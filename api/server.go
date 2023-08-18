package api

import (
	db "lla/db/sqlc"

	fs "lla/golibs/file_store"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router    *gin.Engine
	store     db.Store
	fileStore fs.FileStore
}

func NewServer(store db.Store, filestore fs.FileStore) (*Server, error) {
	server := &Server{
		store:     store,
		fileStore: filestore,
	}

	server.SetupRouter()

	return server, nil
}

func (s *Server) SetupRouter() {
	router := gin.Default()

	// TEST purpose only
	router.GET("/lla", s.handleGetLla)

	// Common
	router.POST("/generate_resumable_upload_url", s.handleGeneratePresignedURL)
	// router.POST("/delete_object", s.handleDeleteObject)

	// Learning items
	router.POST("/learning-items", s.handleCreateLearningItem)
	router.GET("/learning_items", s.handleGetLearningItems)
	router.DELETE("/learning_items/:id", s.handleDeleteLearningItem)
	router.PUT("/learning-items", s.handleUpdateLearningItem)

	// Topics
	router.POST("/topics", s.handleUpsertTopic)
	router.GET("/topics", s.handleGetTopics)
	router.GET("/topics/with_learning_items_count", s.handleGetTopicsAndTotalLearningItems)

	// Flashcards
	router.POST("/flashcards/start_learning", s.handleStartLearningFlashcards)

	s.router = router
}

// Run http server on address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
