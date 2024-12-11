tidy:
	go mod tidy
	go mod vendor

# ==============================================================================

BASE_IMAGE_NAME := miniurl
VERSION         := "0.0.1-$(shell git rev-parse --short HEAD)"
APP_IMAGE       := $(BASE_IMAGE_NAME):$(VERSION)

# ==============================================================================
# Building containers

all: server

server:
	docker build \
		-f Dockerfile \
		-t $(APP_IMAGE) \
		.

# ==============================================================================
# Dev envoironment

dev:
	DEV_MINIURL_IMAGE=$(APP_IMAGE) docker compose -f compose.dev.yaml up