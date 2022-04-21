package test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/akkien/learngo-explorer-api/middlewares"
	"github.com/akkien/learngo-explorer-api/models"
	"github.com/gin-gonic/gin"
)

func (suite *TestSuiteEnv) Test_GetBlocks() {
	a := suite.Assert()

	req, w, err := setGetBlocksRouter(suite.m)
	if err != nil {
		a.Error(err)
	}

	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := []models.Block{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	//expected := []models.Block{}
	a.Equal(10, len(actual))
}

func setGetBlocksRouter(m *models.DBModel,
) (*http.Request, *httptest.ResponseRecorder, error) {
	// Setup
	router := gin.New()

	router.Use(middlewares.SetUserStatus())

	router.Use(func(c *gin.Context) {
		c.Set("db", m)
		c.Next()
	})

	routes(router)

	// Make request
	req, err := http.NewRequest(http.MethodGet, "/api/blocks?page=1&limit=10", nil)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	return req, w, nil
}
