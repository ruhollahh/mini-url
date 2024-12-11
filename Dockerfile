FROM golang:1.23.3-alpine3.20 AS develop
WORKDIR /app
RUN go install github.com/air-verse/air@latest
CMD ["sh", "-c", "go run cmd/tooling/migrate/main.go && air -c .air.toml"]