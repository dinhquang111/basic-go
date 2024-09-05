package models

type Config struct {
	Services ElasticConfig
}

type ElasticConfig struct {
	url string
}
