package test

import (
	"log"
	"testing"

	"github.com/akkien/learngo-explorer-api/models"
	"github.com/stretchr/testify/suite"
)

const (
	TestDbStr = "postgres://postgres:mysecret@127.0.0.1:5432/explorer?sslmode=disable"
)

func ClearTable(m *models.DBModel) {
	m.DB.Exec("DELETE FROM blocks")
	m.DB.Exec("DELETE FROM transactions")
	m.DB.Exec("DELETE FROM receipts")
	m.DB.Exec("DELETE FROM logs")
}

type TestSuiteEnv struct {
	suite.Suite
	m *models.DBModel
}

// Tests are run before they start
func (suite *TestSuiteEnv) SetupSuite() {
	db, err := models.OpenDB(TestDbStr)
	if err != nil {
		log.Fatal(err)
	}
	suite.m = &models.DBModel{DB: db}
}

// Running after each test
func (suite *TestSuiteEnv) TearDownTest() {
	/** Only call suite table if we use fake database */
	// ClearTable(suite.db)
}

// Running after all tests are completed
func (suite *TestSuiteEnv) TearDownSuite() {
	dbConn, err := suite.m.DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()
}

// This gets run automatically by `go test` so we call `suite.Run` inside it
func TestSuite(t *testing.T) {
	// This is what actually runs our suite
	suite.Run(t, new(TestSuiteEnv))
}
