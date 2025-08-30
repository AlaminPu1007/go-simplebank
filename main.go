package main

import (
	"database/sql"
	"log"

	"github.com/alaminpu1007/simplebank/api"
	db "github.com/alaminpu1007/simplebank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8000"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("DB connection is not possible")
	}

	// now create a new store
	store := db.NewStore(conn)

	// using the store, create a new server
	server := api.NewServer(store)

	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
