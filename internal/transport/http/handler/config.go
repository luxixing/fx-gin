package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luxixing/fx-gin/internal/config"
	"github.com/luxixing/fx-gin/internal/domain"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ConfigHandlerParams 配置处理器参数
type ConfigHandlerParams struct {
	fx.In

	Config  *config.Config
	Logger  *zap.Logger
	Service domain.ConfigService
}

// ConfigHandler 配置处理器
type ConfigHandler struct {
	cfg     *config.Config
	logger  *zap.Logger
	service domain.ConfigService
}

// NewConfigHandler 创建配置处理器
func NewConfigHandler(p ConfigHandlerParams) *ConfigHandler {
	return &ConfigHandler{
		cfg:     p.Config,
		logger:  p.Logger,
		service: p.Service,
	}
}

// Create godoc
// @Summary Create a new config
// @Description Create a new config with key and value
// @Tags config
// @Accept json
// @Produce json
// @Param config body domain.ConfigRequest true "Config"
// @Success 200 {object} domain.ConfigResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /config [post]
func (h *ConfigHandler) Create(c *gin.Context) {
	var req domain.ConfigRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response, err := h.service.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// Get godoc
// @Summary Get a config by key
// @Description Get a config by key
// @Tags config
// @Produce json
// @Param key path string true "Config Key"
// @Success 200 {object} domain.ConfigResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /config/{key} [get]
func (h *ConfigHandler) Get(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "key is required",
		})
		return
	}

	response, err := h.service.Get(key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if response == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "config not found",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
