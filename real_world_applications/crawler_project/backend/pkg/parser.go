package pkg

import (
	"strings"

	"golang.org/x/net/html"
)

// Parse now also returns the title, description, and keywords of the HTML content.
func Parse(htmlContent string) ([]string, []string, string, string, []string, error) {
	urls := []string{}
	contents := []string{}
	var title, description string
	keywords := []string{} // Use a slice for keywords

	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return nil, nil, "", "", nil, err
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "a": // Handle both <a> and <link> tags for URLs
				for _, a := range n.Attr {
					if a.Key == "href" {
						urls = append(urls, a.Val)
						break
					}
				}
			case "title":
				if n.FirstChild != nil {
					title = n.FirstChild.Data
				}
			case "meta":
				for _, a := range n.Attr {
					if a.Key == "name" && a.Val == "keywords" {
						for _, a := range n.Attr {
							if a.Key == "content" {
								// Assuming keywords are comma-separated
								keywords = append(keywords, strings.Split(a.Val, ",")...)
							}
						}
					}
				}
			case "p", "h1", "h2", "h3", "h4", "h5", "h6":
				if n.FirstChild != nil {
					contents = append(contents, strings.TrimSpace(n.FirstChild.Data))
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	// Trim and remove duplicates from keywords
	keywords = uniqueAndTrim(keywords)

	return urls, contents, title, description, keywords, nil
}

// uniqueAndTrim cleans up the keywords slice by trimming spaces and removing duplicates.
func uniqueAndTrim(items []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range items {
		entry = strings.TrimSpace(entry)
		if _, value := keys[entry]; !value && entry != "" {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
