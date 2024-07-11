package pkg

import (
	"context"
	"strings"
	"time"
	"web_crawler/models"
)

// CombineStrings takes an array of strings and combines them into a single string.
func CombineStrings(stringsArray []string) string {
	return strings.Join(stringsArray, " ") // Combining with space as delimiter
}

// InsertCombinedContent creates a WebPage with combined content and inserts it into Elasticsearch.
func InsertCombinedContent(ctx context.Context, url string, status int, content []string, title, description string, keywords []string) (string, error) {
	combinedContent := CombineStrings(content) // Combine URLs or any strings array into a single string

	page := models.WebPage{
		URL:        url,    // Not applicable or could be a representative URL
		StatusCode: status, // Assuming a default status code
		Content:    combinedContent,
		CrawledAt:  time.Now(),
		Title:      title,
		Desription: description,
		Keywords:   keywords,
	}

	result, err := models.CreateWebPage(ctx, page)
	if err != nil {
		return "", err
	}
	return result, nil

}
