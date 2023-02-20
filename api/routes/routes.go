package routes

import (
	routes "main/api/routes/workspaces"

	"go.uber.org/fx"
)

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewTestRoutes),
	fx.Provide(NewRoutes),
	fx.Provide(routes.NewWorkspaceRoutes),
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
	WorkspaceRoutes routes.WorkspaceRoutes,

) Routes {
	return Routes{
		TestRoutes,
		WorkspaceRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
