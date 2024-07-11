package pkg

import (
	"context"
	"time"

	"github.com/chromedp/chromedp"
)

func Fetch(url string) (string, error) {
	// Create a context with a timeout to ensure a maximum amount of time spent fetching
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Create a new chromedp context
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	var htmlContent string

	// Run tasks
	// Navigate to the site and get the rendered HTML after JavaScript execution
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.OuterHTML("html", &htmlContent, chromedp.ByQuery),
	)
	if err != nil {
		return "", err
	}

	return htmlContent, nil
}
