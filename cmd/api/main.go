package main

import (
	"test-go/internal/api/routes"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

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

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	router := gin.New()
	router.Use(func(c *gin.Context) {
		start := time.Now()
		raw := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start).Milliseconds()

		var zeroLog *zerolog.Event
		status := c.Writer.Status()
		switch {
		case status < 400:
			zeroLog = log.Info()
		case status < 500:
			zeroLog = log.Warn()
		default:
			zeroLog = log.Error()
		}
		zeroLog.
			Int("status", c.Writer.Status()).
			Int64("latency", latency).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Str("method", c.Request.Method).
			Str("query", raw).
			Msg("HTTP request completed")
	})
	routes.UseRoutes(router)
}
