PKG=./__tests__/entity ./__tests__/usecase
.PHONY: test generate_sqlc create_migration up

# ===============================
# Run all unit tests

test:
	go test $(PKG) -v


# ===============================
# Run SQLC code generation

generate_sqlc:
	sqlc generate --file sqlc.yaml


# ===============================
# Create new migration

create_migration:
	@read -p "Enter migration name: " name; \
		migrate create -ext sql -dir database/migrations -seq $$name


# ===============================
# Run full environment (API + DB + migrate)

up:
	docker-compose up --build
