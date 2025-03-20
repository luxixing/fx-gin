package handler

import "go.uber.org/fx"

// Module 处理器模块
var Module = fx.Options(
	fx.Provide(
		NewTestHandler,
		NewConfigHandler,
		NewUserHandler,
	),
)
