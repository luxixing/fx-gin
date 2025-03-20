package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luxixing/fx-gin/internal/domain"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// UserHandlerParams 用户处理器参数
type UserHandlerParams struct {
	fx.In

	UserService domain.UserService
	Logger      *zap.Logger
}

// UserHandler 用户处理器
type UserHandler struct {
	userService domain.UserService
	logger      *zap.Logger
}

// NewUserHandler 创建用户处理器
func NewUserHandler(p UserHandlerParams) *UserHandler {
	return &UserHandler{
		userService: p.UserService,
		logger:      p.Logger,
	}
}

// Register 用户注册
// @Summary 用户注册
// @Description 注册新用户
// @Tags 用户
// @Accept json
// @Produce json
// @Param user body domain.UserRequest true "用户注册信息"
// @Success 200 {object} domain.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req domain.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
		return
	}

	user, err := h.userService.Register(c.Request.Context(), &req)
	if err != nil {
		h.logger.Error("用户注册失败", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录并获取令牌
// @Tags 用户
// @Accept json
// @Produce json
// @Param login body domain.LoginRequest true "登录信息"
// @Success 200 {object} domain.TokenResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req domain.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
		return
	}

	token, err := h.userService.Login(c.Request.Context(), &req)
	if err != nil {
		h.logger.Error("用户登录失败", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, token)
}

// GetProfile 获取用户档案
// @Summary 获取用户档案
// @Description 获取用户信息和配置文件
// @Tags 用户
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} domain.UserWithProfile
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/{id}/profile [get]
func (h *UserHandler) GetProfile(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	userWithProfile, err := h.userService.GetUserWithProfile(c.Request.Context(), id)
	if err != nil {
		h.logger.Error("获取用户档案失败", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userWithProfile)
}

// GetRoles 获取用户角色
// @Summary 获取用户角色
// @Description 获取用户的所有角色
// @Tags 用户
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} domain.UserWithRoles
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/{id}/roles [get]
func (h *UserHandler) GetRoles(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	userWithRoles, err := h.userService.GetUserWithRoles(c.Request.Context(), id)
	if err != nil {
		h.logger.Error("获取用户角色失败", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userWithRoles)
}

// GetUser 获取用户信息
// @Summary 获取用户信息
// @Description 获取用户基本信息
// @Tags 用户
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} domain.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	user, err := h.userService.GetUserByID(c.Request.Context(), id)
	if err != nil {
		h.logger.Error("获取用户信息失败", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser 更新用户信息
// @Summary 更新用户信息
// @Description 更新用户基本信息
// @Tags 用户
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Param user body domain.UserRequest true "用户信息"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	var req domain.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
		return
	}

	if err := h.userService.UpdateUser(c.Request.Context(), id, &req); err != nil {
		h.logger.Error("更新用户信息失败", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "用户信息更新成功"})
}

// DeleteUser 删除用户
// @Summary 删除用户
// @Description 删除指定用户
// @Tags 用户
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	if err := h.userService.DeleteUser(c.Request.Context(), id); err != nil {
		h.logger.Error("删除用户失败", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "用户删除成功"})
}

// ListUsers 获取用户列表
// @Summary 获取用户列表
// @Description 分页获取用户列表
// @Tags 用户
// @Accept json
// @Produce json
// @Param page query int false "页码，默认1"
// @Param size query int false "每页数量，默认10"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /api/v1/users [get]
func (h *UserHandler) ListUsers(c *gin.Context) {
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

	users, total, err := h.userService.ListUsers(c.Request.Context(), page, size)
	if err != nil {
		h.logger.Error("获取用户列表失败", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"items": users,
	})
}
