package app

import (
	"EasyFinGo/internal/config"
	"EasyFinGo/internal/database"
	"EasyFinGo/internal/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type App struct {
	Config *config.Config
	DB     *gorm.DB
	Router *gin.Engine
}

func New() (*App, error) {
	cfg, err := config.LoadConfig()

	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	db, err := database.NewPostgresDB(&cfg.Database)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	r := setupRoter(cfg)

	app := &App{
		Config: cfg,
		DB:     db,
		Router: r,
	}

	router.SetupRoutes(app.Router, db)

	return app, nil

}

func setupRoter(cfg *config.Config) *gin.Engine {
	if cfg.Server.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//TODO: Add custom middleware later
	//r.Use(middleware.CORS())
	//r.Use(middleware.RequestID())
	//r.Use(middleware.RateLimiter())
	//r.Use(middleware.ErrorHandler())
	return r
}

func (a *App) Run() error {
	addr := fmt.Sprintf("%s:%s", a.Config.Server.Host, a.Config.Server.Port)
	log.Printf("server starting on %s (env:%s)", addr, a.Config.Server.Env)
	return a.Router.Run(addr)
}

func (a *App) Shutdown() {
	log.Println("initiating graceful shutdown")

	sqlDB, err := a.DB.DB()
	if err != nil {
		log.Printf("error getting db instance: %v", err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Printf(" Error closing database: %v", err)
		return
	}
}
