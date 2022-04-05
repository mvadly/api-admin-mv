package middleware

import (
	"api-adminmv/util"
	"fmt"

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

func TokenValidatorDummy() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("req:", c.Request.Header)
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(401, gin.H{
				"responseCode":        "01",
				"responseDescription": "authorization is invalid",
			})
			c.AbortWithStatus(401)
			return
		}

		lenBearer := len(BEARER_SCHEMA)
		tokenString := authHeader[lenBearer:]

		if tokenString != "R04XSUbnm1GXNmDiXx9ysWMpFWBr" {
			c.JSON(401, gin.H{
				"responseCode":        "01",
				"responseDescription": "token is invalid",
			})

			c.AbortWithStatus(401)
			return
		}

		c.Next()
	}
}
