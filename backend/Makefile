.PHONY: build_backend
build_backend:
	docker build -t $(BACKEND_IMAGE_TAG) .

.PHONY: run_backend
run_backend:
	docker-compose up -d backend

.PHONY: test_backend
test_backend:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

.PHONY: test_coverage
test_coverage:
	@echo "Running Go tests with coverage..."
	go test -coverprofile=coverage.out ./...
	@echo "Generating coverage report..."
	go tool cover -html=coverage.out -o coverage.html
	@echo "Opening coverage report..."
	open coverage.html
