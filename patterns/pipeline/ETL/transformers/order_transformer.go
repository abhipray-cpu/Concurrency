package transformer

import (
	"ETL/data"
	"time"
)

// NormalizeAndEncryptOrder normalizes and encrypts order data.
func NormalizeAndEncryptOrder(order *data.Order) *data.Order {
	// Encrypt UserID (placeholder for actual encryption)
	order.UserID = "Enc_" + order.UserID

	// Normalize OrderDate to ISO 8601 (YYYY-MM-DD)
	if parsedTime, err := time.Parse("01/02/2006", order.OrderDate); err == nil {
		order.OrderDate = parsedTime.Format("2006-01-02")
	}

	return order
}

// FlagHighQuantity flags orders with a quantity above a certain threshold.
func FlagHighQuantity(order *data.Order) *data.Order {
	const highQuantityThreshold = 100 // Define the threshold
	if order.Quantity > highQuantityThreshold {
		order.OrderID = -order.OrderID // Negative OrderID as a flag for high quantity
	}
	return order
}

// AnonymizeUserData replaces user-related data with placeholders for privacy.
func AnonymizeUserData(order *data.Order) *data.Order {
	order.UserID = "Anonymous"
	order.ProductID = 0 // Assuming ProductID can be set to 0 as a placeholder
	return order
}

// UpdateOrderStatus adds a status field to the order and sets it based on the order date.
func UpdateOrderStatus(order *data.Order) *data.Order {
	const dateFormat = "2006-01-02"
	orderDate, err := time.Parse(dateFormat, order.OrderDate)
	if err != nil {
		return order // Return the order unchanged if the date format is incorrect
	}

	if time.Since(orderDate).Hours() > 24 {
		order.OrderDate += " [Status: Shipped]"
	} else {
		order.OrderDate += " [Status: Processing]"
	}
	return order
}
