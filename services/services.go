package services

import (
	services "main/services/workspaces"

	"go.uber.org/fx"
)

// Module exports services present
var Module = fx.Options(
	fx.Provide(services.NewWorkspaceService),
	fx.Provide(NewKubernetesService),
)
