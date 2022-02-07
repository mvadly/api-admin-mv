package middleware

import (
	"api-adminmv/util"

	"github.com/gin-gonic/gin"
)

func TokenValidator() gin.HandlerFunc {
	return func(c *gin.Context) {

		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(401, gin.H{"status": false, "message": "Authorization is invalid"})
			c.AbortWithStatus(401)
			return
		}

		lenBearer := len(BEARER_SCHEMA)
		tokenString := authHeader[lenBearer:]
		token, err := util.ValidateTokenString(tokenString)

		if err != nil {
			c.JSON(401, gin.H{"status": false, "message": err.Error()})
			c.AbortWithStatus(401)
			return
		}

		if !token.Valid {
			c.JSON(401, gin.H{"status": false, "message": err.Error()})
			c.AbortWithStatus(401)
			return
		}

		c.Next()
	}
}
