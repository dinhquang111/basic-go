package consul

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
