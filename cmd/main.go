package main

import (
	"time"
	"zimniyles/fibergo/config"
	"zimniyles/fibergo/internal/home"
	"zimniyles/fibergo/internal/post"
	"zimniyles/fibergo/pkg/database"
	"zimniyles/fibergo/pkg/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v3"
)

func main() {
	config.Init()
	config.NewDBConfig()
	logConfig := config.NewLogConfig()
	dbConfig := config.NewDBConfig()
	customLogger := logger.NewLogger(logConfig)

	app := fiber.New()

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New())
	app.Static("/public", "./public")
	dbpool := database.CreateDataBasePool(dbConfig, customLogger)
	defer dbpool.Close()

	storage:=postgres.New(postgres.Config{
		DB: dbpool,
		Table: "sessions",
		Reset: false,
		GCInterval: 10 * time.Second,
	})
	store := session.New(session.Config{
		Storage: storage,
	})

	//Repositories
	postRepository := post.NewPostRepository(dbpool, customLogger)
	homeRepository := home.NewUsersRepository(dbpool, customLogger)

	//Handlers
	home.NewHandler(app, customLogger, postRepository, store, homeRepository)
	post.NewHandler(app, customLogger, postRepository)

	//Listen and serve
	app.Listen(":3000")
}
