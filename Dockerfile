FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# RUN go build -ldflags "-X main.Version=1.0.0 -X main.Commit=$(git rev-parse HEAD) -X main.BuildTime=$(date -u +'%Y-%m-%dT%H:%M:%SZ')" -o main ./
RUN go build -o main ./cmd/api

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
