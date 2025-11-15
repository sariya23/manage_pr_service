FROM golang:1.25-alpine as dev

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN apk add --no-cache postgresql-client make
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go build -o service cmd/server/main.go