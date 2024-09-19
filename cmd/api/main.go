package main

import (
	"test-go/internal/api/routes"
	"test-go/internal/logger"
	"test-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// func readConsulConfig() string {
// 	config := api.DefaultConfig()
// 	config.Address = "http://localhost:8500"
// 	client, err := api.NewClient(config)

// 	if err != nil {
// 		log.Fatalf("Error creating Consul client: %s", err)
// 	}
// 	kv := client.KV()
// 	pair, _, err := kv.Get("cert-service/development", nil)
// 	if err != nil {
// 		log.Fatalf("Error creating Consul client: %s", err)
// 	}
// 	if pair == nil {
// 		log.Fatalf("Key 'SignalR.API/appsettings.Development.json' not found")
// 	}
// 	strVal := string(pair.Value)

// 	var result map[string]interface{}
// 	readJsonErr := json.Unmarshal([]byte(strVal), &result)
// 	if err != nil {
// 		log.Fatalf("Error unmarshaling JSON: %v", readJsonErr)
// 	}

// 	elasticConfig, ok := result["ElasticConfiguration"].(map[string]interface{})
// 	if !ok {
// 		log.Fatal("Error read elastic configuration with key: ElasticConfiguration")
// 	}
// 	elasticHost, ok := elasticConfig["Uri"].(string)
// 	if !ok {
// 		log.Fatal("Error read elastic host with key: Uri")
// 	}
// 	return elasticHost
// }

func main() {
	gin.SetMode(gin.ReleaseMode)
	// var elasticSearchHost = readConsulConfig()
	// search.ConnectElasticSearch(elasticSearchHost)
	// loggerPdt := logger.NewLogger()

	logger.SetupLogger()
	engine := gin.New()
	engine.Use(middleware.Logger())
	engine.Use(middleware.ErrorHandler())
	routes.SetupRoutes(engine)
	engine.Run(":8080")
}
