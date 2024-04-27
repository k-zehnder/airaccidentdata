.PHONY: build
build:
	@echo "Building backend Docker image..."
	docker build -t $(BACKEND_IMAGE_TAG) .

.PHONY: push
push:
	@echo "Pushing backend Docker image..."
	docker push $(BACKEND_IMAGE_TAG)

.PHONY: run
run:
	@echo "Starting backend service..."
	docker-compose up -d backend

.PHONY: test
test:
	@echo "Running backend tests..."
	go test -v -coverprofile=coverage.out ./...

.PHONY: test_coverage
test_coverage:
	@echo "Running Go tests with coverage..."
	go test -coverprofile=coverage.out ./...
	@echo "Generating coverage report..."
	go tool cover -html=coverage.out -o coverage.html
	@echo "Opening coverage report..."
	open coverage.html
