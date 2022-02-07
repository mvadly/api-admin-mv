package auth

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	// r.Use(util.BasicAuth)
	v1 := r.Group("/v1")
	{
		v1.POST("auth/login", Login)
	}

}
