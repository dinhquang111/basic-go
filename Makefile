APP_NAME := basic-go-service
DOCKER_USERNAME := darthmalgus1997
DOCKER_REPO := $(DOCKER_USERNAME)/$(APP_NAME)
VERSION := 1.0.6
PORT := 8080
OUTPUT_FILE= bin/main.exe

push:
	@echo "Committing code to current branch..."
	git add .
	@powershell -Command " \
		$$message = Read-Host 'Please enter commit message'; \
		if ($$message -eq '') { echo 'Commit message cannot be empty'; exit 1 } \
		git commit -m $$message; \
		git push; \
	"

release:
	@echo "Releasing code ..."
	@powershell -Command " \
		$$version = Read-Host 'Please enter relase version'; \
		$$message = Read-Host 'Please enter relase message'; \
		if ($$version -eq '') { echo 'Relase version cannot be empty'; exit 1 } \
		git tag -a v$$version -m message; \
		git push origin v$$version; \
		gh release create $(TAG) --title "$(TITLE)" --notes "$(NOTES)"
	"
	@echo "Release created successfully."

run: swag
	go build -o ./bin/main.exe ./cmd/api/main.go
	@cmd /c start http://localhost:8080/swagger/index.html
	@$(OUTPUT_FILE)

merge:
	@powershell -Command " \
		$$targetBranch = Read-Host 'Enter target branch'; \
		if ($$targetBranch -eq '') { echo 'Relase target branch cannot be empty'; exit 1 } \
		git checkout $$targetBranch; \
		git pull; \
		git checkout -; \
		gh pr create --base $$targetBranch; \
	"

docker-run:
	@echo "Building the Docker image..."
	docker build -t $(DOCKER_REPO):$(VERSION) .
	docker run -d -p ${PORT}:${PORT} --name ${APP_NAME} $(DOCKER_REPO):$(VERSION)

docker-push:
	@echo "Pushing the Docker image to Docker Hub..."
	docker push $(DOCKER_REPO):$(VERSION)

docker-rm-container:
	docker stop ${APP_NAME}
	docker rm ${APP_NAME}

swag:
	swag init -g cmd/api/main.go