package server

import (
	"fmt"
	"gbl-api/controllers/booth"
	"gbl-api/controllers/score"
	"log"

	"github.com/gin-gonic/gin"
)

func makeBooth(c *gin.Context) {
	var b booth.Booth
	err := c.BindJSON(&b)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	err = booth.MakeBooth(b)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
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
	c.JSON(200, gin.H{
		"boothlist": booths,
	})
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

func deleteBooth(c *gin.Context) {
	bid := c.Param("bid")

	err := booth.DeleteBooth(bid)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func deleteBoothUser(c *gin.Context) {
	bid := c.Param("bid")

	err1 := booth.DeleteBooth(bid)
	err2 := booth.DeletePasswordByBID(bid)
	if err1 != nil || err2 != nil {
		log.Println(err1, err2)
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
}

func makeBoothUser(c *gin.Context) {
	type Request struct {
		Password string `json:"password"`
	}

	var req Request
	err := c.BindJSON(&req)
	fmt.Println(err)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	err = booth.AddPassword(req.Password)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "ok",
	})
}
