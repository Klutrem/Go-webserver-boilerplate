package controllers

import (
	"main/lib"
	computing_models "main/models/computing"
	"main/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PodController struct {
	service services.KubernetesService
	logger  lib.Logger
}

func NewPodController(KubeService services.KubernetesService, logger lib.Logger) PodController {
	return PodController{
		service: KubeService,
		logger:  logger,
	}
}

func (u PodController) CreatePodRequest(c *gin.Context) {
	body := computing_models.PodBody{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	pod, err := u.service.CreatePod(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(200, pod)
}
