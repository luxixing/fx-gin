package http

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/luxixing/fx-gin/docs/swagger"
	"github.com/luxixing/fx-gin/internal/transport/http/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
)

// Module HTTP传输层模块
var Module = fx.Options(
	fx.Provide(NewRouter),
)

// RouterParams 路由参数
type RouterParams struct {
	fx.In

	TestHandler   *handler.TestHandler
	ConfigHandler *handler.ConfigHandler
	UserHandler   *handler.UserHandler
}

// Router 路由器
type Router struct {
	*gin.Engine
	testHandler   *handler.TestHandler
	configHandler *handler.ConfigHandler
	userHandler   *handler.UserHandler
}

// NewRouter 创建路由器
func NewRouter(p RouterParams) *Router {
	r := gin.New()

	// 添加第三方CORS中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 添加日志和恢复中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//swagger
	// 创建swagger文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/api/v1")
	{
		v1.GET("/test", p.TestHandler.Test)
		v1.POST("/config", p.ConfigHandler.Create)
		v1.GET("/config/:key", p.ConfigHandler.Get)

		// 添加用户相关路由
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

	return &Router{
		Engine:        r,
		testHandler:   p.TestHandler,
		configHandler: p.ConfigHandler,
		userHandler:   p.UserHandler,
	}
}
