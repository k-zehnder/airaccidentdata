.PHONY: build
build:
	@echo "Building frontend Docker image..."
	docker build -t $(FRONTEND_IMAGE_TAG) .

.PHONY: push
push:
	@echo "Pushing frontend Docker image..."
	docker push $(FRONTEND_IMAGE_TAG)

.PHONY: run
run:
	@echo "Starting frontend service..."
	docker-compose up -d frontend

.PHONY: dev
dev:
	@echo "Setting up development environment for frontend..."
	npm install
	npm run dev
