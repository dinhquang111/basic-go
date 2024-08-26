APP_NAME := basic-go-service
DOCKER_USERNAME := darthmalgus1997
DOCKER_REPO := $(DOCKER_USERNAME)/$(APP_NAME)
VERSION := 1.0.0
PORT := 8080

push-code:
	git add .
	git commit -m "fix"
	git push

build:
	go run main.go

docker-build:
	@echo "Building the Docker image..."
	docker build -t $(DOCKER_REPO):$(VERSION) .

docker-run:
	docker run -d -p ${PORT}:${PORT} --name ${APP_NAME} $(DOCKER_REPO):$(VERSION)

docker-push: docker-build
	@echo "Pushing the Docker image to Docker Hub..."
	docker push $(DOCKER_REPO):$(VERSION)

docker-rm-container:
	docker stop ${APP_NAME}
	docker rm ${APP_NAME}