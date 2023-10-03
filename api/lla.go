// TEST purpose only
package api

import (
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Server) handleGetLla(c *gin.Context) {

	time.Sleep(3 * time.Second)

	c.JSON(200, gin.H{
		"message": "Welcome to LLA!",
	})
}
