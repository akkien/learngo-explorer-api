package main

import (
	"errors"
	"flag"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/akkien/learngo-explorer-api/middlewares"
	"github.com/akkien/learngo-explorer-api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
	secretkey string
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	version  string
	router   *gin.Engine
}

func (app *application) configure(m *models.DBModel, logger *log.Logger) {
	// app.router.Use(gin.Logger())
	// app.router.Use(gin.Recovery())
	// app.router.Use(gin.ErrorLogger())

	app.router.Use(func(c *gin.Context) {
		c.Set("requestId", uuid.New().String())
		c.Set("db", m)
		c.Set("logger", logger)
		c.Next()
	})

	app.router.Use(middlewares.SetUserStatus())

	app.router.Use(middlewares.RequestLoggerMiddleware())

	app.routes()
}

func (app *application) serve() error {
	return app.router.Run("localhost:" + strconv.Itoa(app.config.port))
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 5001, "Server port to listen on")
	flag.StringVar(&cfg.db.dsn, "dsn", "postgres://postgres:mysecret@127.0.0.1:5432/explorer?sslmode=disable", "DSN")
	flag.StringVar(&cfg.env, "env", "development", "Application enviornment {development|production|maintenance}")
	flag.StringVar(&cfg.secretkey, "secret", "bRWmrwNUTqNUuzckjxsFlHZjxHkjrzKP", "secret key")

	flag.Parse()

	f := createLogFile()

	infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(f, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Setup Database
	// "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := models.OpenDB(cfg.db.dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	dbConn, err := db.DB()
	if err != nil {
		errorLog.Fatal(err)
	}
	defer dbConn.Close()

	// Setup Gin router
	gin.SetMode(gin.ReleaseMode)

	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()

	// Init & Run application
	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		version:  version,
		router:   router,
	}

	app.configure(&models.DBModel{DB: db}, infoLog)
	app.serve()
}

func createLogFile() *os.File {
	logsDir := "logs"
	if _, err := os.Stat(logsDir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(logsDir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
	f, err := os.OpenFile("./logs/explorer.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return f
}
