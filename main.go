package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"test-go/elastic"
	"test-go/routes"

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
	if pair != nil {
		log.Fatalf("Consul value is nil: %s", err)
	}
	strVal := string(pair.Value)

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

var (
	Version   = "dev"
	Commit    = "none"
	BuildTime = "unknown"
)

func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version":   Version,
		"commit":    Commit,
		"buildTime": BuildTime,
	})
}

func main() {
	var elasticSearchHost = readConsulConfig()
	elastic.ConnectElasticSearch(elasticSearchHost)
	router := gin.Default()
	router.GET("/health", HealthCheckHandler)
	routes.UseRoutes(router)

	router.Run(":8080")
}
