package db

import (
	"database/sql"

	"go.uber.org/zap"
)

// MigrationConfig 迁移配置
type MigrationConfig struct {
	Logger *zap.Logger
	DB     *sql.DB
}

// RunMigrations 执行数据库迁移，创建必要的表
func RunMigrations(cfg MigrationConfig) error {
	logger := cfg.Logger
	db := cfg.DB

	logger.Info("开始数据库迁移...")

	// 用户表
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL UNIQUE,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			status INTEGER DEFAULT 1,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		)
	`)
	if err != nil {
		logger.Error("创建users表失败", zap.Error(err))
		return err
	}

	// 用户资料表
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS profiles (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			nickname TEXT,
			avatar TEXT,
			bio TEXT,
			phone TEXT,
			gender INTEGER DEFAULT 0,
			birthday TEXT,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		)
	`)
	if err != nil {
		logger.Error("创建profiles表失败", zap.Error(err))
		return err
	}

	// 角色表
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS roles (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL UNIQUE,
			description TEXT,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		)
	`)
	if err != nil {
		logger.Error("创建roles表失败", zap.Error(err))
		return err
	}

	// 用户角色关联表
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS user_roles (
			user_id INTEGER NOT NULL,
			role_id INTEGER NOT NULL,
			created_at TIMESTAMP NOT NULL,
			PRIMARY KEY (user_id, role_id),
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
			FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE
		)
	`)
	if err != nil {
		logger.Error("创建user_roles表失败", zap.Error(err))
		return err
	}

	// 配置表
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS configs (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			key1 TEXT NOT NULL UNIQUE,
			key2 TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		)
	`)
	if err != nil {
		logger.Error("创建configs表失败", zap.Error(err))
		return err
	}

	// 初始化默认角色
	_, err = db.Exec(`
		INSERT OR IGNORE INTO roles (name, description, created_at, updated_at)
		VALUES 
			('admin', '管理员角色', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
			('user', '普通用户角色', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`)
	if err != nil {
		logger.Error("初始化默认角色失败", zap.Error(err))
		return err
	}

	logger.Info("数据库迁移完成")
	return nil
}
