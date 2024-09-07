package main

import (
	"encoding/json"
	"log"
	"test-go/internal/api/routes"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
)

func readConsulConfig() string {
	config := api.DefaultConfig()
	config.Address = "http://localhost:8500"
	client, err := api.NewClient(config)

	if err != nil {
		log.Fatalf("Error creating Consul client: %s", err)
	}
	kv := client.KV()
	pair, _, err := kv.Get("cert-service/development", nil)
	if err != nil {
		log.Fatalf("Error creating Consul client: %s", err)
	}
	if pair == nil {
		log.Fatalf("Key 'SignalR.API/appsettings.Development.json' not found")
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

func main() {
	// gin.SetMode(gin.ReleaseMode)
	// var elasticSearchHost = readConsulConfig()
	// search.ConnectElasticSearch(elasticSearchHost)

	// loggerPdt := logger.NewLogger()
	router := gin.New()
	// router.Use(func(c *gin.Context) {
	// 	path := c.Request.URL.Path
	// 	start := time.Now()
	// 	c.Next()
	// 	latency := time.Since(start)
	// 	loggerPdt.Info("HTTP request",
	// 		logger.Field{Key: "status", Value: c.Writer.Status()},
	// 		logger.Field{Key: "method", Value: c.Request.Method},
	// 		logger.Field{Key: "path", Value: path},
	// 		logger.Field{Key: "latency", Value: latency},
	// 	)
	// })
	routes.UseRoutes(router)
}
