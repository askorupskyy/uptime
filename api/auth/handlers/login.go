package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rcbxd/uptime/api/auth/types"
	"github.com/rcbxd/uptime/api/auth/utils"
)

func LoginHandler(c *gin.Context) {
	utils.NewLogger().Warn("login handler is not fully implemented!!")

	var loginDTO types.LoginDTO

	c.BindJSON(&loginDTO)

	c.JSON(200, gin.H{
		"message": "logging in...",
		"user":    loginDTO,
	})
}
