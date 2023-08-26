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
		TopicID: pgtype.Text{String: req.TopicID},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, entity.CreateFlashCardsFromLIs(lis))
}

type CompleteFlashCardsRequest struct {
	TopicIDs []string `json:"topic_ids" binding:"required"`
}

// func (s *Server) handleCompleteFlashcards(c *gin.Context) {
// 	var req []entity.FlashCard
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	lis := entity.CreateLIsFromFlashCards(req)
// 	if err := s.store.CompleteLearningItems(c, lis); err != nil {
// 		c.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	c.JSON(http.StatusOK, req)
// }
