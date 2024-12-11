FROM golang:1.23.3-alpine3.20 AS develop
WORKDIR /app
CMD ["sh", "-c", "go run cmd/tooling/migrate/main.go && go run cmd/web/main.go"]