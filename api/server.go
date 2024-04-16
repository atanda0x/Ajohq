package api

import (
	db "github.com/atanda0x/FintechConnect/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP request for fintch services
type Server struct {
	store  db.Store
	router *gin.Engine
}

// Newserver create a new HTTP server and setup routing
func NewServer(s db.Store) *Server {
	server := &Server{store: s}
	router := gin.Default()

	router.POST("/accounts", server.CreateAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.PUT("/accounts", server.updateAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)

	server.router = router
	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
