# usage: Нужно указать префикс env файла. То есть
# если используем local.env, то пишем make migrate ENV=local.
ENV ?= dev
ENV_FILE = ./config/$(ENV).env
include ${ENV_FILE}

# DEV
.PHONY: run
run:
	go run cmd/server/main.go --config=./config/dev.env

.PHONY: test
test:
	go test -v ./...

.PHONY: infra
infra:
	docker-compose -p pr_infra -f deployments/docker/dev/docker-compose.yaml  \
	--env-file ./config/dev.env up -d

.PHONY: migrate
migrate:
	goose -dir migrations postgres \
	"postgresql://$(POSTGRES_USERNAME):$(POSTGRES_PASSWORD)\
	@$(POSTGRES_HOST_OUTER):$(POSTGRES_PORT)/$(POSTGRES_DB)\
	?sslmode=$(SSL_MODE)" up
