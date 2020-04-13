run:
	@docker-compose --project-name=covidarnold up -d

build:
	@mkdir -p bin
	@go build -o bin/ingester cmd/ingester/main.go

forecast:
	@docker build . -t covidarnold_forecaster -f docker/Dockerfile.forecaster