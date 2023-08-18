package api

import (
	"database/sql"
	"lla/api/entity"
	db "lla/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
)

type CreateLearningItemRequest struct {
	ImageLink        string   `json:"image_link" binding:"required"`
	EnglishWord      string   `json:"english_word" binding:"required"`
	VietnameseWord   string   `json:"vietnamese_word" binding:"required"`
	EnglishSentences []string `json:"english_sentences"`
	TopicID          string   `json:"topic_id"`
}

func (s *Server) handleUpsertLearningItem(c *gin.Context) {
	var req CreateLearningItemRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	learningItem, err := s.store.CreateLearningItem(c, db.CreateLearningItemParams{
		ID:               ulid.Make().String(),
		ImageLink:        req.ImageLink,
		EnglishWord:      req.EnglishWord,
		VietnameseWord:   sql.NullString{String: req.VietnameseWord, Valid: req.VietnameseWord != ""},
		EnglishSentences: req.EnglishSentences,
		TopicID:          sql.NullString{String: req.TopicID, Valid: req.TopicID != ""},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(200, gin.H{
		"id": learningItem.ID,
	})
}

// type LearningItem struct {
// 	ID               string   `json:"id"`
// 	ImageLink        string   `json:"image_link"`
// 	EnglishWord      string   `json:"english_word"`
// 	VietnameseWord   string   `json:"vietnamese_word"`
// 	EnglishSentences []string `json:"english_sentences"`
// 	CreatedAt        string   `json:"created_at"`
// 	UserID           string   `json:"user_id"`
// 	TopicID          string   `json:"topic_id"`
// }

func (s *Server) handleGetLearningItems(c *gin.Context) {
	learningItems, err := s.store.GetAllLearningItems(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(200, entity.CreateLearningItemsFromLIs(learningItems))
}

func (s *Server) handleDeleteLearningItem(c *gin.Context) {
	id := c.Param("id")

	item, err := s.store.DeleteLearningItem(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(200, gin.H{
		"id": item.ID,
	})
}
