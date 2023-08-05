package api

import (
	db "lla/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	store  db.Store
}

func NewServer(store db.Store) (*Server, error) {
	server := &Server{
		store: store,
	}

	server.SetupRouter()

	return server, nil
}

func (s *Server) SetupRouter() {
	router := gin.Default()

	// TEST purpose only
	router.GET("/lla", s.handleGetLla)

	// Learning items
	router.POST("/learning-items", s.handleUpsertLearningItem)

	s.router = router
}

// Run http server on address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
