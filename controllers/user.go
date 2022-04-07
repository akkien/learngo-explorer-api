package controllers

import (
	"net/http"

	"github.com/akkien/learngo-explorer-api/util"
	"github.com/gin-gonic/gin"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Auth(c *gin.Context) {
	var userCred user

	if err := c.BindJSON(&userCred); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    403,
			"message": "Hello World!",
			"data":    "",
		})
	}

	token, err := util.GenerateToken(userCred.Username, userCred.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Success",
		"data":    token,
	})
}
