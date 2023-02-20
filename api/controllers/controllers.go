package controllers

import (
	controllers "main/api/controllers/workspaces"

	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	fx.Provide(controllers.NewWorkspaceController),
)
