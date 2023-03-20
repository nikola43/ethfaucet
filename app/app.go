package app

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtlogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	database "github.com/nikola43/ethfaucet/database"
	"github.com/nikola43/ethfaucet/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var httpServer *fiber.App

type App struct {
}

func (a *App) Initialize(port string) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	PROD := os.Getenv("PROD")
	MYSQL_USER := os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	MYSQL_DATABASE := os.Getenv("MYSQL_DATABASE")
	XAPIKEY := os.Getenv("XAPIKEY")

	_ = XAPIKEY

	if PROD == "1" {
		MYSQL_USER = os.Getenv("MYSQL_USER_DEV")
		MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD_DEV")
		MYSQL_DATABASE = os.Getenv("MYSQL_DATABASE_DEV")
		XAPIKEY = os.Getenv("XAPIKEYDEV")
	}

	InitializeDatabase(
		MYSQL_USER,
		MYSQL_PASSWORD,
		MYSQL_DATABASE)

	database.Migrate()

	InitializeHttpServer(port)
}

func HandleRoutes(api fiber.Router) {
	routes.UserRoutes(api)
}

func InitializeHttpServer(port string) {
	httpServer = fiber.New()
	/*
		//httpServer.Use(middlewares.XApiKeyMiddleware)
		httpServer.Use(cors.New(cors.Config{
			AllowOrigins: "https://panel.ecox.stelast.com",
		}))
	*/

	httpServer.Use(jwtlogger.New())
	httpServer.Use(cors.New(cors.Config{}))

	api := httpServer.Group("/api") // /api
	v1 := api.Group("/v1")          // /api/v1
	HandleRoutes(v1)

	err := httpServer.Listen(port)
	if err != nil {
		log.Fatal(err)
	}
}

func InitializeDatabase(user, password, database_name string) {
	connectionString := fmt.Sprintf(
		"%s:%s@/%s?parseTime=true",
		user,
		password,
		database_name,
	)

	DB, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	database.GormDB, err = gorm.Open(mysql.New(mysql.Config{Conn: DB}), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal(err)
	}
}
