package http

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/luxixing/fx-gin/docs/swagger"
	"github.com/luxixing/fx-gin/internal/transport/http/handler"
	"github.com/luxixing/fx-gin/internal/transport/http/middleware"
	"github.com/luxixing/fx-gin/pkg/registry"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
)

func init() {
	registry.Register(
		fx.Provide(NewRouter),
	)
}

// RouterParams represents the parameters required for router initialization
type RouterParams struct {
	fx.In

	TestHandler *handler.TestHandler
	UserHandler *handler.UserHandler
}

// NewRouter creates and configures the Gin router
func NewRouter(p RouterParams) *gin.Engine {
	r := gin.New()

	// Add third-party CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Add request context middleware
	r.Use(middleware.RequestContext())
	// Replace default gin.Logger with zap logger middleware
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	// Swagger documentation
	// Create swagger documentation routes
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/api/v1")
	{
		v1.GET("/test", p.TestHandler.Test)
		// Add user-related routes
		users := v1.Group("/users")
		{
			users.POST("/register", p.UserHandler.Register)
			users.POST("/login", p.UserHandler.Login)
			users.GET("/:id", p.UserHandler.GetUser)
			users.PUT("/:id", p.UserHandler.UpdateUser)
			users.DELETE("/:id", p.UserHandler.DeleteUser)
			users.GET("", p.UserHandler.ListUsers)
			users.GET("/:id/profile", p.UserHandler.GetProfile)
			users.GET("/:id/roles", p.UserHandler.GetRoles)
		}
	}
	return r
}
