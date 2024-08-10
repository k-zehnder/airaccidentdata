VERSION ?= latest
DOCKERHUB_REPO := computers33333/airaccidentdata
FRONTEND_IMAGE_TAG := $(DOCKERHUB_REPO)-frontend:$(VERSION)
BACKEND_IMAGE_TAG := $(DOCKERHUB_REPO)-backend:$(VERSION)
AIRCRAFT_SCRAPER_IMAGE_TAG := $(DOCKERHUB_REPO)-aircraft_scraper:$(VERSION)
ELASTIC_INDEXER_IMAGE_TAG := $(DOCKERHUB_REPO)-elastic_indexer:$(VERSION)

export FRONTEND_IMAGE_TAG BACKEND_IMAGE_TAG AIRCRAFT_SCRAPER_IMAGE_TAG ELASTIC_INDEXER_IMAGE_TAG

.PHONY: all
all: build test push

.PHONY: build
build:
	@echo "Building all components..."
	$(MAKE) -C frontend build
	$(MAKE) -C backend build
	$(MAKE) -C aircraft_scraper build

.PHONY: test
test:
	@echo "Running tests for all components..."
	$(MAKE) -C backend test

.PHONY: push
push:
	@echo "Pushing all images..."
	$(MAKE) -C frontend push
	$(MAKE) -C backend push
	$(MAKE) -C aircraft_scraper push
	$(MAKE) -C elastic push

.PHONY: deploy
deploy:
	@echo "Deploying application..."
	docker compose down
	$(MAKE) -C frontend pull
	$(MAKE) -C backend pull
	$(MAKE) -C aircraft_scraper pull
	$(MAKE) -C elastic pull
	docker compose up -d
	@echo "Application deployed successfully."

.PHONY: dev
dev:
	@echo "Starting development environment..."
	docker compose -f docker-compose-dev.yml up -d --build
