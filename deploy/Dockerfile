# 1) Stage для установки Goose (миграции)
FROM golang:1.23 AS goose-base

WORKDIR /tmp/goose
RUN go mod init goose-empty
RUN CGO_ENABLED=0 GOOS=linux go install github.com/pressly/goose/v3/cmd/goose@latest

# 2) Stage для сборки приложения
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/avito-shop ./cmd

# 3) Stage для тестирования
FROM golang:1.23 AS tester

WORKDIR /app

COPY --from=builder /app /app
WORKDIR /app

RUN apk --no-cache add ca-certificates bash

RUN go test ./tests/unit_test/... -v

RUN go test ./tests/e2e/... -v

# 4) Финальный минимальный образ
FROM alpine:latest

WORKDIR /root/

RUN apk --no-cache add ca-certificates bash

COPY --from=goose-base /go/bin/goose /usr/local/bin/goose
COPY --from=builder /app/bin/avito-shop /usr/local/bin/avito-shop
COPY --from=builder /app/migrations /migrations

# Запуск приложения
CMD ["/usr/local/bin/avito-shop"]