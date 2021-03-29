package controllers

import (
	"fmt"
	"net/http"

	"github.com/akkien/explorer-modern/models"
	"github.com/akkien/explorer-modern/util"
	"github.com/gin-gonic/gin"
)

// GET /err?msg=
// Err shows the error message page
func Err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := util.Session(writer, request)
	if err != nil {
		util.GenerateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		util.GenerateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

// Index controller
func Index(c *gin.Context) {
	isLoggedIn := c.GetBool("is_logged_in")

	if isLoggedIn {
		blocks, err := models.Blocks(1, 10)
		txs, err := models.Transactions(1, 10)
		fmt.Println("HEllo1")
		if err != nil {
			res := gin.H{
				"title": "Home Page",
				"error": err,
			}
			render(c, res, "error.html")
		} else {
			response := models.BksTxs{Blocks: blocks, Txs: txs}
			res := gin.H{
				"title":   "Home Page",
				"payload": response,
			}
			fmt.Println("HEllo2")
			render(c, res, "private.index.html")
		}
	} else {
		res := gin.H{
			"title": "Home Page",
		}
		render(c, res, "public.index.html")
	}
}
