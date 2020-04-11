run:
	@docker-compose --project-name=covidarnold up -d

build:
	@mkdir -p bin
	@go build -o bin/ingester cmd/ingester/main.go
