package app

import (
	"film_storage/internal/app/routes"
	"film_storage/pkg/configs"
	"film_storage/pkg/logging"
	"log"
)


type App struct {
	Router *routesы.Router
	Config *configs.Config
	Logger *logging.Logger
}

func NewApp(cfg *configs.Config) *App {
	database, err := db.NewDatabase(cfg)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := user.NewRepository(database.GetDB())
	userService := user.NewService(userRepo, cfg)
	userHandler := user.NewHandler(userService)

	actorRepo := actor.NewRepository(database.GetDB())
	actorService := actor.NewService(actorRepo, cfg)
	actorHandler := actor.NewHandler(actorService)

	filmRepo := film.NewRepository(database.GetDB())
	filmService := film.NewService(filmRepo)
	filmHandler := film.NewHandler(filmService)

	router := router.NewRouter(cfg, userHandler, actorHandler, filmHandler)

	return &App{
		Router: router,
		Config: cfg,
	}
}

func (a *App) Run() {
	log.Printf("server running %s", a.Config.Addr())
	if err := a.Router.Run(a.Config.Addr()); err != nil {
		log.Fatal(err)
	}
}