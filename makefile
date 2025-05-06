ifneq (,$(wildcard .env))
	include .env
	export
endif

MIGRATE_PATH=./internal/infra/database/migrations

migrate-up:
	@migrate -path $(MIGRATE_PATH) -database "postgres://${DATABASE_USER}:${DATABASE_PASSWORD}@localhost:5432/${DATABASE_DB}?sslmode=disable" --verbose up

migrate-down:
	@migrate -path $(MIGRATE_PATH) -database "postgres://${DATABASE_USER}:${DATABASE_PASSWORD}@localhost:5432/${DATABASE_DB}?sslmode=disable" --verbose down

migrate-create:
	@migrate create -ext sql -dir $(MIGRATE_PATH) $(NAME)