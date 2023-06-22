package server

import (
	"gbl-api/controllers/user"
	"github.com/gin-gonic/gin"
)

func AuthLogin(c *gin.Context) {
	var u user.User
	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	if user.IsUserExist(u.UID) {
		c.JSON(200, gin.H{
			"message": "Login success",
		})
	} else {
		c.JSON(404, gin.H{
			"message": "User not found",
		})
	}
}

func AuthRegister(c *gin.Context) {
	var u user.User
	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	err = user.RegisterUser(u)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Register success",
	})
}
