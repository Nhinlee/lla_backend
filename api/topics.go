package api

import (
	db "lla/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
)

type CreateTopicRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateTopicResponse struct {
	ID string `json:"id"`
}

func (s *Server) handleUpsertTopic(ctx *gin.Context) {
	var req CreateTopicRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, errorResponse(err))
		return
	}

	topic, err := s.store.CreateTopic(ctx, db.CreateTopicParams{
		ID:   ulid.Make().String(),
		Name: req.Name,
	})
	if err != nil {
		ctx.JSON(500, errorResponse(err))
		return
	}

	ctx.JSON(200, CreateTopicResponse{
		ID: topic.ID,
	})
}

func (s *Server) handleGetTopics(ctx *gin.Context) {
	topics, err := s.store.GetAllTopics(ctx)
	if err != nil {
		ctx.JSON(500, errorResponse(err))
		return
	}

	ctx.JSON(200, topics)
}

func (s *Server) handleGetTopicsAndTotalLearningItems(ctx *gin.Context) {
	topics, err := s.store.GetTopicsAndTotalLearningItems(ctx)
	if err != nil {
		ctx.JSON(500, errorResponse(err))
		return
	}

	ctx.JSON(200, topics)
}
