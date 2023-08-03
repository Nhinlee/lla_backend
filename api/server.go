package api

import "github.com/gin-gonic/gin"

type Server struct {
	router *gin.Engine
}

func NewServer() (*Server, error) {
	server := &Server{}

	server.SetupRouter()

	return server, nil
}

func (s *Server) SetupRouter() {
	router := gin.Default()

	// TEST purpose only
	router.GET("/lla", s.handleGetLla)

	// Test auto ci/cd
	router.GET("/test", s.handleGetLla)

	s.router = router
}

// Run http server on address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
