package auth

import (
	"api-adminmv/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Login(c *gin.Context) {
	var req RequestLogin
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"status": false, "message": "Validation errors", "errors": util.GetErrorJSON(err)})
		return
	}

	data, err := GetUserPasswordLogin(req)

	if err != nil {
		c.JSON(500, gin.H{"status": false, "message": "Internal Server Error", "errors": err.Error()})
		return
	}

	check := util.CheckPasswordHash(data.Username+req.Password, data.Password)
	if !check {
		c.JSON(401, gin.H{"status": false, "message": "Username / Password not match"})
		return
	}

	var responseData = util.ResponseTokenCreated{
		ID:       uuid.New().ID(),
		Username: data.Username,
		Fullname: data.Fullname,
	}
	token, err := util.GenerateToken(responseData)

	if err != nil {
		c.JSON(500, gin.H{"status": false, "message": "Failed to created the token"})
		return
	}

	// util.SaveCookie(token, c)
	c.SetCookie("cookie-adminmv", token, int(time.Now().Add(time.Minute*1).Unix()), "/", "http://localhost:1000", true, true)

	c.JSON(200, gin.H{
		"status":  false,
		"message": "Login Success",
		"token":   token,
	})

}
