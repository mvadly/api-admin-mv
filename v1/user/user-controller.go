package user

import (
	"api-adminmv/util"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func GetUsersBy(c *gin.Context) {

	count, err := CountListUserBy(c)

	if err != nil {
		c.JSON(500, gin.H{"status": false, "message": err})
		return
	}

	data, err := ListUserBy(c)
	if err != nil {
		c.JSON(500, gin.H{"status": false, "message": err})
		return
	}
	var result = make(map[string]interface{})
	result["page"] = c.Query("page")
	result["per_page"] = c.Query("per_page")
	result["total"] = count
	pp, _ := strconv.ParseInt(c.Query("per_page"), 10, 64)
	ceil := count / pp
	result["total_pages"] = ceil
	result["data"] = data
	c.JSON(200, gin.H{"status": true, "message": "OK", "results": result})
}

func GetUsers(c *gin.Context) {
	data, err := ListUser()
	if err != nil {
		c.JSON(500, gin.H{"status": false, "message": err})
		return
	}

	c.JSON(200, gin.H{"status": true, "message": "OK", "result": data})
}

func CreateUser(c *gin.Context) {
	validate = validator.New()
	var req CreateReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"status": false, "message": "Validation errors", "errors": util.GetErrorJSON(err)})
		return
	}
	err = validate.Struct(req)
	if err != nil {
		c.JSON(400, gin.H{"status": false, "message": "Validation errors", "errors": util.GetError(err)})
		return
	}

	err = AddUser(req)
	if err != nil {
		c.JSON(500, gin.H{"status": false, "message": "Validation errors", "errors": err})
		return
	}

	c.JSON(200, gin.H{"status": true, "message": "User created"})

}

func UpdateUser(c *gin.Context) {
	var id = c.Param("id")
	var req CreateReq
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(400, gin.H{"status": false, "message": "Validation errors", "errors": util.GetErrorJSON(err)})
		return
	}

	err = UserUpdate(id, req)
	if err != nil {
		c.JSON(500, gin.H{"status": false, "message": "Internal Server error: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": true, "message": "User Updated"})

}

func DeleteUser(c *gin.Context) {
	var id = c.Param("id")
	err := UserDelete(id)
	if err != nil {
		c.JSON(500, gin.H{"status": false, "message": "Internal Server error: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": true, "message": "User Deleted"})

}
