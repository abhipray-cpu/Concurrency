package main

import (
	transformer "ETL/transformers"
	"fmt"
	"strings"
)

func Transform(pipeline <-chan DataWrapper) <-chan DataWrapper {
	fmt.Println("Transforming data...")
	output := make(chan DataWrapper)
	go func() {
		for data := range pipeline {
			switch {
			case data.User != nil:
				data.User = transformer.CleanAndHashUser(data.User)

			case data.Order != nil:
				data.Order = transformer.NormalizeAndEncryptOrder(data.Order)
				data.Order = transformer.FlagHighQuantity(data.Order)
				data.Order = transformer.AnonymizeUserData(data.Order)
				data.Order = transformer.UpdateOrderStatus(data.Order)

			case data.Product != nil:
				data.Product = CleanAndAdjustProduct(data.Product)
				data.Product = MarkOutOfStock(data.Product)
				data.Product = TagPriceRange(data.Product)
				data.Product = NormalizeCategoryName(data.Product)
			}
			output <- data
		}
		// Close the output channel when done
		close(output)
	}()
	return output
}

// CleanAndAdjustProduct cleans product data and adjusts its price based on category.
func CleanAndAdjustProduct(product *Product) *Product {
	// Trim leading and trailing spaces from ProductName and Category
	product.ProductName = strings.TrimSpace(product.ProductName)
	product.Category = strings.TrimSpace(product.Category)

	// Correct common data entry errors in Category
	switch strings.ToLower(product.Category) {
	case "elec":
		product.Category = "Electronics"
	case "bk":
		product.Category = "Books"
	}

	// Adjust price based on category
	if product.Category == "Electronics" {
		product.Price *= 1.05 // Increase price by 5%
	} else if product.Category == "Books" {
		product.Price *= 0.95 // Decrease price by 5%
	}

	return product
}

// MarkOutOfStock adds an "OutOfStock" tag to the product name if the stock quantity is 0.
func MarkOutOfStock(product *Product) *Product {
	if product.StockQuantity == 0 {
		product.ProductName = "[OutOfStock] " + product.ProductName
	}
	return product
}

// TagPriceRange adds a tag to the product name based on its price range.
func TagPriceRange(product *Product) *Product {
	switch {
	case product.Price < 50:
		product.ProductName = "[Budget] " + product.ProductName
	case product.Price >= 50 && product.Price <= 200:
		product.ProductName = "[Midrange] " + product.ProductName
	case product.Price > 200:
		product.ProductName = "[Premium] " + product.ProductName
	}
	return product
}

// NormalizeCategoryName ensures consistent capitalization for category names.
func NormalizeCategoryName(product *Product) *Product {
	product.Category = strings.Title(strings.ToLower(product.Category))
	return product
}
