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
	docker compose up -d backend

.PHONY: test
test:
	@echo "Running Go tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	@echo "Generating coverage report..."
	go tool cover -html=coverage.out -o coverage.html

.PHONY: data
data:
	@echo "Building CSV downloader and CSV-to-MySQL processor in the container..."
	docker exec -it $(shell docker ps -qf "name=backend") sh -c "go build -o /app/csvdownloader ./cmd/csvdownloader/main.go && go build -o /app/csvtomysql ./cmd/csvtomysql/main.go"
	@echo "Running CSV downloader and CSV-to-MySQL processor in the container..."
	docker exec -it $(shell docker ps -qf "name=backend") sh -c "/app/csvdownloader && /app/csvtomysql"
