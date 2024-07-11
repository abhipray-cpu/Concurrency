package models

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"go.mongodb.org/mongo-driver/mongo"
)

var es *elasticsearch.Client

type WebPage struct {
	ID         string    `json:"id"`
	URL        string    `json:"url"`
	StatusCode int       `json:"status_code"`
	Content    string    `json:"content"`
	CrawledAt  time.Time `json:"crawled_at"`
	Title      string    `json:"title"`
	Desription string    `json:"description"`
	Keywords   []string  `json:"keywords"`
}

type Models struct {
	WebPage WebPage
	User    User
}

func NewModels(esClient *elasticsearch.Client, mongClient *mongo.Client) *Models {
	es = esClient
	client = mongClient
	return &Models{
		WebPage: WebPage{},
		User:    User{},
	}
}

func CreateWebPage(ctx context.Context, page WebPage) (string, error) {
	data, err := json.Marshal(page)
	if err != nil {
		log.Printf("Error marshalling data: %v", err)
		return "", err
	}

	req := esapi.IndexRequest{
		Index:      "webpages",
		DocumentID: "", // leaving empty to let elastic search generate a unique id
		Body:       strings.NewReader(string(data)),
		Refresh:    "true",
	}

	res, err := req.Do(ctx, es)

	if err != nil {
		log.Printf("Error getting response: %v", err)
		return "", err
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("Error indexing document: %s", res.Status())
		return "", fmt.Errorf("error indexing document: %s", res.String())
	}

	var r map[string]interface{}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Printf("Error parsing the response body: %s", err)
		return "", err
	}
	return r["_id"].(string), nil

}

func ReadWebPage(ctx context.Context, id string) (*WebPage, error) {
	req := esapi.GetRequest{
		Index:      "webpages",
		DocumentID: id,
	}

	res, err := req.Do(ctx, es)
	if err != nil {
		log.Printf("Error getting response: %v", err)
		return nil, err
	}

	defer res.Body.Close()
	if res.IsError() {
		log.Printf("Error getting document: %s", res.Status())
		return nil, fmt.Errorf("error getting document: %s", res.String())
	}
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Printf("Error parsing the response body: %s", err)
		return nil, err
	}

	// Safely assert _source as map[string]interface{}
	doc, ok := r["_source"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("error asserting _source")
	}

	// No need to double extract _source
	crawledAtStr, ok := doc["crawled_at"].(string)
	if !ok {
		return nil, fmt.Errorf("error asserting crawled_at")
	}
	crawledAt, err := time.Parse(time.RFC3339, crawledAtStr)
	if err != nil {
		log.Printf("Error parsing crawled_at time: %s", err)
		return nil, err
	}

	// Safely assert other fields
	url, ok := doc["url"].(string)
	if !ok {
		return nil, fmt.Errorf("error asserting url")
	}

	statusCode, ok := doc["status_code"].(float64)
	if !ok {
		return nil, fmt.Errorf("error asserting status_code")
	}

	content, ok := doc["content"].(string)
	if !ok {
		return nil, fmt.Errorf("error asserting content")
	}

	title, ok := doc["title"].(string)
	if !ok {
		return nil, fmt.Errorf("error asserting title")
	}

	page := WebPage{
		ID:         id, // ID is passed as a parameter, no need to extract from doc
		URL:        url,
		StatusCode: int(statusCode),
		Content:    content,
		CrawledAt:  crawledAt,
		Title:      title,
	}

	return &page, nil
}

// Helper function to convert an interface slice to a string slice.
func convertInterfaceToStringSlice(interfaceSlice []interface{}) []string {
	var stringSlice []string
	for _, v := range interfaceSlice {
		stringSlice = append(stringSlice, v.(string))
	}
	return stringSlice
}
func UpdateWebPage(ctx context.Context, id string, page WebPage) error {
	data, err := json.Marshal(page)
	if err != nil {
		log.Printf("Error marshalling data: %v", err)
		return err
	}
	req := esapi.UpdateRequest{
		Index:      "webpages",
		DocumentID: id,
		Body:       strings.NewReader(fmt.Sprintf(`{"doc":%s}`, string(data))),
	}

	res, err := req.Do(ctx, es)
	if err != nil {
		log.Printf("Error getting response: %v", err)
		return err
	}

	defer res.Body.Close()
	if res.IsError() {
		log.Printf("Error updating document: %s", res.Status())
		return fmt.Errorf("error updating document: %s", res.String())

	}
	return nil
}

