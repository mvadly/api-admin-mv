package user

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {

	v1 := r.Group("v1/")
	{
		v1.GET("list_users", GetUsersBy)
		v1.GET("users", GetUsers)
		v1.POST("users", CreateUser)
		v1.POST("users/:id", UpdateUser)
		v1.DELETE("users/:id", DeleteUser)
	}
}
