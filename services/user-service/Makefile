.PHONY: migrate-up migrate-down migrate-force migrate-version migrate-create

DB_URL=postgresql://user_service:user123@localhost:5432/user_service_db?sslmode=disable

migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" down

migrate-force:
	migrate -path migrations -database "$(DB_URL)" force $(VERSION)

migrate-version:
	migrate -path migrations -database "$(DB_URL)" version

migrate-create:
	migrate create -ext sql -dir migrations -seq $(NAME)