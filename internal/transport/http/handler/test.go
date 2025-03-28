package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luxixing/fx-gin/internal/config"
	"github.com/luxixing/fx-gin/pkg/registry"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func init() {
	registry.Register(
		fx.Provide(NewTestHandler),
	)
}

// TestHandlerParams embed fx.In for dependency injection
type TestHandlerParams struct {
	fx.In
	Config *config.Config
}

// TestHandler for handling test requests
type TestHandler struct {
	cfg    *config.Config
	logger *zap.Logger
}

// NewTestHandler creates a new TestHandler
func NewTestHandler(p TestHandlerParams) *TestHandler {
	return &TestHandler{
		cfg: p.Config,
	}
}

// Test handles test endpoint requests
func (h *TestHandler) Test(c *gin.Context) {
	zap.S().Info("test handler called")
	c.JSON(http.StatusOK, gin.H{
		"message": "test",
	})
}
