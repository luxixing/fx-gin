package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/luxixing/fx-gin/internal/config"
	_ "github.com/luxixing/fx-gin/internal/infra/db"
	_ "github.com/luxixing/fx-gin/internal/repo"
	_ "github.com/luxixing/fx-gin/internal/service"
	_ "github.com/luxixing/fx-gin/internal/transport/http"
	_ "github.com/luxixing/fx-gin/internal/transport/http/handler"
	_ "github.com/luxixing/fx-gin/pkg/logger"
	"github.com/luxixing/fx-gin/pkg/registry"

	"go.uber.org/fx"

	"go.uber.org/zap"
)

func main() {
	envFile := flag.String("env", ".env", "path to the env file, default is .env")
	flag.Parse()
	if err := godotenv.Load(*envFile); err != nil {
		log.Printf("waring: failed to load env file: %v", err)
	}

	app := fx.New(
		registry.GetModules(),
		fx.Invoke(
			func(
				lc fx.Lifecycle,
				cfg *config.Config,
				db *sql.DB,
				router *gin.Engine,
			) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						zap.L().Info("Starting application...")
						zap.S().Infow("Getting config...", "config", cfg)
						srv := &http.Server{
							Addr:    fmt.Sprintf("%s:%d", cfg.App.Host, cfg.App.Port),
							Handler: router,
						}
						//start server
						go func() {
							if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
								zap.S().Fatalw("Listening failed", "error", err)
							}
						}()
						//todo more
						return nil
					},
					OnStop: func(ctx context.Context) error {
						zap.S().Info("Stopping server...")
						//todo
						db.Close()
						zap.S().Sync()
						return nil
					},
				})
			},
		),
	)
	app.Run()
}
