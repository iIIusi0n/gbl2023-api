package server

import (
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api") // "/api" prefix added
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", authLogin)
			auth.POST("/register", authRegister)
		}

		booth := api.Group("/booth")
		{
			booth.GET("/", getBooths)
			booth.GET("/:bid", getBooth)
			booth.GET("/check/:bid/:uid", checkBooth)

			booth.POST("/make", makeBooth)
		}

		problem := api.Group("/problem")
		{
			problem.GET("/:bid", problemList)

			problem.POST("/submit/:bid", problemSubmit)
		}

		ranking := api.Group("/ranking")
		{
			ranking.GET("/", rankingList)
		}

		user := api.Group("/user")
		{
			user.GET("/:uid", userInfo)
		}
	}

	return r
}
