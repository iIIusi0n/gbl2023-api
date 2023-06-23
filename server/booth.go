package server

import (
	"gbl-api/controllers/booth"
	"gbl-api/controllers/score"
	"github.com/gin-gonic/gin"
	"log"
)

func makeBooth(c *gin.Context) {
	var booths []booth.Booth
	err := c.BindJSON(&booths)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	for _, b := range booths {
		err = booth.MakeBooth(b)
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{
				"message": "Internal server error",
			})
			return
		}
	}

	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func getBooths(c *gin.Context) {
	booths, err := booth.GetBooths()
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, booths)
}

func getBooth(c *gin.Context) {
	bid := c.Param("bid")
	b, err := booth.GetBooth(bid)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Booth not found",
		})
		return
	}
	c.JSON(200, b)
}

func checkBooth(c *gin.Context) {
	bid := c.Param("bid")
	uid := c.Param("uid")

	p, err := score.IsUserParticipated(bid, uid)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	if p {
		c.JSON(200, gin.H{
			"participate": true,
		})
	} else {
		c.JSON(200, gin.H{
			"participate": false,
		})
	}
}
