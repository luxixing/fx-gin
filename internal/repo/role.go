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
		fx.Provide(NewRoleRepo),
	)
}

type RoleRepoParams struct {
	fx.In

	DB *sql.DB
}

type roleRepo struct {
	db *sql.DB
}

func NewRoleRepo(p RoleRepoParams) domain.RoleRepo {
	return &roleRepo{
		db: p.DB,
	}
}

func (r *roleRepo) Create(ctx context.Context, role *domain.Role) error {
	now := time.Now()
	role.CreatedAt = now
	role.UpdatedAt = now

	query := `INSERT INTO roles (name, description, created_at, updated_at) VALUES (?, ?, ?, ?)`

	result, err := r.db.ExecContext(ctx, query,
		role.Name,
		role.Description,
		role.CreatedAt,
		role.UpdatedAt,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	role.ID = id
	return nil
}

func (r *roleRepo) GetByID(ctx context.Context, id int64) (*domain.Role, error) {
	query := `SELECT id, name, description, created_at, updated_at FROM roles WHERE id = ?`

	var role domain.Role
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&role.ID,
		&role.Name,
		&role.Description,
		&role.CreatedAt,
		&role.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &role, nil
}

// GetByName 根据名称获取角色
func (r *roleRepo) GetByName(ctx context.Context, name string) (*domain.Role, error) {
	query := `SELECT id, name, description, created_at, updated_at FROM roles WHERE name = ?`

	var role domain.Role
	err := r.db.QueryRowContext(ctx, query, name).Scan(
		&role.ID,
		&role.Name,
		&role.Description,
		&role.CreatedAt,
		&role.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &role, nil
}

func (r *roleRepo) Update(ctx context.Context, role *domain.Role) error {
	role.UpdatedAt = time.Now()

	query := `UPDATE roles SET name = ?, description = ?, updated_at = ? WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query,
		role.Name,
		role.Description,
		role.UpdatedAt,
		role.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

// Delete removes a role from the database
func (r *roleRepo) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM roles WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

// List retrieves all roles from the database
func (r *roleRepo) List(ctx context.Context) ([]*domain.Role, error) {
	query := `SELECT id, name, description, created_at, updated_at FROM roles`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []*domain.Role
	for rows.Next() {
		role := &domain.Role{}
		err := rows.Scan(
			&role.ID,
			&role.Name,
			&role.Description,
			&role.CreatedAt,
			&role.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

// AddRoleToUser assigns a role to a user
func (r *roleRepo) AddRoleToUser(ctx context.Context, userID int64, roleID int64) error {
	now := time.Now()

	query := `INSERT INTO user_roles (user_id, role_id, created_at) VALUES (?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, userID, roleID, now)
	if err != nil {
		return err
	}

	return nil
}

// RemoveRoleFromUser removes a role from a user
func (r *roleRepo) RemoveRoleFromUser(ctx context.Context, userID int64, roleID int64) error {
	query := `DELETE FROM user_roles WHERE user_id = ? AND role_id = ?`

	_, err := r.db.ExecContext(ctx, query, userID, roleID)
	if err != nil {
		return err
	}

	return nil
}

// GetUserRoles retrieves all roles assigned to a user
func (r *roleRepo) GetUserRoles(ctx context.Context, userID int64) ([]*domain.Role, error) {
	query := `
		SELECT r.id, r.name, r.description, r.created_at, r.updated_at 
		FROM roles r
		JOIN user_roles ur ON r.id = ur.role_id
		WHERE ur.user_id = ?
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []*domain.Role
	for rows.Next() {
		role := &domain.Role{}
		err := rows.Scan(
			&role.ID,
			&role.Name,
			&role.Description,
			&role.CreatedAt,
			&role.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

// GetUsersByRoleID retrieves all users who have a specific role
func (r *roleRepo) GetUsersByRoleID(ctx context.Context, roleID int64) ([]*domain.User, error) {
	query := `
		SELECT u.id, u.username, u.email, u.password, u.status, u.created_at, u.updated_at
		FROM users u
		JOIN user_roles ur ON u.id = ur.user_id
		WHERE ur.role_id = ?
	`

	rows, err := r.db.QueryContext(ctx, query, roleID)
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
