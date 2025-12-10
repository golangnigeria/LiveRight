# Makefile for LiveRight backend using Goose

# Default migration directory
MIGRATIONS_DIR = migrations
ENV_FILE = .env
GO_CMD = go run ./cmd/api

# Run all pending migrations
.PHONY: migrate
migrate:
	@echo "Applying migrations from $(MIGRATIONS_DIR) using Goose..."
	goose -env $(ENV_FILE) -dir $(MIGRATIONS_DIR) up
	@echo "Migrations applied successfully!"

# Run migrations and then start the server
.PHONY: dev
dev:
	@echo "Running migrations and starting the server..."
	$(MAKE) migrate
	$(GO_CMD)

# Rollback last migration
.PHONY: rollback
rollback:
	@echo "Rolling back last migration..."
	goose -env $(ENV_FILE) -dir $(MIGRATIONS_DIR) down
	@echo "Rollback completed!"

# Check migration status
.PHONY: status
status:
	@echo "Checking migration status..."
	goose -env $(ENV_FILE) -dir $(MIGRATIONS_DIR) status

# Create a new migration
# Usage: make create NAME=add_some_table
.PHONY: create
create:
ifndef NAME
	$(error NAME is not set. Usage: make create NAME=add_some_table)
endif
	@echo "Creating new migration: $(NAME)"
	goose -dir $(MIGRATIONS_DIR) create $(NAME) sql
	@echo "Migration created in $(MIGRATIONS_DIR)"

# Reset database (roll back all migrations)
.PHONY: reset
reset:
	@echo "Resetting all migrations..."
	goose -env $(ENV_FILE) -dir $(MIGRATIONS_DIR) reset
	@echo "Database reset completed!"

