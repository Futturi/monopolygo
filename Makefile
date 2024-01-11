dropdb:
	migrate -path internal/repository/migrate/ -database "postgres://root:12345@localhost:5432/monopoly?sslmode=disable" down
updb:
	migrate -path internal/repository/migrate/ -database "postgres://root:12345@localhost:5432/monopoly?sslmode=disable" up
sqlup:
	docker-compose up

.PHONY: updb dropdb sqlup