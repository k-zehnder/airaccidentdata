.PHONY: build
build:
	@echo "Building scraper Docker image..."
	docker build -t $(SCRAPER_IMAGE_TAG) .

.PHONY: push
push:
	@echo "Pushing scraper Docker image..."
	docker push $(SCRAPER_IMAGE_TAG)

.PHONY: run
run:
	@echo "Starting scraper service..."
	docker-compose up -d scraper

.PHONY: fetch-aircraft-images
fetch-aircraft-images:
	@echo "Collecting aircraft images from Wikipedia..."
	npm run dev 
