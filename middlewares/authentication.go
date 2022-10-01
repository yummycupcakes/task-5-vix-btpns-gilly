package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yummycupcakes/task-5-vix-btpns-gilly/helpers"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})

			return
		}

		c.Set("userData", verifyToken)
		c.Next()
	}
}