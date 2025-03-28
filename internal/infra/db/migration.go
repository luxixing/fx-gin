package db

import (
	"database/sql"

	"go.uber.org/zap"
)

// MigrationConfig represents the configuration for database migration
type MigrationConfig struct {
	DB *sql.DB
}

// RunMigrations executes database migrations to create necessary tables
func RunMigrations(cfg MigrationConfig) error {
	db := cfg.DB

	zap.S().Info("Starting database migration...")

	// Users table
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
		zap.S().Error("Failed to create users table", "error", err)
		return err
	}

	// User profiles table
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
		zap.S().Error("Failed to create profiles table", "error", err)
		return err
	}

	// Roles table
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
		zap.S().Error("Failed to create roles table", "error", err)
		return err
	}

	// User roles association table
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
		zap.S().Error("Failed to create user_roles table", "error", err)
		return err
	}

	// Initialize default roles
	_, err = db.Exec(`
		INSERT OR IGNORE INTO roles (name, description, created_at, updated_at)
		VALUES 
			('admin', 'Administrator role', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
			('user', 'Regular user role', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`)
	if err != nil {
		zap.S().Errorw("Failed to initialize default roles", "error", err)
		return err
	}
	zap.S().Info("Database migration completed")
	return nil
}
