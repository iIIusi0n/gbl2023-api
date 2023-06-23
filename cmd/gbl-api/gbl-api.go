package main

import (
	"fmt"
	"gbl-api/config"
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

	r := gin.Default()
	r.Run(config.Hostname + ":" + config.Port)
}
