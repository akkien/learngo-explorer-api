package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/akkien/learngo-explorer-api/middlewares"
	"github.com/akkien/learngo-explorer-api/models"
	"github.com/gin-gonic/gin"
)

// func (suite *TestSuiteEnv) Test_CreateAlbum() {
// 	req, w, err := setCreateAlbumRouter(suite.m)
// 	a := suite.Assert()
// 	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
// 	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")
// 	body, err := ioutil.ReadAll(w.Body)
// 	if err != nil {
// 		a.Error(err)
// 	}

// 	actual := models.Book{}
// 	if err := json.Unmarshal(body, &actual); err != nil {
// 		a.Error(err)
// 	}

// 	expected := models.Book{}
// 	a.Equal(expected, actual)

// }

func setCreateAlbumRouter(m *models.DBModel,
	body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	// Setup
	router := gin.Default()

	router.Use(middlewares.SetUserStatus())

	router.Use(func(c *gin.Context) {
		c.Set("db", m)
		c.Next()
	})

	routes(router)

	// Make request
	req, err := http.NewRequest(http.MethodPost, "/api/blocks", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	return req, w, nil
}
