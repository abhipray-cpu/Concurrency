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
	fmt.Println(r["_id"].(string))
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
	doc := r["_source"].(map[string]interface{})
	page := WebPage{
		URL:        doc["url"].(string),
		StatusCode: int(doc["status_code"].(float64)),
		Content:    doc["content"].(string),
		CrawledAt:  time.Time(doc["crawled_at"].(time.Time)),
		Title:      doc["title"].(string),
		Desription: doc["description"].(string),
		Keywords:   convertInterfaceToStringSlice(doc["keywords"].([]interface{})),
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

	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		doc := hit.(map[string]interface{})["_source"].(map[string]interface{})

		page := WebPage{
			URL:        doc["url"].(string),
			StatusCode: int(doc["status_code"].(float64)),
			Content:    doc["content"].(string),
			CrawledAt:  time.Time(doc["crawled_at"].(time.Time)),
			Title:      doc["title"].(string),
			Desription: doc["description"].(string),
			Keywords:   convertInterfaceToStringSlice(doc["keywords"].([]interface{})),
		}
		pages = append(pages, page)
	}
	return pages, nil
}
