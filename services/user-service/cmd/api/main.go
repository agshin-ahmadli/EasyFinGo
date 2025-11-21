package main

import (
	"EasyFinGo/internal/app"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	_ "go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	_ "go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(newLogger),
		app.Module,
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	).Run()
}

func newLogger() (*zap.Logger, error) {
	return zap.NewDevelopment()
}
