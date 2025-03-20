package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/luxixing/fx-gin/internal/domain"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// UserRepoParams 用户仓库参数
type UserRepoParams struct {
	fx.In

	DB     *sql.DB
	Logger *zap.Logger
}

// userRepo 用户仓库实现
type userRepo struct {
	db     *sql.DB
	logger *zap.Logger
}

// NewUserRepo 创建用户仓库
func NewUserRepo(p UserRepoParams) domain.UserRepo {
	return &userRepo{
		db:     p.DB,
		logger: p.Logger,
	}
}

// Create 创建用户
func (r *userRepo) Create(ctx context.Context, user *domain.User) error {
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
		r.logger.Error("创建用户失败", zap.Error(err))
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = id
	return nil
}

// GetByID 根据ID获取用户
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
		r.logger.Error("获取用户失败", zap.Error(err))
		return nil, err
	}

	return &user, nil
}

// GetByUsername 根据用户名获取用户
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
		r.logger.Error("根据用户名获取用户失败", zap.Error(err))
		return nil, err
	}

	return &user, nil
}

// GetByEmail 根据邮箱获取用户
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
		r.logger.Error("根据邮箱获取用户失败", zap.Error(err))
		return nil, err
	}

	return &user, nil
}

// Update 更新用户信息
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
		r.logger.Error("更新用户失败", zap.Error(err))
		return err
	}

	return nil
}

// Delete 删除用户
func (r *userRepo) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		r.logger.Error("删除用户失败", zap.Error(err))
		return err
	}

	return nil
}

// List 获取用户列表
func (r *userRepo) List(ctx context.Context, offset, limit int) ([]*domain.User, error) {
	query := `SELECT id, username, email, password, status, created_at, updated_at 
              FROM users ORDER BY id DESC LIMIT ?, ?`

	rows, err := r.db.QueryContext(ctx, query, offset, limit)
	if err != nil {
		r.logger.Error("获取用户列表失败", zap.Error(err))
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
			r.logger.Error("扫描用户数据失败", zap.Error(err))
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Count 获取用户总数
func (r *userRepo) Count(ctx context.Context) (int, error) {
	query := `SELECT COUNT(*) FROM users`

	var count int
	err := r.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		r.logger.Error("获取用户总数失败", zap.Error(err))
		return 0, err
	}

	return count, nil
}
