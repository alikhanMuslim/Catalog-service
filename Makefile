postgres:
	docker run --name postgres12 -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=root -p 5432:5432 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root catalog_service

dropdb:
	docker exec -it postgres12 dropdb catalog_service

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/catalog_service?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/catalog_service?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb