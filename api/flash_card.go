package api

import (
	"lla/api/entity"
	db "lla/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
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
		TopicID: pgtype.Text{String: req.TopicID, Valid: true},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, entity.CreateFlashCardsFromLIs(lis))
}

type CompleteFlashCardsRequest struct {
	IDs []string `json:"ids" binding:"required"`
}

func (s *Server) handleCompleteFlashcards(c *gin.Context) {
	var req CompleteFlashCardsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res := s.store.UpdateCompletedAt(c, req.IDs)
	res.Exec(func(i int, err error) {
		if err != nil {
			c.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
