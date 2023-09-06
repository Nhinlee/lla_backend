package api

import (
	"lla/api/entity"
	db "lla/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/oklog/ulid/v2"
)

type CreateLearningItemRequest struct {
	ImageLink        string   `json:"image_link" binding:"required"`
	EnglishWord      string   `json:"english_word" binding:"required"`
	VietnameseWord   string   `json:"vietnamese_word"`
	EnglishSentences []string `json:"english_sentences"`
	TopicID          string   `json:"topic_id"`
}

func (s *Server) handleCreateLearningItem(c *gin.Context) {
	var req CreateLearningItemRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	learningItem, err := s.store.CreateLearningItem(c, db.CreateLearningItemParams{
		ID:               ulid.Make().String(),
		ImageLink:        req.ImageLink,
		EnglishWord:      req.EnglishWord,
		VietnameseWord:   pgtype.Text{String: req.VietnameseWord, Valid: req.VietnameseWord != ""},
		EnglishSentences: req.EnglishSentences,
		TopicID:          pgtype.Text{String: req.TopicID, Valid: req.TopicID != ""},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(200, gin.H{
		"id": learningItem.ID,
	})
}

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

type UpdateLearningItemRequest struct {
	ID               string   `json:"id" binding:"required"`
	ImageLink        string   `json:"image_link"`
	EnglishWord      string   `json:"english_word"`
	VietnameseWord   string   `json:"vietnamese_word"`
	EnglishSentences []string `json:"english_sentences"`
	TopicID          string   `json:"topic_id"`
}

func (s *Server) handleUpdateLearningItem(c *gin.Context) {
	var req UpdateLearningItemRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	updatedItem, err := s.store.GetLearningItemById(c, req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// apply updates
	if req.ImageLink != "" {
		updatedItem.ImageLink = req.ImageLink
	}

	if req.EnglishWord != "" {
		updatedItem.EnglishWord = req.EnglishWord
	}

	if req.VietnameseWord != "" {
		updatedItem.VietnameseWord = pgtype.Text{String: req.VietnameseWord, Valid: req.VietnameseWord != ""}
	}

	if len(req.EnglishSentences) > 0 {
		updatedItem.EnglishSentences = req.EnglishSentences
	}

	if req.TopicID != "" {
		updatedItem.TopicID = pgtype.Text{String: req.TopicID, Valid: req.TopicID != ""}
	}

	err = s.store.UpdateLearningItem(c, db.UpdateLearningItemParams{
		ID:               updatedItem.ID,
		ImageLink:        updatedItem.ImageLink,
		EnglishWord:      updatedItem.EnglishWord,
		VietnameseWord:   updatedItem.VietnameseWord,
		EnglishSentences: updatedItem.EnglishSentences,
		TopicID:          updatedItem.TopicID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(200, gin.H{})
}
