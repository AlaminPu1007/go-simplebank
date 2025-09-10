package api

import (
	"fmt"

	db "github.com/alaminpu1007/simplebank/db/sqlc"
	"github.com/alaminpu1007/simplebank/token"
	"github.com/alaminpu1007/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Serve http request for our banking service
type Server struct {
	config     util.Config
	store      *db.Store
	tokenMaker token.Maker
	router     *gin.Engine
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
		config:     config, // we will get token maker related info later
	}

	// binding our custom validation
	// can be used as a custom validation
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()

	return server, nil
}

// ALL INITIALIZED ROUTER WILL BE GOES HERE
func (server *Server) setupRouter() {

	router := gin.Default()

	/* USER ROUTE WILL BE GOES HERE */
	// create a user
	router.POST("/create-user", server.createUser)
	// login user route
	router.POST("/users/signin", server.loginUser)
	// get new access token
	router.POST("/tokens/renew_access", server.renewAccessToken)

	// create a router group for protected route or users
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// Create a account post method
	authRoutes.POST("/accounts", server.createAccount)
	// Get account by id
	authRoutes.GET("/accounts/:id", server.getAccountById)
	// get list of account with offset/limit
	authRoutes.GET("/accounts", server.getListAccounts)

	/* TRANSFER BALANCE ROUTE WILL BE GOES HERE */
	// create a transfer
	authRoutes.POST("/transfers", server.createTransfer)

	server.router = router
}

// START: runs the HTTP server on a specif address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// CREATE ERROR HANDLER TO SERVER ERROR JSON GLOBALLY
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
