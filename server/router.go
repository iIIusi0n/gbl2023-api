package server

import (
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/login", authLogin)
		auth.POST("/register", authRegister)
	}

	booth := r.Group("/booth")
	{
		booth.GET("/", getBooths)
		booth.GET("/:bid", getBooth)
		booth.GET("/check/:bid/:uid", checkBooth)

		booth.POST("/make", makeBooth)
	}

	problem := r.Group("/problem")
	{
		problem.GET("/:bid", problemList)

		problem.POST("/submit/:bid", problemSubmit)
	}

	ranking := r.Group("/ranking")
	{
		ranking.GET("/", rankingList)
	}

	user := r.Group("/user")
	{
		user.GET("/:uid", userInfo)
	}

	return r
}
