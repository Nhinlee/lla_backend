// TEST purpose only
package api

import "github.com/gin-gonic/gin"

func (s *Server) handleGetLla(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to LLA @@!",
	})
}
