package middlewares

import (
	"eliarms.events.com/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}
	userId, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}
	c.Set("userId", userId)
	c.Next()
}
