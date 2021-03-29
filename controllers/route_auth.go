package controllers

import (
	"fmt"
	"net/http"

	"github.com/akkien/explorer-modern/models"
	"github.com/akkien/explorer-modern/util"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

// GET /login
// Login show the login page
func LoginPage(c *gin.Context) {
	res := gin.H{
		"title": "Login",
	}
	render(c, res, "login.html")
}

// // GET /signup
// // Signup show the signup page
func SignupPage(c *gin.Context) {
	res := gin.H{
		"title": "Signup",
	}
	render(c, res, "signup.html")
}

// POST /signup
// SignupAccount create the user account
func Signup(c *gin.Context) {
	user := models.User{
		Name:     c.PostForm("name"),
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}

	val := validator.New()
	err := val.Struct(user)
	if err != nil {
		responseBody := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			responseBody[e.Field()] = fmt.Sprint(e)
		}
		res := gin.H{
			"title":   "Signup",
			"payload": responseBody,
		}
		render(c, res, "signup.html")
		return
	}

	user.Password, err = util.HashPassword(user.Password)
	if err != nil {
		util.Danger(err, "Cannot hash password")
	}

	if err := user.Create(); err != nil {
		util.Danger(err, "Cannot create user")
	}
	c.Redirect(http.StatusFound, "/login")
}

// // POST /login
// // Login the user given the email and password
func Login(c *gin.Context) {
	user, err := models.UserByEmail(c.PostForm("email"))
	if err != nil {
		util.Danger(err, "Cannot find user")
	}

	password := c.PostForm("password")
	fmt.Println(user, password)
	isValidPassword := util.CheckPasswordHash(password, user.Password)
	if isValidPassword != true {
		util.Danger(err, "Invalid password")
	}

	session, err := user.CreateSession()
	if err != nil {
		util.Danger(err, "Cannot create session")
	}
	cookie := http.Cookie{
		Name:     "_cookie",
		Value:    session.UUID,
		HttpOnly: true,
	}
	c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
	c.Redirect(http.StatusFound, "/")
}

// GET /logout
// Logout logs the user out
func Logout(c *gin.Context) {
	cookie, err := c.Cookie("_cookie")
	if err != http.ErrNoCookie {
		util.Warning(err, "Failed to get cookie")
		session := models.Session{UUID: cookie}
		session.DeleteByUUID()
	}
	c.Redirect(http.StatusFound, "/")
}
