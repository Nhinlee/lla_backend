package api

import "github.com/gin-gonic/gin"

type GeneratePresignedURLRequest struct {
	FileName string `json:"file_name" binding:"required"`
}

type GeneratePresignedURLResponse struct {
	PresignedURL string `json:"presigned_url"`
	PublicURL    string `json:"public_url"`
}

func (s *Server) handleGeneratePresignedURL(c *gin.Context) {
	var req GeneratePresignedURLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, errorResponse(err))
		return
	}

	url, err := s.fileStore.GenerateResumableUploadURL(req.FileName)
	if err != nil {
		c.JSON(500, errorResponse(err))
		return
	}

	c.JSON(200, GeneratePresignedURLResponse{
		PresignedURL: url.String(),
		PublicURL:    s.fileStore.GeneratePublicObjectURL(req.FileName),
	})
}
