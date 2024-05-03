.PHONY: build
build:
	@echo "Building aviation scraper Docker image..."
	docker build -t $(AVIATION_SCRAPER_IMAGE_TAG) .

.PHONY: push
push:
	@echo "Pushing aviation scraper Docker image..."
	docker push $(AVIATION_SCRAPER_IMAGE_TAG)

.PHONY: run
run:
	@echo "Starting aviation scraper service..."
	docker-compose up -d aviation_scraper

.PHONY: fetch-aircraft-images
fetch-aircraft-images:
	@echo "Collecting aircraft images from Wikipedia..."
	npm run dev 
