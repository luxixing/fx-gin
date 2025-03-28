package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/luxixing/fx-gin/internal/domain"
	"github.com/luxixing/fx-gin/pkg/logger"
	"github.com/luxixing/fx-gin/pkg/registry"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func init() {
	registry.Register(
		fx.Provide(NewUserRepo),
	)
}

// UserRepoParams represents the parameters required for user repository initialization
type UserRepoParams struct {
	fx.In

	DB *sql.DB
}

// userRepo implements the user repository interface
type userRepo struct {
	db *sql.DB
}

// NewUserRepo creates a new user repository instance
func NewUserRepo(p UserRepoParams) domain.UserRepo {
	return &userRepo{
		db: p.DB,
	}
}

// Create creates a new user
func (r *userRepo) Create(ctx context.Context, user *domain.User) error {
	logger.Debug(ctx, "Create", zap.Any("user", user))
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	query := `INSERT INTO users (username, email, password, status, created_at, updated_at) 
              VALUES (?, ?, ?, ?, ?, ?)`

	result, err := r.db.ExecContext(ctx, query,
		user.Username,
		user.Email,
		user.Password,
		user.Status,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = id
	return nil
}

// GetByID retrieves a user by ID
func (r *userRepo) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	query := `SELECT id, username, email, password, status, created_at, updated_at 
              FROM users WHERE id = ?`

	var user domain.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// GetByUsername retrieves a user by username
func (r *userRepo) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	query := `SELECT id, username, email, password, status, created_at, updated_at 
              FROM users WHERE username = ?`

	var user domain.User
	err := r.db.QueryRowContext(ctx, query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// GetByEmail retrieves a user by email
func (r *userRepo) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT id, username, email, password, status, created_at, updated_at 
              FROM users WHERE email = ?`

	var user domain.User
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// Update updates user information
func (r *userRepo) Update(ctx context.Context, user *domain.User) error {
	user.UpdatedAt = time.Now()

	query := `UPDATE users SET username = ?, email = ?, password = ?, status = ?, updated_at = ? 
              WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query,
		user.Username,
		user.Email,
		user.Password,
		user.Status,
		user.UpdatedAt,
		user.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

// Delete deletes a user
func (r *userRepo) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

// List retrieves a list of users
func (r *userRepo) List(ctx context.Context, offset, limit int) ([]*domain.User, error) {
	query := `SELECT id, username, email, password, status, created_at, updated_at 
              FROM users ORDER BY id DESC LIMIT ?, ?`

	rows, err := r.db.QueryContext(ctx, query, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		user := &domain.User{}
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.Status,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Count retrieves the total number of users
func (r *userRepo) Count(ctx context.Context) (int, error) {
	query := `SELECT COUNT(*) FROM users`

	var count int
	err := r.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
