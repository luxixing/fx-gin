package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/luxixing/fx-gin/internal/domain"
	"github.com/luxixing/fx-gin/pkg/registry"
	"go.uber.org/fx"
)

func init() {
	registry.Register(
		fx.Provide(NewProfileRepo),
	)
}

// ProfileRepoParams represents the parameters required for profile repository initialization
type ProfileRepoParams struct {
	fx.In

	DB *sql.DB
}

// profileRepo implements the profile repository interface
type profileRepo struct {
	db *sql.DB
}

// NewProfileRepo creates a new profile repository instance
func NewProfileRepo(p ProfileRepoParams) domain.ProfileRepo {
	return &profileRepo{
		db: p.DB,
	}
}

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
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	profile.ID = id
	return nil
}

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
		return nil, err
	}

	return &profile, nil
}

func (r *profileRepo) Update(ctx context.Context, profile *domain.Profile) error {
	now := time.Now()
	profile.UpdatedAt = now

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
		return err
	}

	return nil
}

func (r *profileRepo) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM profiles WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
