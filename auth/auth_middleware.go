package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(tokenIssuer TokenIssuer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			err := fmt.Errorf("Authorization header is missing")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse(err))
			return
		}

		payload, err := tokenIssuer.VerifyToken(authHeader)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse(err))
			return
		}

		ctx.Set("payload", payload)
		ctx.Next()
	}
}

func ErrorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
