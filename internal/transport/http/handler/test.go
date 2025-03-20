package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luxixing/fx-gin/internal/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// TestHandlerParams 测试处理器参数
type TestHandlerParams struct {
	fx.In
	Config *config.Config
	Logger *zap.Logger
	//service
}

// TestHandler 测试处理器
type TestHandler struct {
	cfg    *config.Config
	logger *zap.Logger
}

// NewTestHandler 创建测试处理器
func NewTestHandler(p TestHandlerParams) *TestHandler {
	return &TestHandler{
		cfg:    p.Config,
		logger: p.Logger,
	}
}

// Test 测试接口
func (h *TestHandler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "test",
	})
}
