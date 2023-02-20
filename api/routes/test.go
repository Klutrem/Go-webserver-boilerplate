package routes

import (
	"main/api/kubescontrollers"
	"main/api/middlewares"
	"main/api/ws"
	"main/lib"

	"github.com/gin-gonic/gin"
)

// TestRoutes struct
type TestRoutes struct {
	logger          lib.Logger
	handler         lib.RequestHandler
	authMiddleware  middlewares.JWTAuthMiddleware
	kubescontroller kubescontrollers.KubeController
	Websocket       ws.Ws
}

// Setup Test routes
func (s TestRoutes) Setup() {
	s.logger.Info("Setting up routes")
	api := s.handler.Gin.Group("/api") //.Use(s.authMiddleware.Handler())
	{
		api.GET("/kube_get/:namespace", s.kubescontroller.GetPodList)
		api.GET("/helm_get", s.kubescontroller.HGetReleaseRequest)
		api.POST("/kube/add", s.kubescontroller.CreatePodRequest)
		api.POST("/kube/create_config_map", s.kubescontroller.CreateOrUpdateConfigMapRequest)
		api.POST("/kube/create_secret", s.kubescontroller.CreateOrUpdateSecretRequest)
		api.POST("/kube/create_namespace", s.kubescontroller.CreateNamespaceRequest)
		api.POST("/kube/create_pv", s.kubescontroller.CreatePersistentVolumeRequest)
		api.POST("/kube/create_pvc", s.kubescontroller.CreatePersistentVolumeClaimRequest)
		api.POST("/kube/create_nodeport", s.kubescontroller.CreateNodePortRequest)
		api.POST("/helm_create", s.kubescontroller.HCreateReleaseRequest)
		api.POST("helm_create_repo", s.kubescontroller.HCreateRepoRequest)
		api.POST("/kube/create_role", s.kubescontroller.CreateRoleRequest)
		api.POST("/kube/role_bind", s.kubescontroller.CreateRoleBindingRequest)
		api.POST("/kube/create_account", s.kubescontroller.CreateServiceAccountRequest)
		api.DELETE("/kube_delete/:namespace/:pod_name", s.kubescontroller.DeletePodRequest)
	}

	r := gin.Default()
	r.GET("/", s.Websocket.MessageHandler)
	go r.Run(":12121")

}

// NewTestRoutes creates new Test controller
func NewTestRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	kubescontroller kubescontrollers.KubeController,
) TestRoutes {
	return TestRoutes{
		handler:         handler,
		logger:          logger,
		kubescontroller: kubescontroller,
	}
}
