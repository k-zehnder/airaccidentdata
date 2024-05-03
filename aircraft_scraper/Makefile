.PHONY: build
build:
	@echo "Building aircraft scraper Docker image..."
	docker build -t $(AIRCRAFT_SCRAPER_IMAGE_TAG) .

.PHONY: push
push:
	@echo "Pushing aircraft scraper Docker image..."
	docker push $(AIRCRAFT_SCRAPER_IMAGE_TAG)

.PHONY: run
run:
	@echo "Starting aircraft scraper service..."
	docker-compose up -d aircraft_scraper

.PHONY: fetch-aircraft-images
fetch-aircraft-images:
	@echo "Collecting aircraft images from Wikipedia..."
	npm run dev 
