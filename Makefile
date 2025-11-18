include ${ENV_FILE}

# DEV
.PHONY: run
run:
	go run cmd/server/main.go --config=./config/dev.env

.PHONY: test
test:
	go clean -testcache && go test -v ./...

.PHONY: migrate
migrate:
	goose -dir migrations postgres \
	"postgresql://$(POSTGRES_USERNAME):$(POSTGRES_PASSWORD)\
	@$(POSTGRES_HOST_OUTER):$(POSTGRES_PORT)/$(POSTGRES_DB)\
	?sslmode=$(SSL_MODE)" up

.PHONY: lint
lint:
	revive ./...

.PHONY: gen
gen:
	oapi-codegen --config .oapi-codegen.yaml api/openapi.yaml

## LOCAL
.PHONY: service_migrate_inner
service_migrate_inner:
	goose -dir migrations postgres \
	"postgresql://$(POSTGRES_USERNAME):$(POSTGRES_PASSWORD)\
	@$(POSTGRES_HOST_INNER):$(POSTGRES_PORT)/$(POSTGRES_DB)\
	?sslmode=$(SSL_MODE)" up

.PHONY: test_integrations
test_integrations:
	 go clean -testcache && go test -v -tags=integrations ./tests/...
