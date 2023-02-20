package routes

import (
	controllers "main/api/controllers/workspaces"
	"main/api/middlewares"
	"main/lib"
)

type WorkspaceRoutes struct {
	logger              lib.Logger
	handler             lib.RequestHandler
	WorkspaceController controllers.WorkspaceController
	authMiddleware      middlewares.JWTAuthMiddleware
}

func (s WorkspaceRoutes) Setup() {
	s.logger.Info("setting up worksapce routes")
	workspace := s.handler.Gin.Group("/workspace") //.Use(s.authMiddleware.Handler())
	workspace.GET("/", s.WorkspaceController.GetWorkspaces)
	workspace.GET("/:id", s.WorkspaceController.GetOneWorkspace)
	workspace.GET("/trash", s.WorkspaceController.GetDeletedWorkspaces)
	workspace.POST("/create", s.WorkspaceController.CreateWorkspace)
	workspace.POST("/:id/node/add", s.WorkspaceController.AddNode)
	workspace.POST("/:id/edge/add", s.WorkspaceController.AddEdge)
	workspace.POST("/:id/node/update", s.WorkspaceController.UpdateNode)
	workspace.DELETE("/:id/:node_id", s.WorkspaceController.DeleteNode)
	workspace.DELETE("/:id", s.WorkspaceController.DeleteWorkspace)
}

func NewWorkspaceRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	WorkspaceController controllers.WorkspaceController,
) WorkspaceRoutes {
	return WorkspaceRoutes{
		handler:             handler,
		logger:              logger,
		WorkspaceController: WorkspaceController,
	}
}
