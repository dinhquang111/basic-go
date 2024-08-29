package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
)

func readConsulConfig() string {
	config := api.DefaultConfig()
	config.Address = "http://192.168.222.150:8500"
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("Error creating Consul client: %s", err)
	}
	kv := client.KV()
	pair, _, err := kv.Get("SignalR.API/appsettings.Development.json", nil)
	if err != nil {
		log.Fatalf("Error creating Consul client: %s", err)
	}
	if pair == nil {
		fmt.Println("Key 'SignalR.API/appsettings.Development.json' not found")
	}

	strVal := string(pair.Value)
	if err != nil {
		log.Fatalf("Consul value is nil: %s", err)
	}

	var result map[string]interface{}
	readJsonErr := json.Unmarshal([]byte(strVal), &result)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", readJsonErr)
	}

	elasticConfig, ok := result["ElasticConfiguration"].(map[string]interface{})
	if !ok {
		log.Fatal("Error read elastic configuration with key: ElasticConfiguration")
	}
	elasticHost, ok := elasticConfig["Uri"].(string)
	if !ok {
		log.Fatal("Error read elastic host with key: Uri")
	}
	return elasticHost
}

func connectElasticSearch(host string) *elasticsearch.Client {
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			host,
		},
		Username: "elastic",
		Password: "malgus123",
	})

	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// Verify that the connection is working
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	return es
}

func validateIndexSearch(es *elasticsearch.Client) {
	indexName := "cls-feature"
	req := esapi.IndicesExistsRequest{
		Index: []string{indexName},
	}
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error checking if the index exists: %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		return
	} else if res.StatusCode == 404 {
		es.Indices.Create("cls-feature")
	} else {
		log.Fatalf("Error checking if the index exists: %s", err)
	}
}

func insertSearch(es *elasticsearch.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		indexName := "course"
		body, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid json body",
			})
		}
		res, err := es.Index(indexName, strings.NewReader(string(body)))
		// response, err := es.Get(indexName, "-S4vmJEBt5q-sVw_Lt80")
		fmt.Println(res, err)
	}
}
func handleSearch(es *elasticsearch.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// indexName := "cls-feature"
		query := c.Query("q")
		queryBody := map[string]interface{}{
			"query": map[string]interface{}{
				"query_string": map[string]interface{}{
					"query": query,
				},
			},
		}
		queryJSON, err := json.Marshal(queryBody)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal query"})
			return
		}
		res, err := es.Search(
			es.Search.WithContext(context.Background()),
			es.Search.WithIndex("*"),
			es.Search.WithBody(strings.NewReader(string(queryJSON))),
			es.Search.WithPretty(),
		)
		// Check for errors in the response
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Elasticsearch error: %s", res.String())})
			return
		}

		// Parse the search results
		var response map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
			return
		}

		// Extract hits from the response
		hits := response["hits"].(map[string]interface{})["hits"].([]interface{})
		results := []map[string]interface{}{}

		for _, hit := range hits {
			doc := hit.(map[string]interface{})["_source"].(map[string]interface{})
			results = append(results, doc)
		}

		c.JSON(http.StatusOK, results)
	}
}

func main() {
	var elasticSearchHost = readConsulConfig()
	esClient := connectElasticSearch(elasticSearchHost)
	fmt.Println(esClient)
	router := gin.Default()
	version := 4
	router.GET("/ping/:name", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong " + c.Params.ByName("name"),
			"version": "v" + strconv.Itoa(version),
		})
	})

	router.POST("/search", insertSearch(esClient))
	router.GET("/search", handleSearch(esClient))
	router.Run(":8080")
}
