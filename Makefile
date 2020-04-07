run:
	@docker-compose up -d

build:
	@mkdir -p bin
	@go build -o bin/ingester cmd/ingester/main.go
