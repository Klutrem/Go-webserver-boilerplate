package repository

import (
	repository "main/repository/workspaces"

	"go.uber.org/fx"
)

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewKubernetesRepository),
	fx.Provide(repository.NewWorkspaceRepository),
)
