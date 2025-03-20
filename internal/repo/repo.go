package repo

import "go.uber.org/fx"

var (
	//register repo
	Module = fx.Options(
		fx.Provide(
			NewConfigRepo,
			NewUserRepo,
			NewRoleRepo,
			NewProfileRepo,
		),
	)
)