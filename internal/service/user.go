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
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// UserServiceParams 用户服务参数
type UserServiceParams struct {
	fx.In

	UserRepo    domain.UserRepo
	ProfileRepo domain.ProfileRepo
	RoleRepo    domain.RoleRepo
	Logger      *zap.Logger
}

// userService 用户服务实现
type userService struct {
	userRepo    domain.UserRepo
	profileRepo domain.ProfileRepo
	roleRepo    domain.RoleRepo
	logger      *zap.Logger
}

// NewUserService 创建用户服务
func NewUserService(p UserServiceParams) domain.UserService {
	return &userService{
		userRepo:    p.UserRepo,
		profileRepo: p.ProfileRepo,
		roleRepo:    p.RoleRepo,
		logger:      p.Logger,
	}
}

// Register 注册新用户
func (s *userService) Register(ctx context.Context, req *domain.UserRequest) (*domain.User, error) {
	// 检查用户名是否已存在
	existingUser, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, fmt.Errorf("检查用户名失败: %w", err)
	}
	if existingUser != nil {
		return nil, errors.New("用户名已存在")
	}

	// 检查邮箱是否已存在
	existingUser, err = s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("检查邮箱失败: %w", err)
	}
	if existingUser != nil {
		return nil, errors.New("邮箱已被注册")
	}

	// 哈希密码
	hashedPassword, err := s.hashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("密码加密失败: %w", err)
	}

	// 创建用户
	user := &domain.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		Status:   domain.UserStatusActive, // 默认激活
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("创建用户失败: %w", err)
	}

	// 创建默认配置文件
	profile := &domain.Profile{
		UserID:   user.ID,
		Nickname: user.Username,
	}
	if err := s.profileRepo.Create(ctx, profile); err != nil {
		s.logger.Warn("创建用户配置文件失败", zap.Error(err))
		// 不返回错误，用户仍然可以继续使用
	}

	// 赋予默认角色（如有）
	defaultRole, _ := s.roleRepo.GetByName(ctx, "user")
	if defaultRole != nil {
		if err := s.roleRepo.AddRoleToUser(ctx, user.ID, defaultRole.ID); err != nil {
			s.logger.Warn("为用户添加默认角色失败", zap.Error(err))
		}
	}

	// 不返回密码
	user.Password = ""
	return user, nil
}

// GetUserByID 根据ID获取用户
func (s *userService) GetUserByID(ctx context.Context, id int64) (*domain.User, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("获取用户失败: %w", err)
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}

	// 不返回密码
	user.Password = ""
	return user, nil
}

// UpdateUser 更新用户信息
func (s *userService) UpdateUser(ctx context.Context, id int64, req *domain.UserRequest) error {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("获取用户失败: %w", err)
	}
	if user == nil {
		return errors.New("用户不存在")
	}

	// 检查用户名是否已被其他用户使用
	if req.Username != user.Username {
		existingUser, err := s.userRepo.GetByUsername(ctx, req.Username)
		if err != nil {
			return fmt.Errorf("检查用户名失败: %w", err)
		}
		if existingUser != nil && existingUser.ID != id {
			return errors.New("用户名已存在")
		}
		user.Username = req.Username
	}

	// 检查邮箱是否已被其他用户使用
	if req.Email != user.Email {
		existingUser, err := s.userRepo.GetByEmail(ctx, req.Email)
		if err != nil {
			return fmt.Errorf("检查邮箱失败: %w", err)
		}
		if existingUser != nil && existingUser.ID != id {
			return errors.New("邮箱已被注册")
		}
		user.Email = req.Email
	}

	// 更新密码（如果提供）
	if req.Password != "" {
		hashedPassword, err := s.hashPassword(req.Password)
		if err != nil {
			return fmt.Errorf("密码加密失败: %w", err)
		}
		user.Password = hashedPassword
	}

	// 更新用户
	if err := s.userRepo.Update(ctx, user); err != nil {
		return fmt.Errorf("更新用户失败: %w", err)
	}

	return nil
}

// DeleteUser 删除用户
func (s *userService) DeleteUser(ctx context.Context, id int64) error {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("获取用户失败: %w", err)
	}
	if user == nil {
		return errors.New("用户不存在")
	}

	// 删除用户
	if err := s.userRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("删除用户失败: %w", err)
	}

	return nil
}

