package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/luxixing/fx-gin/internal/domain"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ProfileRepoParams 配置文件仓库参数
type ProfileRepoParams struct {
	fx.In

	DB     *sql.DB
	Logger *zap.Logger
}

// profileRepo 配置文件仓库实现
type profileRepo struct {
	db     *sql.DB
	logger *zap.Logger
}

// NewProfileRepo 创建配置文件仓库
func NewProfileRepo(p ProfileRepoParams) domain.ProfileRepo {
	return &profileRepo{
		db:     p.DB,
		logger: p.Logger,
	}
}

// Create 创建用户配置文件
func (r *profileRepo) Create(ctx context.Context, profile *domain.Profile) error {
	now := time.Now()
	profile.CreatedAt = now
	profile.UpdatedAt = now

	query := `INSERT INTO profiles (user_id, nickname, avatar, bio, phone, gender, birthday, created_at, updated_at) 
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := r.db.ExecContext(ctx, query,
		profile.UserID,
		profile.Nickname,
		profile.Avatar,
		profile.Bio,
		profile.Phone,
		profile.Gender,
		profile.Birthday,
		profile.CreatedAt,
		profile.UpdatedAt,
	)
	if err != nil {
		r.logger.Error("创建用户配置文件失败", zap.Error(err))
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	profile.ID = id
	return nil
}

// GetByUserID 根据用户ID获取配置文件
func (r *profileRepo) GetByUserID(ctx context.Context, userID int64) (*domain.Profile, error) {
	query := `SELECT id, user_id, nickname, avatar, bio, phone, gender, birthday, created_at, updated_at 
			  FROM profiles WHERE user_id = ?`

	var profile domain.Profile
	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&profile.ID,
		&profile.UserID,
		&profile.Nickname,
		&profile.Avatar,
		&profile.Bio,
		&profile.Phone,
		&profile.Gender,
		&profile.Birthday,
		&profile.CreatedAt,
		&profile.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		r.logger.Error("获取用户配置文件失败", zap.Error(err))
		return nil, err
	}

	return &profile, nil
}

// Update 更新用户配置文件
func (r *profileRepo) Update(ctx context.Context, profile *domain.Profile) error {
	profile.UpdatedAt = time.Now()

	query := `UPDATE profiles SET nickname = ?, avatar = ?, bio = ?, phone = ?, gender = ?, birthday = ?, updated_at = ? 
			  WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query,
		profile.Nickname,
		profile.Avatar,
		profile.Bio,
		profile.Phone,
		profile.Gender,
		profile.Birthday,
		profile.UpdatedAt,
		profile.ID,
	)

	if err != nil {
		r.logger.Error("更新用户配置文件失败", zap.Error(err))
		return err
	}

	return nil
}

// Delete 删除用户配置文件
func (r *profileRepo) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM profiles WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		r.logger.Error("删除用户配置文件失败", zap.Error(err))
		return err
	}

	return nil
}
