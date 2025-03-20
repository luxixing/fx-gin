package db

import (
	"database/sql"
	"sync"

	"github.com/luxixing/fx-gin/internal/config"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var (
	sqliteDBInstance *sql.DB
	sqliteOnce       sync.Once
)

// CreateSQLiteConnection creates and returns a SQLite database connection
func CreateSQLiteConnection(cfg *config.Config, logger *zap.Logger) (*sql.DB, error) {
	var err error
	sqliteOnce.Do(func() {
		sqliteDBInstance, err = sql.Open("sqlite3", cfg.Database.Database)
		if err != nil {
			logger.Fatal("failed to open SQLite database", zap.Error(err))
			return
		}

		if err = sqliteDBInstance.Ping(); err != nil {
			logger.Fatal("failed to ping SQLite database", zap.Error(err))
			return
		}

		// 执行数据库迁移
		migrationCfg := MigrationConfig{
			Logger: logger,
			DB:     sqliteDBInstance,
		}
		if err = RunMigrations(migrationCfg); err != nil {
			logger.Error("数据库迁移失败", zap.Error(err))
			return
		}
	})

	return sqliteDBInstance, err
}

// Module provides the SQLite connection for dependency injection
var Module = fx.Options(
	fx.Provide(func(cfg *config.Config, logger *zap.Logger) (*sql.DB, error) {
		return CreateSQLiteConnection(cfg, logger)
	}),
)
