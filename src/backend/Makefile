deps:
	go mod download

build: deps
	go build -o server ./cmd/server/main.go

docker-build:
	docker build . -t jfreeman/url-backend:0.1
