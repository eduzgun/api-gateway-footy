package middlewares

import (
	"github.com/eduzgun/api-gateway-footy/utils"
	"github.com/gin-gonic/gin"
)

func isAuthorised() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")

		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorised"})
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(cookie)

		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorised"})
			c.Abort()
			return
		}

		c.Set("role", claims.Role)
		c.Next()

	}
}
