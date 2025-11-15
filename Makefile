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


# LOCAL
.PHONY: service_up
service_up:
	docker-compose -p pg_manage_service -f deployments/docker/local/docker-compose.yaml  \
	--env-file ./config/local.env up -d

.PHONY: service_migrate_inner
service_migrate_inner:
	goose -dir migrations postgres \
	"postgresql://$(POSTGRES_USERNAME):$(POSTGRES_PASSWORD)\
	@$(POSTGRES_HOST_INNER):$(POSTGRES_PORT)/$(POSTGRES_DB)\
	?sslmode=$(SSL_MODE)" up


.PHONY: service_down
service_down:
	docker-compose -p pg_manage_service -f deployments/docker/local/docker-compose.yaml \
	--env-file ./config/local.env rm -fvs


# TEST

# ДЛЯ ТЕСТОВ В ДОКЕРЕ
.PHONY: test_service_up
test_service_up:
	docker-compose -p test_pr_manage_service -f deployments/docker/test/docker-compose.yaml  \
	--env-file ./config/test.env up -d

.PHONY: test_migrate
test_migrate:
	goose -dir migrations postgres \
	"postgresql://$(POSTGRES_USERNAME):$(POSTGRES_PASSWORD)\
	@$(POSTGRES_HOST_INNER):$(POSTGRES_PORT)/$(POSTGRES_DB)\
	?sslmode=$(SSL_MODE)" up

.PHONY: test_integrations
test_integrations:
	 go test -v -tags=integrations ./tests/...

.PHONY: test_service_down
test_service_down:
	docker-compose -p test_pr_manage -f deployments/docker/test/docker-compose.yaml \
	--env-file ./config/test.env rm -fvs