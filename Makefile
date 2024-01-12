dropdb:
	migrate -path internal/repository/migrate/ -database "postgres://root:12345@localhost:5432/monopoly?sslmode=disable" down
updb:
	migrate -path internal/repository/migrate/ -database "postgres://root:12345@localhost:5432/monopoly?sslmode=disable" up
sqlup:
	docker compose up
run:
	go run cmd/main.go
.PHONY: updb dropdb sqlup run