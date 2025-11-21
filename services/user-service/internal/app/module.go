package app

import (
	"EasyFinGo/internal/config"
	"EasyFinGo/internal/database"
	"EasyFinGo/internal/router"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var Module = fx.Options(
	fx.Provide(config.LoadConfig),
	fx.Provide(database.NewPostgresDB),
	fx.Invoke(router.SetupRoutes),
	fx.Invoke(registerHooks),
)

func newRouter(cfg *config.Config) *gin.Engine {
	if cfg.Server.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	return r
}

func registerHooks(
	lc fx.Lifecycle,
	cfg *config.Config,
	router *gin.Engine,
	db *gorm.DB,
	logger *zap.Logger,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
			logger.Info("Starting server",
				zap.String("address", addr),
				zap.String("environment", cfg.Server.Env))

			go func() {
				if err := router.Run(addr); err != nil {
					logger.Fatal("failed to start server", zap.Error(err))
				}
			}()
			return nil
		},

		OnStop: func(ctx context.Context) error {
			logger.Info("Shutting down server gracefully")

			sqlDB, err := db.DB()
			if err != nil {
				return err
			}

			if err := sqlDB.Close(); err != nil {
				logger.Error("Error closing database", zap.Error(err))
				return err
			}
			logger.Info("database connection closed")
			return nil
		},
	})
}
