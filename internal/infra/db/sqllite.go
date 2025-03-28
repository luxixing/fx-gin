package db

import (
	"database/sql"
	"sync"

	"github.com/luxixing/fx-gin/internal/config"
	"github.com/luxixing/fx-gin/pkg/registry"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func init() {
	registry.Register(
		fx.Provide(
			NewSQLiteConnection,
		),
	)
}

var (
	sqliteDBInstance *sql.DB
	sqliteOnce       sync.Once
)

type SQLiteConnectionParams struct {
	fx.In
	Config *config.Config
}

// CreateSQLiteConnection creates and returns a SQLite database connection
func NewSQLiteConnection(p SQLiteConnectionParams) (*sql.DB, error) {
	var err error
	sqliteOnce.Do(func() {
		sqliteDBInstance, err = sql.Open("sqlite3", p.Config.Database.Database)
		if err != nil {
			zap.S().Fatalw("failed to open SQLite database", "error", err)
			return
		}

		if err = sqliteDBInstance.Ping(); err != nil {
			zap.S().Fatalw("failed to ping SQLite database", "error", err)
			return
		}

		// Run database migrations
		migrationCfg := MigrationConfig{
			DB: sqliteDBInstance,
		}
		if err = RunMigrations(migrationCfg); err != nil {
			zap.S().Fatalw("failed to run migrations", "error", err)
			return
		}
	})

	return sqliteDBInstance, err
}
