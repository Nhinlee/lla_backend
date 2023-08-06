package api

import (
	db "lla/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
)

type CreateLearningItemRequest struct {
	ImageUrl         string   `json:"image_url" binding:"required"`
	EnglishWord      string   `json:"english_word" binding:"required"`
	VietnameseWord   string   `json:"vietnamese_word" binding:"required"`
	EnglishSentences []string `json:"english_sentences"`
}

func (s *Server) handleUpsertLearningItem(c *gin.Context) {
	var req CreateLearningItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	learningItem, err := s.store.CreateLearningItem(c, db.CreateLearningItemParams{
		ID:               ulid.Make().String(),
		ImageLink:        req.ImageUrl,
		EnglishWord:      req.EnglishWord,
		VietnameseWord:   req.VietnameseWord,
		EnglishSentences: req.EnglishSentences,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(200, gin.H{
		"id": learningItem.ID,
	})
}
