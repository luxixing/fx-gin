package service

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/luxixing/fx-gin/internal/domain"
	"github.com/luxixing/fx-gin/pkg/logger"
	"github.com/luxixing/fx-gin/pkg/registry"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func init() {
	registry.Register(
		fx.Provide(NewUserService),
	)
}

// UserServiceParams represents the parameters required for user service initialization
type UserServiceParams struct {
	fx.In

	UserRepo    domain.UserRepo
	ProfileRepo domain.ProfileRepo
	RoleRepo    domain.RoleRepo
}

// userService implements the user service interface
type userService struct {
	userRepo    domain.UserRepo
	profileRepo domain.ProfileRepo
	roleRepo    domain.RoleRepo
}

// NewUserService creates a new user service instance
func NewUserService(p UserServiceParams) domain.UserService {
	return &userService{
		userRepo:    p.UserRepo,
		profileRepo: p.ProfileRepo,
		roleRepo:    p.RoleRepo,
	}
}

// Register registers a new user
func (s *userService) Register(ctx context.Context, req *domain.UserRequest) (*domain.User, error) {
	logger.Debug(ctx, "register", zap.Any("req", req))
	// Check if username already exists
	existingUser, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to check username: %w", err)
	}
	if existingUser != nil {
		return nil, errors.New("username already exists")
	}

	// Check if email already exists
	existingUser, err = s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to check email: %w", err)
	}
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := s.hashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Create user
	user := &domain.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		Status:   domain.UserStatusActive, // Default to active
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Create default profile
	profile := &domain.Profile{
		UserID:   user.ID,
		Nickname: user.Username,
	}
	if err := s.profileRepo.Create(ctx, profile); err != nil {
		// Don't return error, user can still continue
	}

	// Assign default role (if exists)
	defaultRole, _ := s.roleRepo.GetByName(ctx, "user")
	if defaultRole != nil {
		if err := s.roleRepo.AddRoleToUser(ctx, user.ID, defaultRole.ID); err != nil {
		}
	}

	// Don't return password
	user.Password = ""
	return user, nil
}

// GetUserByID retrieves a user by ID
func (s *userService) GetUserByID(ctx context.Context, id int64) (*domain.User, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// Don't return password
	user.Password = ""
	return user, nil
}

// UpdateUser updates user information
func (s *userService) UpdateUser(ctx context.Context, id int64, req *domain.UserRequest) error {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		return errors.New("user not found")
	}

	// Check if username is already used by another user
	if req.Username != user.Username {
		existingUser, err := s.userRepo.GetByUsername(ctx, req.Username)
		if err != nil {
			return fmt.Errorf("failed to check username: %w", err)
		}
		if existingUser != nil && existingUser.ID != id {
			return errors.New("username already exists")
		}
		user.Username = req.Username
	}

	// Check if email is already used by another user
	if req.Email != user.Email {
		existingUser, err := s.userRepo.GetByEmail(ctx, req.Email)
		if err != nil {
			return fmt.Errorf("failed to check email: %w", err)
		}
		if existingUser != nil && existingUser.ID != id {
			return errors.New("email already registered")
		}
		user.Email = req.Email
	}

	// Update password if provided
	if req.Password != "" {
		hashedPassword, err := s.hashPassword(req.Password)
		if err != nil {
			return fmt.Errorf("failed to hash password: %w", err)
		}
		user.Password = hashedPassword
	}

	// Update user
	if err := s.userRepo.Update(ctx, user); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

// DeleteUser deletes a user
func (s *userService) DeleteUser(ctx context.Context, id int64) error {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		return errors.New("user not found")
	}

	// Delete user
	if err := s.userRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

// ListUsers retrieves a list of users
func (s *userService) ListUsers(ctx context.Context, page, pageSize int) ([]*domain.User, int, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	users, err := s.userRepo.List(ctx, offset, pageSize)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get user list: %w", err)
	}

	// Get total count
	total, err := s.userRepo.Count(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get user count: %w", err)
	}

	// Remove passwords
	for _, user := range users {
		user.Password = ""
	}

	return users, total, nil
}

// Login user login
func (s *userService) Login(ctx context.Context, req *domain.LoginRequest) (*domain.TokenResponse, error) {
	// Find user by username
	user, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		return nil, errors.New("invalid username or password")
	}

	// Check user status
	if user.Status != domain.UserStatusActive {
		return nil, errors.New("account is not active or has been locked")
	}

	// Verify password
	if !s.verifyPassword(user.Password, req.Password) {
		return nil, errors.New("invalid username or password")
	}

	// Generate token (simple implementation, should use JWT or other standard methods in production)
	tokenBytes := make([]byte, 32)
	if _, err := rand.Read(tokenBytes); err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}
	token := base64.StdEncoding.EncodeToString(tokenBytes)
	expiresAt := time.Now().Add(24 * time.Hour).Unix()

	// In production, token should be saved to database or cache
	// ...

	return &domain.TokenResponse{
		Token:     token,
		ExpiresAt: expiresAt,
	}, nil
}

// ValidateToken validates a token
func (s *userService) ValidateToken(ctx context.Context, token string) (int64, error) {
	// In production, token should be validated against database or cache
	// For demo purposes, return a fixed user ID
	return 1, nil
}

// GetUserWithProfile retrieves a user and their profile
func (s *userService) GetUserWithProfile(ctx context.Context, id int64) (*domain.UserWithProfile, error) {
	// Get user
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// Get profile
	profile, err := s.profileRepo.GetByUserID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user profile: %w", err)
	}
	if profile == nil {
		// Create default profile if not exists
		profile = &domain.Profile{
			UserID:   user.ID,
			Nickname: user.Username,
		}
		if err := s.profileRepo.Create(ctx, profile); err != nil {
			return nil, fmt.Errorf("failed to create user profile: %w", err)
		}
	}

	// Don't return password
	user.Password = ""
	return &domain.UserWithProfile{
		User:    *user,
		Profile: *profile,
	}, nil
}

// GetUserWithRoles retrieves a user and their roles
func (s *userService) GetUserWithRoles(ctx context.Context, id int64) (*domain.UserWithRoles, error) {
	// Get user
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// Get roles
	rolePointers, err := s.roleRepo.GetUserRoles(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user roles: %w", err)
	}

	// Convert pointer slice to value slice
	roles := make([]domain.Role, len(rolePointers))
	for i, r := range rolePointers {
		if r != nil {
			roles[i] = *r
		}
	}

	// Don't return password
	user.Password = ""
	return &domain.UserWithRoles{
		User:  *user,
		Roles: roles,
	}, nil
}

// Helper function: Hash password
func (s *userService) hashPassword(password string) (string, error) {
	// Simple implementation, in production use bcrypt or argon2
	hash := sha256.Sum256([]byte(password))
	return base64.StdEncoding.EncodeToString(hash[:]), nil
}

// Helper function: Verify password
func (s *userService) verifyPassword(hashedPassword, password string) bool {
	// Simple implementation, in production use bcrypt or argon2
	hash := sha256.Sum256([]byte(password))
	inputHash := base64.StdEncoding.EncodeToString(hash[:])
	return subtle.ConstantTimeCompare([]byte(hashedPassword), []byte(inputHash)) == 1
}
