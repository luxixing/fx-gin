package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/luxixing/fx-gin/internal/config"
	"github.com/luxixing/fx-gin/internal/infra/db"
	"github.com/luxixing/fx-gin/internal/repo"
	"github.com/luxixing/fx-gin/internal/service"
	router "github.com/luxixing/fx-gin/internal/transport/http"
	"github.com/luxixing/fx-gin/internal/transport/http/handler"
	"github.com/luxixing/fx-gin/pkg/logger"

	"go.uber.org/fx"

	"go.uber.org/zap"
)

func provide() fx.Option {
	return fx.Options(
		//config
		config.Module,
		//infra
		db.Module,
		logger.Module,
		//transport
		handler.Module,
		router.Module,
		//repo
		repo.Module,
		//service
		service.Module,
	)
}

func main() {
	envFile := flag.String("env", ".env", "path to the env file")
	flag.Parse()
	if err := godotenv.Load(*envFile); err != nil {
		log.Printf("waring: failed to load env file: %v", err)
	}
	app := fx.New(
		provide(),
		fx.Invoke(
			func(
				lc fx.Lifecycle,
				cfg *config.Config,
				logger *zap.Logger,
				db *sql.DB,
				router *router.Router,
			) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						logger.Info("Starting application...")
						logger.Info("Getting config...", zap.Any("config", cfg))
						srv := &http.Server{
							Addr:    fmt.Sprintf("%s:%d", cfg.App.Host, cfg.App.Port),
							Handler: router,
						}
						//start server
						go func() {
							if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
								log.Fatalf("listen failed, err=%s", err.Error())
								logger.Fatal("listen failed", zap.Error(err))
							}
						}()
						//todo more
						return nil
					},
					OnStop: func(ctx context.Context) error {
						logger.Info("Stopping application...")
						//todo
						db.Close()
						return nil
					},
				})
			},
		),
	)
	app.Run()
}
