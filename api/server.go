package api

import (
	"fmt"
	"lla/auth"
	db "lla/db/sqlc"

	fs "lla/golibs/file_store"
	"lla/golibs/open_ai"
	"lla/golibs/vision_ai"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router      *gin.Engine
	store       db.Store
	fileStore   fs.FileStore
	visionAI    *vision_ai.VisionAI
	openAI      *open_ai.OpenAI
	tokenIssuer auth.TokenIssuer
}

func NewServer(store db.Store, filestore fs.FileStore, visionAI *vision_ai.VisionAI, openAI *open_ai.OpenAI) (*Server, error) {

	// TODO: move secret key to secret manager
	issuer, err := auth.NewPasetoTokenIssuer("12345678901234567890123456789012")
	if err != nil {
		return nil, fmt.Errorf("cannot create token issuer: %w", err)
	}

	server := &Server{
		store:       store,
		fileStore:   filestore,
		tokenIssuer: issuer,
		visionAI:    visionAI,
		openAI:      openAI,
	}

	server.SetupRouter()

	return server, nil
}

func (s *Server) SetupRouter() {
	router := gin.Default()

	// Public
	router.POST("/login", s.handleLogin)
	router.POST("/signup", s.handleSignUp)

	// Middleware
	router.Use(auth.AuthMiddleware(s.tokenIssuer))

	// TEST purpose only
	router.GET("/lla", s.handleGetLla)

	// Common
	router.POST("/generate_resumable_upload_url", s.handleGeneratePresignedURL)
	// router.POST("/delete_object", s.handleDeleteObject)

	// Learning items
	router.POST("/learning_items", s.handleCreateLearningItem)
	router.GET("/learning_items", s.handleGetLearningItems)
	router.DELETE("/learning_items/:id", s.handleDeleteLearningItem)
	router.PUT("/learning_items", s.handleUpdateLearningItem)
	router.POST("/learning_items/generate_label", s.handleGenerateLearningItemLabel)
	router.POST("/learning_items/generate_sentences", s.handleGenerateEnSentences)

	// Topics
	router.POST("/topics", s.handleUpsertTopic)
	router.GET("/topics", s.handleGetTopics)
	router.GET("/topics/with_learning_items_count", s.handleGetTopicsAndTotalLearningItems)

	// Flashcards
	router.POST("/flashcards/learning", s.handleStartLearningFlashcards)
	router.POST("/flashcards/complete", s.handleCompleteFlashcards)

	// Users
	router.GET("/user/profile/:id", s.handleGetUserProfile)

	s.router = router
}

// Run http server on address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
