package routes

import (
	controllers "main/api/controllers/computing"
	"main/lib"
)

type ComputingRoutes struct {
	logger        lib.Logger
	handler       lib.RequestHandler
	PodController controllers.PodController
}

func (s ComputingRoutes) Setup() {
	s.logger.Info("Setting up computing routes")
	computing := s.handler.Gin.Group("/computing")

	computing.POST("/pod/create", s.PodController.CreatePodRequest)
	computing.POST("/vm/create", s.PodController.CreateVM)
}

func NewComputingRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	PodController controllers.PodController,
) ComputingRoutes {
	return ComputingRoutes{
		handler:       handler,
		logger:        logger,
		PodController: PodController,
	}
}
