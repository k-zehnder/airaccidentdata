.PHONY: index
index:
	@echo "Indexing data into Elasticsearch..."
	@echo "Use `make index ARGS=--clear` to clear index"
	docker exec -it $(shell docker ps -qf "name=elastic_indexer") sh -c "ts-node src/index.ts $(ARGS)"

.PHONY: build
build:
	@echo "Building Elasticsearch Indexer Docker image..."
	docker build -t $(ELASTIC_INDEXER_IMAGE_TAG) .

.PHONY: push
push:
	@echo "Pushing Elasticsearch Indexer Docker image..."
	docker push $(ELASTIC_INDEXER_IMAGE_TAG)

.PHONY: pull
pull:
	@echo "Pulling Elasticsearch Indexer Docker image..."
	docker pull $(ELASTIC_INDEXER_IMAGE_TAG)

.PHONY: run
run:
	@echo "Running Elasticsearch Indexer service..."
	docker compose up -d elastic_indexer
