package server

import (
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/login", AuthLogin)
		auth.POST("/register", AuthRegister)
	}

	booth := r.Group("/booth")
	{
		booth.GET("/", GetBooths)
		booth.GET("/:bid", GetBooth)
		booth.GET("/check/:bid/:uid", CheckBooth)

		booth.POST("/make", MakeBooth)
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

	user := r.Group("/user")
	{
		user.GET(":/uid")
	}

	return r
}
