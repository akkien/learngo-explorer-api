package controllers

import (
	"fmt"
	"net/http"

	"github.com/akkien/explorer-modern/models"
	"github.com/akkien/explorer-modern/util"
	"gopkg.in/go-playground/validator.v9"
)

// GET /login
// Login show the login page
func Login(writer http.ResponseWriter, request *http.Request) {
	t := util.ParseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(writer, nil)
}

// GET /signup
// Signup show the signup page
func Signup(writer http.ResponseWriter, request *http.Request) {
	util.GenerateHTML(writer, nil, "login.layout", "public.navbar", "signup")
}

// POST /signup
// SignupAccount create the user account
func SignupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		util.Danger(err, "Cannot parse form")
	}
	user := models.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}

	val := validator.New()
	err = val.Struct(user)
	if err != nil {
		responseBody := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			responseBody[e.Field()] = fmt.Sprint(e)
		}
		util.GenerateHTML(writer, responseBody, "login.layout", "public.navbar", "signup")
		return
	}

	user.Password, err = util.HashPassword(user.Password)
	if err != nil {
		util.Danger(err, "Cannot hash password")
	}

	if err := user.Create(); err != nil {
		util.Danger(err, "Cannot create user")
	}
	http.Redirect(writer, request, "/login", 302)
}

// POST /authenticate
// Authenticate the user given the email and password
func Authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	user, err := models.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		util.Danger(err, "Cannot find user")
	}

	password := request.PostFormValue("password")
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
	http.SetCookie(writer, &cookie)
	http.Redirect(writer, request, "/", 302)
}

// GET /logout
// Logout logs the user out
func Logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		util.Warning(err, "Failed to get cookie")
		session := models.Session{UUID: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(writer, request, "/", 302)
}
