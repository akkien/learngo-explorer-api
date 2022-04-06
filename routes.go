package main

import (
	"github.com/akkien/learngo-explorer-api/controllers"
	"github.com/akkien/learngo-explorer-api/middlewares"
)

func initializeRoutes() {
	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(middlewares.SetUserStatus())

	// Handle the index route
	router.GET("/", controllers.Index)
	router.GET("/login", controllers.LoginPage)
	router.GET("/signup", controllers.SignupPage)
	router.GET("/signup", controllers.Signup)

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

	// // Group article related routes together
	// articleRoutes := router.Group("/article")
	// {
	// 	// Handle GET requests at /article/view/some_article_id
	// 	articleRoutes.GET("/view/:article_id", getArticle)
	// }
}

// // index
// mux.HandleFunc("/", route.Index)
// // error
// mux.HandleFunc("/err", route.Err)

// // defined in route_auth.go
// mux.HandleFunc("/login", route.Login)
// mux.HandleFunc("/logout", route.Logout)
// mux.HandleFunc("/signup", route.Signup)
// mux.HandleFunc("/signup_account", route.SignupAccount)
// mux.HandleFunc("/authenticate", route.Authenticate)

// // defined in route_thread.go
// mux.HandleFunc("/blocks", route.ReadBlocks)
// mux.HandleFunc("/block/", route.ReadBlock)
// mux.HandleFunc("/txs", route.ReadTransactions)
// mux.HandleFunc("/tx/", route.ReadTransaction)

// // starting up the server
// server := &http.Server{
// 	Addr:           util.Config.Address,
// 	Handler:        mux,
// 	ReadTimeout:    time.Duration(util.Config.ReadTimeout * int64(time.Second)),
// 	WriteTimeout:   time.Duration(util.Config.WriteTimeout * int64(time.Second)),
// 	MaxHeaderBytes: 1 << 20,
// }
// server.ListenAndServe()
