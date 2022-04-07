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
	api.GET("/", controllers.Index)

	userRoutes := api.Group("/users")
	{
		userRoutes.POST("/auth", controllers.Auth)
	}

	blockRoutes := api.Group("/blocks")
	{
		blockRoutes.GET("/", controllers.GetBlocks)
		blockRoutes.GET("/:number", controllers.GetBlock)
		blockRoutes.GET("/:number/txs", controllers.GetTransactionsInBlock)
	}

	txRoutes := api.Group("/txs")
	{
		txRoutes.GET("/", controllers.GetTransactions)
		txRoutes.GET("/:hash", controllers.GetTransaction)
	}

	albumRoutes := api.Group("/albums")
	{
		albumRoutes.POST("/", controllers.PostAlbums)
		albumRoutes.GET("/", controllers.GetAlbums)
		albumRoutes.GET("/:id", middlewares.ValidateToken(), controllers.GetAlbumByID)
		albumRoutes.DELETE("/:id", controllers.DeleteAlbumByID)
	}
	// app.router.GET("/login", controllers.LoginPage)
	// app.router.GET("/signup", controllers.SignupPage)
	// app.router.GET("/signup", controllers.Signup)

	// // Group user related routes together
	// userRoutes := router.Group("/u")
	// {
	// 	// Handle the GET requests at /u/login
	// 	// Show the login page
	// 	// Ensure that the user is not logged in by using the middleware
	// 	userRoutes.GET("/login", EnsureNotLoggedIn(), showLoginPage)

	// 	// Handle POST requests at /u/login
	// 	// Ensure that the user is not logged in by using the middleware
	// 	userRoutes.POST("/login", EnsureNotLoggedIn(), performLogin)

	// 	// Handle GET requests at /u/logout
	// 	// Ensure that the user is logged in by using the middleware
	// 	userRoutes.GET("/logout", EnsureLoggedIn(), logout)

	// 	// Handle the GET requests at /u/register
	// 	// Show the registration page
	// 	// Ensure that the user is not logged in by using the middleware
	// 	userRoutes.GET("/register", EnsureNotLoggedIn(), showRegistrationPage)

	// 	// Handle POST requests at /u/register
	// 	// Ensure that the user is not logged in by using the middleware
	// 	userRoutes.POST("/register", EnsureNotLoggedIn(), register)
	// }
}
