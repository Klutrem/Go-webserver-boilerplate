package controllers

import (
	"errors"
	"main/lib"
	models "main/models/workspaces"
	services "main/services/workspaces"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TestController data type
type WorkspaceController struct {
	service services.WorkspaceService
	logger  lib.Logger
}

// NewTestController creates new Test controller
func NewWorkspaceController(TestService services.WorkspaceService, logger lib.Logger) WorkspaceController {
	return WorkspaceController{
		service: TestService,
		logger:  logger,
	}
}

func (u WorkspaceController) ValidateCardLabel(cardLabel string) error {
	allowedNames := []string{
		"PV",
		"VM",
		"pod",
		"ingress",
		"service",
		"storage",
		"claim",
		"endpoints",
		"nodejs",
		"PVC",
		"rules",
	}
	for _, label := range allowedNames {
		if strings.EqualFold(cardLabel, label) {
			return nil
		}
	}
	return errors.New("invalid CardLabel")
}

func (u WorkspaceController) ValidateCpuNumber(n int) error {
	if 1 <= n && n <= 16 {
		return nil
	}
	return errors.New("CpuNumber must be between 1 and 16")
}

func (u WorkspaceController) ValidateMemoryNumber(n int) error {
	if 8 <= n && n <= 64 {
		return nil
	}
	return errors.New("MemoryNumber must be between 8 and 64")
}

func (u WorkspaceController) ValidateStorageNumber(n int) error {
	if 8 <= n && n <= 64 {
		return nil
	}
	return errors.New("StorageNumber must be between 8 and 64")
}

func (u WorkspaceController) ValidateNode(n models.Node) error {
	err := u.ValidateCardLabel(n.CardLabel)
	if err != nil {
		return err
	}
	err = u.ValidateCpuNumber(n.CpuNumber)
	if err != nil {
		return err
	}
	err = u.ValidateMemoryNumber(n.MemoryNumber)
	if err != nil {
		return err
	}
	err = u.ValidateStorageNumber(n.StorageNumber)
	if err != nil {
		return err
	}
	return nil
}

// @Summary Gets one test
// @Tags get tests
// @Description Get one test by id
// @Param id path string true "Test id"
// @Produce json
// @Security ApiKeyAuth
// @Router /api/test/{id} [get]
func (u WorkspaceController) GetOneWorkspace(c *gin.Context) {
	paramID := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(paramID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	Test, err := u.service.GetOneWorkspace(objID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(200, gin.H{
		"nodes": Test.Node,
		"edges": Test.Edges,
	})

}

// @Summary Get all test
// @Tags get tests
// @Description Get all the Tests
// @Accept */*
// @Security ApiKeyAuth
// @Router /api/test [get]
func (u WorkspaceController) GetWorkspaces(c *gin.Context) {
	Tests, err := u.service.GetAllWorkspaces()
	if err != nil || Tests == nil {
		c.JSON(http.StatusNotFound, err)
	}

	c.JSON(200, Tests)

}

func (u WorkspaceController) GetDeletedWorkspaces(c *gin.Context) {
	Deleted, err := u.service.GetDeletedWorkspaces()
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}
	c.JSON(200, Deleted)
}

// @Summary Create GetTests
// @Tags create test
// @Description Create new test
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param input body string true "test data"
// @Router /api/test [post]
func (u WorkspaceController) CreateWorkspace(c *gin.Context) {
	Workspace := models.Workspace{}
	//trxHandle := c.MustGet(constants.DBTransaction).(*gorm.DB)

	if err := c.ShouldBindJSON(&Workspace); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	id, err := u.service.CreateWorkspace(Workspace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"workspace_id": id,
	})
}

// @Summary Update test
// @Tags update test
// @Description Update an old test
// @Param id path int true "Test id"
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param input body string true "test data"
// @Router /api/test/{id} [post]
func (u WorkspaceController) AddNode(c *gin.Context) {
	paramID := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(paramID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	Node := models.Node{}
	if err := c.ShouldBindJSON(&Node); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := u.ValidateNode(Node); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := u.service.AddNode(Node, objID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, "Workspace updated")
}

func (u WorkspaceController) AddEdge(c *gin.Context) {
	paramID := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(paramID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	Edge := models.Edge{}
	if err := c.ShouldBindJSON(&Edge); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := u.service.AddEdge(Edge, objID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, "Workspace updated")

}

// @Summary delete test
// @Tags Delete
// @Description delete test
// @ID delete-test
// @Param id path int true "Test id"
// @Produce json
// @Security ApiKeyAuth
// @Router /api/test/{id} [delete]
func (u WorkspaceController) DeleteWorkspace(c *gin.Context) {
	paramID := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(paramID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	if err := u.service.DeleteWorkspace(objID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, "Test deleted")
}

func (u WorkspaceController) DeleteNode(c *gin.Context) {
	workspace_id := c.Param("id")
	node_id := c.Param("node_id")
	objID, err := primitive.ObjectIDFromHex(workspace_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	if err := u.service.DeleteNode(objID, node_id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, "Node deleted")
}

func (u WorkspaceController) UpdateNode(c *gin.Context) {
	Node := models.Node{}
	workspace_id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(workspace_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	if err := c.ShouldBindJSON(&Node); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := u.ValidateNode(Node); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := u.service.UpdateNode(objID, Node); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, "node updated")

}

// func (u WorkspaceController) GetCode(c *gin.Context) {
// 	paramcode := c.Query("code")
// 	Url := "http://localhost:8080/realms/master/protocol/openid-connect/token"
// 	url_form := url.Values{
// 		"grant_type":    {"authorization_code"},
// 		"code":          {paramcode},
// 		"client_id":     {"skyfarm"},
// 		"redirect_uri":  {"http://localhost:3000/get_code"},
// 		"client_secret": {os.Getenv("JWT_SECRET")},
// 	}
// 	resp, err := http.PostForm(Url, url_form)
// 	if err != nil {
// 		c.String(500, err.Error())
// 	}

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		c.String(500, err.Error())
// 	}
// 	resp.Body.Close()
// 	c.JSON(200, strings.Split(strings.Split(string(body), ",")[0], ":")[1])
// 	u.logger.Info(resp)
// }
