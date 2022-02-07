package util

import (
	config "api-adminmv/config"
	entity "api-adminmv/entity"
	"fmt"

	"github.com/gin-gonic/gin"
)

var sql = config.ConnectDB().Debug()

func GetAuth(c *gin.Context) string {
	authorization := fmt.Sprintf("%v", c.Request.Header["Authorization"])
	return authorization
}

func GetUserBasicAuth(c *gin.Context) (username string, password string, ok bool) {
	return c.Request.BasicAuth()
}

func BasicAuth(c *gin.Context) {
	var result entity.UserBasicAuth
	authorization := GetAuth(c)
	if len(authorization) >= 6 && authorization[1:6] == "Basic" {
		username, password, hasAuth := GetUserBasicAuth(c)
		check := sql.Select("username, password").
			Where("username = ? AND password = ? AND is_active = 1", username, SaltMD5(username, password)).
			Limit(1).
			Find(&result)

		config.CloseDB()

		if hasAuth && check.Error != nil {
			c.JSON(401, gin.H{"status": false, "message": "Unauthorized"})
			c.Abort()
		}

	} else {
		c.JSON(401, gin.H{"status": false, "message": "Unauthorized"})
		c.Abort()
	}

	c.Next()
}

func JWTAuth() {

}
