package controllers

import (
	models "main/models/computing"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s PodController) CreateVM(c *gin.Context) {
	var VMBody models.VMBody
	if err := c.ShouldBindJSON(&VMBody); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	s.logger.Info(VMBody)
	err := s.service.CreateCRD(VMBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(200, "VM Created")
}
