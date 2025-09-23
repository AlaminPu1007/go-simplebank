package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/alaminpu1007/simplebank/api"
	db "github.com/alaminpu1007/simplebank/db/sqlc"
	"github.com/alaminpu1007/simplebank/gapi"
	pb "github.com/alaminpu1007/simplebank/pb"
	"github.com/alaminpu1007/simplebank/util"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	if err != nil {
		log.Fatal("Cannot create server:", err)
	}

	runGinServer(config, store)

	// runGrpcServer(config, store)

}

// This method will run grpc server
func runGrpcServer(config util.Config, store *db.Store) {

	server, err := gapi.NewServer(config, store)

	if err != nil {
		log.Fatal("Cannot create server:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)

	// optional (but recommended)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)

	if err != nil {
		log.Fatal("Cannot create listener:")
	}

	log.Printf("Start grpc server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("can not start grpc server")
	}
}

// This method will run gin server
func runGinServer(config util.Config, store *db.Store) {
	// using the store, create a new server
	server, err := api.NewServer(config, store)

	err = server.Start(config.HTTPServerAddress)

	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
