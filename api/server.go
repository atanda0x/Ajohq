package api

import (
	"fmt"

	db "github.com/atanda0x/Ajohq/db/sqlc"
	"github.com/atanda0x/Ajohq/token"
	"github.com/atanda0x/Ajohq/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serve HTTP requests for the banking services
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
	}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// Accounts router
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccount)

	// Transfer router
	authRoutes.POST("/transfers", server.createTransfer)

	// User router
	router.POST("/users", server.CreateUser)
	router.POST("?users/login", server.loginUser)

	server.router = router
	return server, nil
}

// Start runs server HTTP on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
