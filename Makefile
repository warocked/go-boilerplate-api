.PHONY: migrate-up migrate-version migrate-validate migrate-create build run

# Database migrations
migrate-up:
	@echo "Running database migrations..."
	go run ./cmd/migrate -command up

migrate-version:
	@echo "Checking migration version..."
	go run ./cmd/migrate -command version

migrate-validate:
	@echo "Validating migrations..."
	go run ./cmd/migrate -command validate

migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "Error: Migration name is required. Use: make migrate-create name=your_migration_name"; \
		exit 1; \
	fi
	@echo "Creating migration files: $(name)"
	@migrate create -ext sql -dir internal/api/db/migrations -seq $(name) || \
	 echo "Note: If migrate CLI is not installed, create files manually:"
	@echo "  1. Create internal/api/db/migrations/XXXXXX_$(name).up.sql"
	@echo "  2. Create internal/api/db/migrations/XXXXXX_$(name).down.sql"
	@echo "  (Replace XXXXXX with next sequential version number)"

# Build commands
build:
	@echo "Building application..."
	go build -o api.exe ./cmd/api

build-migrate:
	@echo "Building migration tool..."
	go build -o migrate.exe ./cmd/migrate

# Run application
run:
	@echo "Running application..."
	go run ./cmd/api
