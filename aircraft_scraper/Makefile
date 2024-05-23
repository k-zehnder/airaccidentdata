.PHONY: build
build:
	@echo "Building aircraft_scraper Docker image..."
	docker build -t $(AIRCRAFT_SCRAPER_IMAGE_TAG) .

.PHONY: push
push:
	@echo "Pushing aircraft_scraper Docker image..."
	docker push $(AIRCRAFT_SCRAPER_IMAGE_TAG)

.PHONY: run
run:
	@echo "Starting aircraft_scraper service..."
	docker compose up -d aircraft_scraper

.PHONY: fetch-images 
fetch-images:
	@echo "Collecting aircraft images from Wikipedia..."
	docker exec -it $(shell docker ps -qf "name=aircraft_scraper") sh -c "npm run start"
