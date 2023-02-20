package routes

import (
	computing "main/api/routes/computing"

	workspaces "main/api/routes/workspaces"

	"go.uber.org/fx"
)

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewTestRoutes),
	fx.Provide(NewRoutes),
	fx.Provide(workspaces.NewWorkspaceRoutes),
	fx.Provide(computing.NewComputingRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(
	TestRoutes TestRoutes,
	WorkspaceRoutes workspaces.WorkspaceRoutes,
	ComputingRoutes computing.ComputingRoutes,

) Routes {
	return Routes{
		TestRoutes,
		WorkspaceRoutes,
		ComputingRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
