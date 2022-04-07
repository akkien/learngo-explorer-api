package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/akkien/learngo-explorer-api/middlewares"
	"github.com/akkien/learngo-explorer-api/models"
	"github.com/gin-gonic/gin"
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

func (app *application) configure(m *models.DBModel) {
	// app.router.Use(gin.Logger())
	// app.router.Use(gin.Recovery())
	// app.router.Use(gin.ErrorLogger())

	app.router.Use(middlewares.SetUserStatus())

	app.router.Use(func(c *gin.Context) {
		c.Set("db", m)
		c.Next()
	})
}

func (app *application) serve() error {
	app.routes()

	return app.router.Run("localhost:" + strconv.Itoa(app.config.port))
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 5001, "Server port to listen on")
	flag.StringVar(&cfg.db.dsn, "dsn", "postgres://postgres:mysecret@127.0.0.1:5432/explorer?sslmode=disable", "DSN")
	flag.StringVar(&cfg.env, "env", "development", "Application enviornment {development|production|maintenance}")
	flag.StringVar(&cfg.secretkey, "secret", "bRWmrwNUTqNUuzckjxsFlHZjxHkjrzKP", "secret key")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Setup Database
	// "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := models.OpenDB(cfg.db.dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	dbConn, err := db.DB()
	defer dbConn.Close()

	// Setup Gin router
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// Init & Run application
	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		version:  version,
		router:   router,
	}

	app.configure(&models.DBModel{DB: db})
	app.serve()
}
