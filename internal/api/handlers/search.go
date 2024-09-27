package handlers

import (
	"basic-go/internal/search"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Insert(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid json body",
		})
	}
	res, err := search.EsClient.Index(string(search.Feature), strings.NewReader(string(body)))
	fmt.Println(res, err)
}

func Query(c *gin.Context) {
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
	res, err := search.EsClient.Search(
		search.EsClient.Search.WithContext(context.Background()),
		search.EsClient.Search.WithIndex(string(search.All)),
		search.EsClient.Search.WithBody(strings.NewReader(string(queryJSON))),
		search.EsClient.Search.WithPretty(),
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

// func validateIndexSearch(es *elasticsearch.Client) {
// 	indexName := "cls-feature"
// 	req := esapi.IndicesExistsRequest{
// 		Index: []string{indexName},
// 	}
// 	res, err := req.Do(context.Background(), es)
// 	if err != nil {
// 		log.Fatalf("Error checking if the index exists: %s", err)
// 	}
// 	defer res.Body.Close()
// 	if res.StatusCode == 200 {
// 		return
// 	} else if res.StatusCode == 404 {
// 		es.Indices.Create("cls-feature")
// 	} else {
// 		log.Fatalf("Error checking if the index exists: %s", err)
// 	}
// }
