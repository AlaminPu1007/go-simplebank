package api

import (
	db "github.com/alaminpu1007/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Serve http request for our banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// Create a account post method
	router.POST("/accounts", server.createAccount)

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
