package repo

import "go.uber.org/fx"

var (
	//register repo
	Module = fx.Options(
		fx.Provide(
			NewConfigRepo,
		),
		fx.Provide(
			NewUserRepo,
		),
		fx.Provide(
			NewRoleRepo,
		),
		fx.Provide(
			NewProfileRepo,
		),
	)
)