package util

import (
	"errors"
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/akkien/learngo-explorer-api/models"
)

// Convenience function for printing to stdout
func P(a ...interface{}) {
	fmt.Println(a...)
}

// ErrorMessage Convenience function to redirect to the error message page
func ErrorMessage(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

// Session checks if the user is logged in and has a session, if not err is not nil
func Session(writer http.ResponseWriter, request *http.Request) (sess models.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

// ShiftPath help parse params from url path
func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i+1:]
}
