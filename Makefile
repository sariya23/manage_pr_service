# usage: Нужно указать префикс env файла. То есть
# если используем local.env, то пишем make migrate ENV=local.
ENV ?= dev
ENV_FILE = ./config/$(ENV).env
include ${ENV_FILE}

# DEV
.PHONY: run
run:
	go run cmd/server/main.go --config=./config/dev.env

