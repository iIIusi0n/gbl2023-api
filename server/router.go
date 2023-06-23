package server

import (
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()

	apiGroup := r.Group("/api") // "/api" prefix added

	auth := apiGroup.Group("/auth")
	{
		auth.POST("/login", authLogin)
		auth.POST("/register", authRegister)
	}

	booth := apiGroup.Group("/booth")
	{
		booth.GET("/", getBooths)
		booth.GET("/:bid", getBooth)
		booth.GET("/check/:bid/:uid", checkBooth)

		booth.POST("/make", makeBooth)
	}

	problem := apiGroup.Group("/problem")
	{
		problem.GET("/:bid", problemList)

		problem.POST("/submit/:bid", problemSubmit)
	}

	ranking := apiGroup.Group("/ranking")
	{
		ranking.GET("/", rankingList)
	}

	user := apiGroup.Group("/user")
	{
		user.GET("/:uid", userInfo)
	}

	return r
}
