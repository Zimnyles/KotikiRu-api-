package main

import (
	"time"
	"zimniyles/fibergo/config"
	"zimniyles/fibergo/internal/announcements"
	"zimniyles/fibergo/internal/feed"
	"zimniyles/fibergo/internal/home"
	"zimniyles/fibergo/internal/messenger"
	"zimniyles/fibergo/internal/post"
	"zimniyles/fibergo/internal/profile"
	"zimniyles/fibergo/pkg/database"
	"zimniyles/fibergo/pkg/logger"
	"zimniyles/fibergo/pkg/middleware"

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
	authConfig := config.NewAuthConfig()
	customLogger := logger.NewLogger(logConfig)

	app := fiber.New()

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New())
	app.Static("/public", "./public")
	dbpool := database.CreateDataBasePool(dbConfig, customLogger)
	defer dbpool.Close()

	storage := postgres.New(postgres.Config{
		DB:         dbpool,
		Table:      "sessions",
		Reset:      false,
		GCInterval: 10 * time.Second,
	})
	store := session.New(session.Config{
		Storage: storage,
	})


	app.Static("/static", "./static")

	app.Use(middleware.AuthMiddleware(store))
	//Repositories
	postRepository := post.NewPostRepository(dbpool, customLogger)
	homeRepository := home.NewUsersRepository(dbpool, customLogger)
	profileRepository := profile.NewProfileRepository(dbpool, customLogger)
	announcmentsRepository := announcements.NewAnnouncementsRepository(dbpool, customLogger)
	feedRepository := feed.NewFeedRepository(dbpool, customLogger)
	messengerRepository := messenger.NewMessengerRepository(dbpool, customLogger)

	//Handlers
	home.NewHandler(app, customLogger, postRepository, store, homeRepository, authConfig)
	post.NewHandler(app, customLogger, postRepository)
	profile.NewHandler(app, customLogger, profileRepository, store)
	announcements.NewHandler(app, customLogger, announcmentsRepository, postRepository)
	feed.NewFeedHandler(app, customLogger, feedRepository, store)
	messenger.NewMessengerHandler(app, customLogger, messengerRepository, store)
	//Listen and serve
	app.Listen(":3000")
}
