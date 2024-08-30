choco install gh
docker run -p 9200:9200 -d --name elasticsearch --network elastic-net -e ELASTIC_PASSWORD=malgus123 -e "discovery.type=single-node" -e "xpack.security.http.ssl.enabled=false" -e "xpack.license.self_generated.type=trial" docker.elastic.co/elasticsearch/elasticsearch:8.15.0
go get github.com/elastic/go-elasticsearch/v8