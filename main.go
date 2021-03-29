package main

import (
	"github.com/gin-gonic/gin"

	"github.com/akkien/explorer-modern/util"
)

var router *gin.Engine

func main() {

	util.P("ChitChat", util.Config.Version, "started at", util.Config.Address)

	// Gin config
	gin.SetMode(gin.ReleaseMode)

	router = gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	initializeRoutes()

	router.Run(":8081")
}
