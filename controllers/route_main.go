package controllers

import (
	"github.com/gin-gonic/gin"
)

// Index controller
func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}
