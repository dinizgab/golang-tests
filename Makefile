args?="./..."
.PHONY: test
test: up-db migrate
	docker compose --profile test run --rm --build apitest go test --tags=integration -failfast -v $(args)

up-db:
	docker compose up -d database

.PHONY: migrate
migrate:
	goose -dir ./migrations postgres "user=postgres password=mysecretpassword host=127.0.0.1 port=5432 dbname=local_db sslmode=disable" up
