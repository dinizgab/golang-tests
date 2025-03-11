.PHONY: run
run: up-db migrate
	docker compose --profile local up --build

args?="./..."
.PHONY: test
test: up-db migrate
	docker compose --profile test run --rm --build go test --tags=integration -failfast -v -p 1 -count=1 $(args)

up-db:
	docker compose up -d database

.PHONY: migrate
migrate:
	goose -dir ./migrations postgres "user=postgres password=mysecretpassword host=127.0.0.1 port=5432 dbname=local_db sslmode=disable" up
