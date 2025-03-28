package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luxixing/fx-gin/internal/domain"
	"github.com/luxixing/fx-gin/pkg/logger"
	"github.com/luxixing/fx-gin/pkg/registry"
	"github.com/luxixing/fx-gin/pkg/utils"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func init() {
	registry.Register(
		fx.Provide(NewUserHandler),
	)
}

// UserHandlerParams embed fx.In for dependency injection
type UserHandlerParams struct {
	fx.In

	UserService domain.UserService
}

// UserHandler for handling user requests
type UserHandler struct {
	userService domain.UserService
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(p UserHandlerParams) *UserHandler {
	return &UserHandler{
		userService: p.UserService,
	}
}

// Register handles user registration
// @Summary User registration
// @Description Register a new user
// @Tags User
// @Accept json
// @Produce json
// @Param user body domain.UserRequest true "User registration information"
// @Success 200 {object} domain.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req domain.UserRequest
	// Create context with trace information
	ctx := utils.WithContext(c)
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error(ctx, "Invalid request parameters", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request parameters"})
		return
	}

	// Call service layer with context
	user, err := h.userService.Register(ctx, &req)
	if err != nil {
		logger.Error(ctx, "Failed to register user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	logger.Info(ctx, "User registered successfully", zap.Int64("user_id", user.ID))
	c.JSON(http.StatusOK, user)
}

// Login handles user login
// @Summary User login
// @Description User login and get token
// @Tags User
// @Accept json
// @Produce json
// @Param login body domain.LoginRequest true "Login information"
// @Success 200 {object} domain.TokenResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	ctx := utils.WithContext(c)
	logger.Info(ctx, "Processing user login request")

	var req domain.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error(ctx, "Invalid request parameters", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request parameters"})
		return
	}

	token, err := h.userService.Login(ctx, &req)
	if err != nil {
		logger.Error(ctx, "User login failed", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logger.Info(ctx, "User login successful", zap.String("username", req.Username))
	c.JSON(http.StatusOK, token)
}

// GetProfile retrieves user profile
// @Summary Get user profile
// @Description Get user information and profile
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} domain.UserWithProfile
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/{id}/profile [get]
func (h *UserHandler) GetProfile(c *gin.Context) {
	ctx := utils.WithContext(c)
	logger.Info(ctx, "Retrieving user profile")

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		logger.Error(ctx, "Invalid user ID", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	userWithProfile, err := h.userService.GetUserWithProfile(ctx, id)
	if err != nil {
		logger.Error(ctx, "Failed to get user profile", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	logger.Info(ctx, "User profile retrieved successfully", zap.Int64("user_id", id))
	c.JSON(http.StatusOK, userWithProfile)
}

// GetRoles retrieves user roles
// @Summary Get user roles
// @Description Get all roles assigned to the user
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} domain.UserWithRoles
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/{id}/roles [get]
func (h *UserHandler) GetRoles(c *gin.Context) {
	ctx := utils.WithContext(c)
	logger.Info(ctx, "Getting user roles")

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		logger.Error(ctx, "Invalid user ID", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	userWithRoles, err := h.userService.GetUserWithRoles(ctx, id)
	if err != nil {
		logger.Error(ctx, "Failed to get user roles", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	logger.Info(ctx, "Successfully retrieved user roles", zap.Int64("user_id", id))
	c.JSON(http.StatusOK, userWithRoles)
}

// GetUser retrieves user information
// @Summary Get user information
// @Description Get user basic information
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} domain.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	ctx := utils.WithContext(c)
	logger.Info(ctx, "Getting user information")

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		logger.Error(ctx, "Invalid user ID", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.userService.GetUserByID(ctx, id)
	if err != nil {
		logger.Error(ctx, "Failed to get user information", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	logger.Info(ctx, "Successfully retrieved user information", zap.Int64("user_id", id))
	c.JSON(http.StatusOK, user)
}

// UpdateUser updates user information
// @Summary Update user information
// @Description Update user basic information
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body domain.UserRequest true "User information"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	ctx := utils.WithContext(c)
	logger.Info(ctx, "Updating user information")

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		logger.Error(ctx, "Invalid user ID", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var req domain.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error(ctx, "Invalid request parameters", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request parameters"})
		return
	}

	if err := h.userService.UpdateUser(ctx, id, &req); err != nil {
		logger.Error(ctx, "Failed to update user information", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	logger.Info(ctx, "User information updated successfully", zap.Int64("user_id", id))
	c.JSON(http.StatusOK, gin.H{"message": "User information updated successfully"})
}

// DeleteUser deletes a user
// @Summary Delete user
// @Description Delete the specified user
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	ctx := utils.WithContext(c)
	logger.Info(ctx, "Deleting user")

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		logger.Error(ctx, "Invalid user ID", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := h.userService.DeleteUser(ctx, id); err != nil {
		logger.Error(ctx, "Failed to delete user", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	logger.Info(ctx, "User deleted successfully", zap.Int64("user_id", id))
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// ListUsers retrieves a list of users
// @Summary List users
// @Description Get paginated list of users
// @Tags User
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param size query int false "Items per page" default(10)
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /api/v1/users [get]
func (h *UserHandler) ListUsers(c *gin.Context) {
	ctx := utils.WithContext(c)
	logger.Info(ctx, "Getting user list")

	pageStr := c.DefaultQuery("page", "1")
	sizeStr := c.DefaultQuery("size", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		size = 10
	}

	users, total, err := h.userService.ListUsers(ctx, page, size)
	if err != nil {
		logger.Error(ctx, "Failed to get user list", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info(ctx, "Successfully retrieved user list",
		zap.Int("page", page),
		zap.Int("size", size),
		zap.Int("total", total),
	)
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"items": users,
	})
}
