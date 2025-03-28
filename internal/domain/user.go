package domain

import (
	"context"
	"time"
)

// User represents a user entity
type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Password is not exposed in JSON
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Profile represents a user profile entity
type Profile struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Nickname  string    `json:"nickname"`
	Avatar    string    `json:"avatar"`
	Bio       string    `json:"bio"`
	Phone     string    `json:"phone"`
	Gender    int       `json:"gender"` // 0:Unknown 1:Male 2:Female
	Birthday  string    `json:"birthday"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Role represents a role entity
type Role struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// UserRole represents the relationship between users and roles
type UserRole struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	RoleID    int64     `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
}

// User status constants
const (
	UserStatusInactive = 0 // Inactive
	UserStatusActive   = 1 // Active
	UserStatusLocked   = 2 // Locked
)

// Gender constants
const (
	GenderUnknown = 0
	GenderMale    = 1
	GenderFemale  = 2
)

// UserWithRoles represents a user with their associated roles
type UserWithRoles struct {
	User  User   `json:"user"`
	Roles []Role `json:"roles"`
}

// UserRequest represents the request for creating/updating a user
type UserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// ProfileRequest represents the request for creating/updating a user profile
type ProfileRequest struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Bio      string `json:"bio"`
	Phone    string `json:"phone"`
	Gender   int    `json:"gender"`
	Birthday string `json:"birthday"`
}

// LoginRequest represents the request for user login
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// TokenResponse represents the response for user login
type TokenResponse struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

// UserWithProfile represents a user with their profile information
type UserWithProfile struct {
	User    User    `json:"user"`
	Profile Profile `json:"profile"`
}

// UserRepo defines the interface for user repository operations
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

// ProfileRepo defines the interface for user profile repository operations
type ProfileRepo interface {
	Create(ctx context.Context, profile *Profile) error
	GetByUserID(ctx context.Context, userID int64) (*Profile, error)
	Update(ctx context.Context, profile *Profile) error
	Delete(ctx context.Context, id int64) error
}

// RoleRepo defines the interface for role repository operations
type RoleRepo interface {
	Create(ctx context.Context, role *Role) error
	GetByID(ctx context.Context, id int64) (*Role, error)
	GetByName(ctx context.Context, name string) (*Role, error)
	Update(ctx context.Context, role *Role) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]*Role, error)
	AddRoleToUser(ctx context.Context, userID, roleID int64) error
	RemoveRoleFromUser(ctx context.Context, userID, roleID int64) error
	GetUserRoles(ctx context.Context, userID int64) ([]*Role, error)
}

type UserService interface {
	Register(ctx context.Context, req *UserRequest) (*User, error)
	GetUserByID(ctx context.Context, id int64) (*User, error)
	UpdateUser(ctx context.Context, id int64, req *UserRequest) error
	DeleteUser(ctx context.Context, id int64) error
	ListUsers(ctx context.Context, page, pageSize int) ([]*User, int, error)

	Login(ctx context.Context, req *LoginRequest) (*TokenResponse, error)
	ValidateToken(ctx context.Context, token string) (int64, error)

	GetUserWithProfile(ctx context.Context, id int64) (*UserWithProfile, error)
	GetUserWithRoles(ctx context.Context, id int64) (*UserWithRoles, error)
}
