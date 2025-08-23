postgress:
# TO GENERATE A NEW POSTGRES CONTAINER
# 1. docker rm -f postgres12
# 2. make postgress
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:18beta3-alpine

createdb:
# RUN POSTGRES SHELL THROUGH DOCKER
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

startpostgress: 
	docker start postgres12

sqlc:
	sqlc generate

test:
	go test -v --cover ./...

.PHONY: createdb dropdb postgress migrateup migratedown startpostgress sqlc