func DeleteWebPage(ctx context.Context, id string) error {
	req := esapi.DeleteRequest{
		Index:      "webpages",
		DocumentID: id,
	}

	res, err := req.Do(ctx, es)
	if err != nil {
		log.Printf("Error getting response: %v", err)
		return err
	}
	defer res.Body.Close()
	if res.IsError() {
		log.Printf("Error deleting document: %s", res.Status())
		return fmt.Errorf("error deleting document: %s", res.String())

	}
	return nil
}

func SearchWebPage(ctx context.Context, query string) ([]WebPage, error) {
	var pages []WebPage

	req := esapi.SearchRequest{
		Index: []string{"webpages"},
		Body:  strings.NewReader(fmt.Sprintf(`{"query": {"match": {"content": "%s"}}}`, query)),
	}

	res, err := req.Do(ctx, es)
	if err != nil {
		log.Printf("Error getting response: %v", err)
		return nil, err
	}

	defer res.Body.Close()

	if res.IsError() {
		log.Printf("Error searching document: %s", res.Status())
		return nil, fmt.Errorf("error searching document: %s", res.String())
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Printf("Error parsing the response body: %s", err)
		return nil, err
	}

	hitsWrapper, ok := r["hits"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("error asserting hits")
	}

	hits, ok := hitsWrapper["hits"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("error asserting hits array")
	}

	for _, hitInterface := range hits {
		hit, ok := hitInterface.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("error asserting hit")
		}

		source, ok := hit["_source"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("error asserting source")
		}

		crawledAtStr, ok := source["crawled_at"].(string)
		if !ok {
			return nil, fmt.Errorf("error asserting crawled_at")
		}
		crawledAt, err := time.Parse(time.RFC3339, crawledAtStr)
		if err != nil {
			log.Printf("Error parsing crawled_at time: %s", err)
			return nil, err
		}

		id, ok := hit["_id"].(string)
		if !ok {
			return nil, fmt.Errorf("error asserting id")
		}

		url, ok := source["url"].(string)
		if !ok {
			return nil, fmt.Errorf("error asserting url")
		}

		statusCode, ok := source["status_code"].(float64)
		if !ok {
			return nil, fmt.Errorf("error asserting status_code")
		}

		content, ok := source["content"].(string)
		if !ok {
			return nil, fmt.Errorf("error asserting content")
		}

		title, ok := source["title"].(string)
		if !ok {
			return nil, fmt.Errorf("error asserting title")
		}

		page := WebPage{
			ID:         id,
			URL:        url,
			StatusCode: int(statusCode),
			Content:    content,
			CrawledAt:  crawledAt,
			Title:      title,
		}
		pages = append(pages, page)
	}
	return pages, nil
}

func GetWebPages(ctx context.Context) ([]WebPage, error) {
	req := esapi.SearchRequest{
		Index: []string{"webpages"},
		Body:  strings.NewReader(`{"query": {"match_all": {}}}`),
	}
	res, err := req.Do(ctx, es)
	if err != nil {
		log.Printf("Error getting response: %v", err)
		return nil, err
	}
	defer res.Body.Close()
	if res.IsError() {
		log.Printf("Error searching document: %s", res.Status())
		return nil, fmt.Errorf("error searching document: %s", res.String())
	}
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Printf("Error parsing the response body: %s", err)
		return nil, err
	}
	var pages []WebPage
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		doc := hit.(map[string]interface{})
		source := doc["_source"].(map[string]interface{})
		crawledAtStr := source["crawled_at"].(string)            // Assert the type as string
		crawledAt, err := time.Parse(time.RFC3339, crawledAtStr) // Parse the string to time.Time
		if err != nil {
			log.Printf("Error parsing crawled_at time: %s", err)
			return nil, err
		}
		page := WebPage{
			ID:         doc["_id"].(string), // Extract the ID
			URL:        source["url"].(string),
			StatusCode: int(source["status_code"].(float64)),
			Content:    source["content"].(string),
			CrawledAt:  crawledAt, // Use the parsed time
			Title:      source["title"].(string),
			// Continue with the rest of your fields...
		}
		pages = append(pages, page)
	}
	return pages, nil
}
