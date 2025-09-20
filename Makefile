DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

postgress:
# TO GENERATE A NEW POSTGRES CONTAINER
# 1. docker rm -f postgres12
# 2. make postgress
	docker run --name postgres12 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:18beta3-alpine

createdb:
# RUN POSTGRES SHELL THROUGH DOCKER
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up 1


migratedown:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down 1


startpostgress: 
	docker start postgres12

sqlc:
	sqlc generate

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

test:
	go test -v --cover ./...

server:
	go run main.go

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

evans:
	evans --host localhost --port 9000 -r repl


.PHONY: server createdb dropdb postgress migrateup migratedown startpostgress sqlc migratedown1 migrateup1 db_docs db_schema proto evans
