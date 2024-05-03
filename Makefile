VERSION ?= latest
DOCKERHUB_REPO := computers33333/airaccidentdata
FRONTEND_IMAGE_TAG := $(DOCKERHUB_REPO)-frontend:$(VERSION)
BACKEND_IMAGE_TAG := $(DOCKERHUB_REPO)-backend:$(VERSION)
IMAGE_SCRAPER_IMAGE_TAG := $(DOCKERHUB_REPO)-aviation_scraper:$(VERSION)

export FRONTEND_IMAGE_TAG BACKEND_IMAGE_TAG IMAGE_SCRAPER_IMAGE_TAG

.PHONY: all
all: build test push

.PHONY: build
build:
	@echo "Building all components..."
	$(MAKE) -C frontend build
	$(MAKE) -C backend build
	$(MAKE) -C aviationScraper build

.PHONY: test
test:
	@echo "Running tests for all components..."
	$(MAKE) -C backend test

.PHONY: push
push:
	@echo "Pushing all images..."
	$(MAKE) -C frontend push
	$(MAKE) -C backend push
	$(MAKE) -C aviationScraper push

.PHONY: deploy
deploy:
	@echo "Deploying application..."
	docker compose down
	docker pull $(FRONTEND_IMAGE_TAG)
	docker pull $(BACKEND_IMAGE_TAG)
	docker compose up -d
	@echo "Application deployed successfully."

.PHONY: dev
dev:
	@echo "Starting development environment..."
	docker compose -f docker-compose-dev.yml up -d --build
