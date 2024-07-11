SHELL=/bin/bash
CRAWLER_BINARY=CRAWLER
## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_crawler
	@echo "Stopping docker images (if running)..."
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:	
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## builds the crawler binary
build_crawler:
	@echo Building crawler binary...
	cd ../backend && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${CRAWLER_BINARY} ./cmd/api
	@echo Done!


