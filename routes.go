package main

import (
	"github.com/akkien/learngo-explorer-api/controllers"
	"github.com/akkien/learngo-explorer-api/middlewares"
)

func (app *application) routes() {
	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not

	api := app.router.Group("/api")

	// Handle the index route
	api.GET("", controllers.Index)

	userRoutes := api.Group("/users")
	{
		userRoutes.POST("/auth", controllers.Auth)
	}

	blockRoutes := api.Group("/blocks")
	{
		blockRoutes.GET("", controllers.GetBlocks)
		blockRoutes.GET("/:number", controllers.GetBlock)
		blockRoutes.GET("/:number/txs", controllers.GetTransactionsInBlock)
	}

	txRoutes := api.Group("/txs")
	{
		txRoutes.GET("", controllers.GetTransactions)
		txRoutes.GET("/:hash", controllers.GetTransaction)
	}

	albumRoutes := api.Group("/albums")
	{
		albumRoutes.POST("", controllers.PostAlbums)
		albumRoutes.GET("", controllers.GetAlbums)
		albumRoutes.GET("/:id", middlewares.ValidateToken(), controllers.GetAlbumByID)
		albumRoutes.DELETE("/:id", controllers.DeleteAlbumByID)
	}
}
