run:
	@docker-compose --project-name=covidarnold up -d

build:
	@mkdir -p bin
	@go build -o bin/ingester cmd/ingester/main.go

forecast:
	@docker build . -t covidarnold_forecaster -f docker/Dockerfile.forecaster
	@docker run \
		-it \
		--rm \
		--name=forecaster \
		--net=covidarnold_data \
		-e DB_USER_NAME=covid \
		-e DB_PASSWORD=arnold \
		-e DB_NAME=covid \
		-e DB_HOSTNAME=postgres \
		covidarnold_forecaster:latest \
		bash