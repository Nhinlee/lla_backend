package api

import (
	"database/sql"
	"lla/api/entity"
	db "lla/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetFlashCardRequest struct {
	TopicID string `json:"topic_id" binding:"required"`
	Limit   int32  `json:"limit" binding:"required"`
}

func (s *Server) handleStartLearningFlashcards(c *gin.Context) {
	var req GetFlashCardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	lis, err := s.store.GetLearningItemsByTopicAndCompleted(c, db.GetLearningItemsByTopicAndCompletedParams{
		Limit:   req.Limit,
		TopicID: sql.NullString{String: req.TopicID, Valid: req.TopicID != ""},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, entity.CreateFlashCardsFromLIs(lis))
}
