package elastic

type ElasticIndex string

const (
	All     ElasticIndex = "*"
	Feature ElasticIndex = "feature"
	User    ElasticIndex = "user"
	Course  ElasticIndex = "course"
)
