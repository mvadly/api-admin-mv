package main

import (
	config "api-adminmv/config"
	middleware "api-adminmv/middleware"
	"api-adminmv/v1/auth"
	"api-adminmv/v1/product"
	"api-adminmv/v1/register"
	"api-adminmv/v1/user"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	routes := gin.Default()
	routes.Use(middleware.CORSMiddleware())
	auth.Router(routes)
	register.Router(routes)
	product.Router(routes)

	routes.Use(middleware.TokenValidator())
	user.Router(routes)

	routes.Run(":1000")
}