// ListUsers 获取用户列表
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
		return nil, 0, fmt.Errorf("获取用户列表失败: %w", err)
	}

	// 获取总数
	total, err := s.userRepo.Count(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("获取用户总数失败: %w", err)
	}

	// 移除密码
	for _, user := range users {
		user.Password = ""
	}

	return users, total, nil
}

// Login 用户登录
func (s *userService) Login(ctx context.Context, req *domain.LoginRequest) (*domain.TokenResponse, error) {
	// 通过用户名查找用户
	user, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, fmt.Errorf("获取用户失败: %w", err)
	}
	if user == nil {
		return nil, errors.New("用户名或密码错误")
	}

	// 检查用户状态
	if user.Status != domain.UserStatusActive {
		return nil, errors.New("账户未激活或已被锁定")
	}

	// 验证密码
	if !s.verifyPassword(user.Password, req.Password) {
		return nil, errors.New("用户名或密码错误")
	}

	// 生成令牌（简单实现，实际应使用JWT或其他标准方式）
	tokenBytes := make([]byte, 32)
	if _, err := rand.Read(tokenBytes); err != nil {
		return nil, fmt.Errorf("生成令牌失败: %w", err)
	}
	token := base64.StdEncoding.EncodeToString(tokenBytes)
	expiresAt := time.Now().Add(24 * time.Hour).Unix()

	// 实际项目中应该将令牌保存到数据库或缓存中
	// ...

	return &domain.TokenResponse{
		Token:     token,
		ExpiresAt: expiresAt,
	}, nil
}

// ValidateToken 验证令牌
func (s *userService) ValidateToken(ctx context.Context, token string) (int64, error) {
	// 实际项目中应该从数据库或缓存中验证令牌
	// 这里为了演示简单返回一个固定的用户ID
	return 1, nil
}

// GetUserWithProfile 获取用户及其配置文件
func (s *userService) GetUserWithProfile(ctx context.Context, id int64) (*domain.UserWithProfile, error) {
	// 获取用户
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("获取用户失败: %w", err)
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}

	// 获取配置文件
	profile, err := s.profileRepo.GetByUserID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("获取用户配置文件失败: %w", err)
	}
	if profile == nil {
		// 如果配置文件不存在，创建一个默认的
		profile = &domain.Profile{
			UserID:   user.ID,
			Nickname: user.Username,
		}
		if err := s.profileRepo.Create(ctx, profile); err != nil {
			return nil, fmt.Errorf("创建用户配置文件失败: %w", err)
		}
	}

	// 不返回密码
	user.Password = ""
	return &domain.UserWithProfile{
		User:    *user,
		Profile: *profile,
	}, nil
}

// GetUserWithRoles 获取用户及其角色
func (s *userService) GetUserWithRoles(ctx context.Context, id int64) (*domain.UserWithRoles, error) {
	// 获取用户
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("获取用户失败: %w", err)
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}

	// 获取角色
	rolePointers, err := s.roleRepo.GetRolesByUserID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("获取用户角色失败: %w", err)
	}

	// 将指针切片转换为值切片
	roles := make([]domain.Role, len(rolePointers))
	for i, r := range rolePointers {
		if r != nil {
			roles[i] = *r
		}
	}

	// 不返回密码
	user.Password = ""
	return &domain.UserWithRoles{
		User:  *user,
		Roles: roles,
	}, nil
}

// 辅助函数：哈希密码
func (s *userService) hashPassword(password string) (string, error) {
	// 简单实现，实际应使用bcrypt或argon2等算法
	hash := sha256.Sum256([]byte(password))
	return base64.StdEncoding.EncodeToString(hash[:]), nil
}

// 辅助函数：验证密码
func (s *userService) verifyPassword(hashedPassword, password string) bool {
	// 简单实现，实际应使用bcrypt或argon2等算法
	hash := sha256.Sum256([]byte(password))
	inputHash := base64.StdEncoding.EncodeToString(hash[:])
	return subtle.ConstantTimeCompare([]byte(hashedPassword), []byte(inputHash)) == 1
}
