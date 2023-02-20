package services

import (
	"main/lib"
	models "main/models/workspaces"
	repository "main/repository/workspaces"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// WorkspaceService service layer
type WorkspaceService struct {
	logger     lib.Logger
	repository repository.WorkspaceRepository
}

// NewWorkspaceService creates a new Testservice
func NewWorkspaceService(logger lib.Logger, repository repository.WorkspaceRepository) WorkspaceService {
	return WorkspaceService{
		logger:     logger,
		repository: repository,
	}
}

// GetOneTest gets one Test
func (s WorkspaceService) GetOneWorkspace(id primitive.ObjectID) (result models.Workspace, err error) {
	return s.repository.GetOneWorkspace(id)
}

func (s WorkspaceService) GetDeletedWorkspaces() ([]models.Workspaces, error) {
	return s.repository.GetDeletedWorkspaces()
}

// GetAllTest get all the Test
func (s WorkspaceService) GetAllWorkspaces() ([]models.Workspaces, error) {
	return s.repository.GetAllWorkspaces()
}

// CreateTest call to create the Test
func (s WorkspaceService) CreateWorkspace(Workspace models.Workspace) (primitive.ObjectID, error) {
	return s.repository.CreateWorkspace(Workspace)
}

// UpdateTest updates the Test
func (s WorkspaceService) AddNode(Node models.Node, id primitive.ObjectID) error {
	return s.repository.AddNode(Node, id)
}

func (s WorkspaceService) AddEdge(Edge models.Edge, id primitive.ObjectID) error {
	return s.repository.AddEdge(Edge, id)
}

// DeleteTest deletes the Test
func (s WorkspaceService) DeleteWorkspace(id primitive.ObjectID) error {
	return s.repository.DeleteWorkspace(id)
}

func (s WorkspaceService) DeleteNode(workspace_id primitive.ObjectID, id string) error {
	return s.repository.DeleteNode(workspace_id, id)
}

func (s WorkspaceService) UpdateNode(workspace_id primitive.ObjectID, node models.Node) error {
	return s.repository.UpdateNode(workspace_id, node)
}
