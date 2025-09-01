package api

import (
	db "github.com/alaminpu1007/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Serve http request for our banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// binding our custom validation
	// can be used as a custom validation
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	// Create a account post method
	router.POST("/accounts", server.createAccount)
	// Get account by id
	router.GET("/accounts/:id", server.getAccountById)
	// get list of account with offset/limit
	router.GET("/accounts", server.getListAccounts)

	/* TRANSFER BALANCE ROUTE WILL BE GOES HERE */
	// create a transfer
	router.POST("/transfers", server.createTransfer)

	/* USER ROUTE WILL BE GOES HERE */
	// create a user
	router.POST("/create-user", server.createUser)

	server.router = router
	return server
}

// START: runs the HTTP server on a specif address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// CREATE ERROR HANDLER TO SERVER ERROR JSON GLOBALLY
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
