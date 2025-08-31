package main

import (
	"database/sql"
	"log"

	"github.com/alaminpu1007/simplebank/api"
	db "github.com/alaminpu1007/simplebank/db/sqlc"
	"github.com/alaminpu1007/simplebank/util"
	_ "github.com/lib/pq"
)

// const (
// 	dbDriver      = "postgres"
// 	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
// 	serverAddress = "0.0.0.0:8000"
// )

func main() {

	// load from app.env
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("DB connection is not possible")
	}

	// now create a new store
	store := db.NewStore(conn)

	// using the store, create a new server
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
