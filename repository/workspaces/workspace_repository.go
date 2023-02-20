package repository

import (
	"context"
	"fmt"
	"main/lib"
	models "main/models/workspaces"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WorkspaceRepository struct {
	lib.Database
	logger lib.Logger
}

// NewTestRepository creates a new Test repository
func NewWorkspaceRepository(db lib.Database, logger lib.Logger) WorkspaceRepository {
	return WorkspaceRepository{
		Database: db,
		logger:   logger,
	}
}

func (s WorkspaceRepository) GetOneWorkspace(id primitive.ObjectID) (result models.Workspace, err error) {
	filter := bson.D{{Key: "_id", Value: id}}
	err = s.Database.Collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, err
}

func (s WorkspaceRepository) GetDeletedWorkspaces() ([]models.Workspaces, error) {
	filter := bson.D{{}}
	curr, err := s.Trash.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	var Tests []models.Workspace
	if err = curr.All(context.TODO(), &Tests); err != nil {
		return nil, err
	}
	deleted := make([]models.Workspaces, len(Tests))
	for i, test := range Tests {
		deleted[i].ID = test.ID
		deleted[i].Name = test.Name
		if test.Status {
			deleted[i].Status = "online"
		} else {
			deleted[i].Status = "offline"
		}
	}
	return deleted, err
}

func (s WorkspaceRepository) GetAllWorkspaces() ([]models.Workspaces, error) {
	filter := bson.D{{}}
	curr, err := s.Collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	var Tests []models.Workspace
	if err = curr.All(context.TODO(), &Tests); err != nil {
		return nil, err
	}
	response := make([]models.Workspaces, len(Tests))
	for i, test := range Tests {
		response[i].ID = test.ID
		response[i].Name = test.Name
		if test.Status {
			response[i].Status = "online"
		} else {
			response[i].Status = "offline"
		}
	}
	return response, err
}

func (s WorkspaceRepository) CreateWorkspace(Workspace models.Workspace) (primitive.ObjectID, error) {
	Workspace.ID = primitive.NewObjectID()
	Workspace.Status = true
	_, err := s.Collection.InsertOne(context.TODO(), Workspace)
	return Workspace.ID, err
}

func (s WorkspaceRepository) AddEdge(Edge models.Edge, id primitive.ObjectID) error {
	var workspace models.Workspace
	filter := bson.D{{Key: "_id", Value: id}}
	err := s.Collection.FindOne(context.TODO(), filter).Decode(&workspace)
	if err != nil {
		return err
	}
	Edge.ID = fmt.Sprintf("e%s-%s", Edge.Source, Edge.Target)
	workspace.Edges = append(workspace.Edges, Edge)
	update := bson.D{{Key: "$set", Value: workspace}}
	_, err = s.Collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (s WorkspaceRepository) DeleteWorkspace(id primitive.ObjectID) error {
	filter := bson.D{{Key: "_id", Value: id}}
	var workspace models.Workspace
	err := s.Collection.FindOne(context.TODO(), filter).Decode(&workspace)
	if err != nil {
		return err
	}
	_, err = s.Collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	workspace.Status = false
	_, err = s.Trash.InsertOne(context.TODO(), workspace)
	if err != nil {
		return err
	}
	return err
}

func (s WorkspaceRepository) DeleteNode(id primitive.ObjectID, node_id string) error {
	int_node_id, err := strconv.Atoi(node_id)
	if err != nil {
		return err
	}
	int_node_id -= 1
	filter := bson.D{{Key: "_id", Value: id}}
	var workspace models.Workspace
	err = s.Collection.FindOne(context.TODO(), filter, nil).Decode(&workspace)
	if err != nil {
		return err
	}
	workspace.Node = append(workspace.Node[:int_node_id], workspace.Node[int_node_id+1:]...)

	update := bson.D{{Key: "$set", Value: workspace}}
	_, err = s.Collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (s WorkspaceRepository) UpdateNode(id primitive.ObjectID, node models.Node) error {
	node_id, err := strconv.Atoi(node.ID)
	if err != nil {
		return err
	}
	node_id -= 1
	filter := bson.D{{Key: "_id", Value: id}}
	var workspace models.Workspace
	err = s.Collection.FindOne(context.TODO(), filter, nil).Decode(&workspace)
	if err != nil {
		return err
	}
	workspace.Node[node_id] = node
	update := bson.D{{Key: "$set", Value: workspace}}
	_, err = s.Collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (s WorkspaceRepository) AddNode(Node models.Node, id primitive.ObjectID) error {
	var workspace models.Workspace
	filter := bson.D{{Key: "_id", Value: id}}
	err := s.Collection.FindOne(context.TODO(), filter).Decode(&workspace)
	if err != nil {
		return err
	}
	Node.ID = fmt.Sprint(len(workspace.Node) + 1)
	int_id, _ := strconv.Atoi(Node.ID)
	for _, node := range workspace.Node {
		if Node.ID == node.ID {
			Node.ID = fmt.Sprint(int_id - 1)
		}
	}
	workspace.Node = append(workspace.Node, Node)
	update := bson.D{{Key: "$set", Value: workspace}}
	_, err = s.Collection.UpdateOne(context.TODO(), filter, update)
	return err
}
