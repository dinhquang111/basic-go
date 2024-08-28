package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
)

func ReadConsulConfig() {
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
	fmt.Print(elasticHost)
}
func main() {
	ReadConsulConfig()
	router := gin.Default()
	version := 4
	router.GET("/ping/:name", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong " + c.Params.ByName("name"),
			"version": "v" + strconv.Itoa(version),
		})
	})

	// Run the server
	router.Run(":8080")
}
