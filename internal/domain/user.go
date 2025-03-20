package domain

import (
	"context"
	"time"
)

// User 用户实体
type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // 不在JSON中暴露密码
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Profile 用户配置文件实体
type Profile struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Nickname  string    `json:"nickname"`
	Avatar    string    `json:"avatar"`
	Bio       string    `json:"bio"`
	Phone     string    `json:"phone"`
	Gender    int       `json:"gender"` // 0:未知 1:男 2:女
	Birthday  string    `json:"birthday"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Role 角色实体
type Role struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// UserRole 用户角色关联
type UserRole struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	RoleID    int64     `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
}

// 用户状态常量
const (
	UserStatusInactive = 0 // 未激活
	UserStatusActive   = 1 // 已激活
	UserStatusLocked   = 2 // 已锁定
)

// Gender 性别常量
const (
	GenderUnknown = 0
	GenderMale    = 1
	GenderFemale  = 2
)

// UserWithRoles 包含角色信息的用户
type UserWithRoles struct {
	User  User   `json:"user"`
	Roles []Role `json:"roles"`
}

// UserRequest 用户创建/更新请求
type UserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// ProfileRequest 配置文件创建/更新请求
type ProfileRequest struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Bio      string `json:"bio"`
	Phone    string `json:"phone"`
	Gender   int    `json:"gender"`
	Birthday string `json:"birthday"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// TokenResponse 登录响应
type TokenResponse struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

// UserWithProfile 用户及其配置文件信息
type UserWithProfile struct {
	User    User    `json:"user"`
	Profile Profile `json:"profile"`
}

// UserRepo 用户仓库接口
type UserRepo interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id int64) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, offset, limit int) ([]*User, error)
	Count(ctx context.Context) (int, error)
}

// ProfileRepo 用户配置文件仓库接口
type ProfileRepo interface {
	Create(ctx context.Context, profile *Profile) error
	GetByUserID(ctx context.Context, userID int64) (*Profile, error)
	Update(ctx context.Context, profile *Profile) error
	Delete(ctx context.Context, id int64) error
}

// RoleRepo 角色仓库接口
type RoleRepo interface {
	Create(ctx context.Context, role *Role) error
	GetByID(ctx context.Context, id int64) (*Role, error)
	GetByName(ctx context.Context, name string) (*Role, error)
	Update(ctx context.Context, role *Role) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]*Role, error)

	// 用户角色相关方法
	AddRoleToUser(ctx context.Context, userID int64, roleID int64) error
	RemoveRoleFromUser(ctx context.Context, userID int64, roleID int64) error
	GetRolesByUserID(ctx context.Context, userID int64) ([]*Role, error)
	GetUsersByRoleID(ctx context.Context, roleID int64) ([]*User, error)
}

// UserService 用户服务接口
type UserService interface {
	// 用户基本操作
	Register(ctx context.Context, req *UserRequest) (*User, error)
	GetUserByID(ctx context.Context, id int64) (*User, error)
	UpdateUser(ctx context.Context, id int64, req *UserRequest) error
	DeleteUser(ctx context.Context, id int64) error
	ListUsers(ctx context.Context, page, pageSize int) ([]*User, int, error)

	// 认证相关
	Login(ctx context.Context, req *LoginRequest) (*TokenResponse, error)
	ValidateToken(ctx context.Context, token string) (int64, error)

	// 高级功能
	GetUserWithProfile(ctx context.Context, id int64) (*UserWithProfile, error)
	GetUserWithRoles(ctx context.Context, id int64) (*UserWithRoles, error)
}
