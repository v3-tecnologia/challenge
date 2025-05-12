APP_NAME=telemetry-challenge-api

.PHONY: migrate-create migrate-up migrate-down migrate-force

migrate-create:
	@if [ -z "$(name)" ]; then echo "Erro: Especifique um nome. Ex: make migrate-create name=add_users_table"; exit 1; fi
	@migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)
	@echo "Migration criada em $(MIGRATIONS_DIR)"

migrate-up:
	@if [ -z "$(step)" ]; then \
		migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up; \
	else \
		migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up $(step); \
	fi

migrate-down:
	@if [ -z "$(step)" ]; then \
		migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down 1; \
	else \
		migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down $(step); \
	fi
