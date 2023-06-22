package server

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.GET("/login")

		auth.POST("/register")
	}

	booth := r.Group("/booth")
	{
		booth.GET("/")
		booth.GET("/:bid")
		booth.GET("/check/:bid/:uid")

		booth.POST("/make")
	}

	problem := r.Group("/problem")
	{
		problem.GET("/:bid")

		problem.POST("/submit/:bid")
	}

	ranking := r.Group("/ranking")
	{
		ranking.GET("/")
	}

	return r
}
