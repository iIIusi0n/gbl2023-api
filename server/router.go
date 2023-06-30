package server

import (
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/getfile", "./upload")

	api := r.Group("/api") // "/api" prefix added
	{
		api.POST("/upload", uploadFile)

		auth := api.Group("/auth")
		{
			auth.POST("/login", authLogin)
			auth.POST("/register", authRegister)
			auth.POST("/boothadmin", authBoothAdmin)
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

			problem.POST("/make/:bid", problemMake)
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

		notification := api.Group("/notification")
		{
			notification.GET("/", notificationList)

			notification.POST("/make", notificationMake)
		}
	}

	return r
}
