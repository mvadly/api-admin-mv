package register

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	v1 := r.Group("v1/register")
	{
		v1.POST("get_valid_giro", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": true, "message": "Norek Giro is valid"})
		})
	}
}
