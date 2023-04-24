deps:
	go mod download

build: deps
	go build -o server ./cmd/server/main.go

docker-restart: docker-build
	docker-compose down -v
	docker-compose up -d
	sleep 2
	docker-compose restart app

docker-build:
	docker build . -t jfreeman/url-backend:0.1
