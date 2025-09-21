package gapi

import (
	"fmt"

	db "github.com/alaminpu1007/simplebank/db/sqlc"
	pb "github.com/alaminpu1007/simplebank/pb"
	"github.com/alaminpu1007/simplebank/token"
	"github.com/alaminpu1007/simplebank/util"
)

// Serve http request for our banking service
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      *db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {

	// create a token maker
	// if you want to use JWT token maker, just replace the method with: token.NewJWTMaker()

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)

	if err != nil {
		return nil, fmt.Errorf("can not create token maker: %w", err)
	}

	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
	}

	return server, nil
}
