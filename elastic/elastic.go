package elastic

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

var EsClient *elasticsearch.Client

func ConnectElasticSearch(host string) *elasticsearch.Client {
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

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	return es
}
