package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gbl-api/config"
	"gbl-api/controllers/booth"
	"gbl-api/controllers/notification"
	"gbl-api/controllers/problem"
	"gbl-api/controllers/score"
	"gbl-api/controllers/user"
	"gbl-api/data"
	"gbl-api/server"
	"github.com/gin-gonic/gin"
)

func main() {
	if config.DebugMode {
		fmt.Println("DEBUG_MODE is enabled.")
		gin.SetMode(gin.DebugMode)
	} else {
		fmt.Println("DEBUG_MODE is disabled.")
		gin.SetMode(gin.ReleaseMode)
	}

	serverLog, err := os.Create(fmt.Sprintf("server-%s.log", time.Now().Format("2006-01-02-15-04-05")))
	if err != nil {
		panic(err)
	}
	defer serverLog.Close()

	gin.DefaultWriter = serverLog

	errorLog, err := os.Create(fmt.Sprintf("error-%s.log", time.Now().Format("2006-01-02-15-04-05")))
	if err != nil {
		panic(err)
	}
	defer errorLog.Close()

	gin.DefaultErrorWriter = errorLog
	log.SetOutput(errorLog)

	score.UpdateLastScoreChanged()

	db := data.GetDatabase()
	db.AutoMigrate(&booth.Booth{})
	db.AutoMigrate(&booth.BoothPassword{})
	db.AutoMigrate(&score.Participation{})
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&problem.Problem{})
	db.AutoMigrate(&notification.Notification{})

	r := server.CreateRouter()
	r.Run(config.Hostname + ":" + config.Port)
}
