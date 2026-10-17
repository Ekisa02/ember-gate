package services

import (
	"errors"

	"github.com/example/ember-gate/internal/models"
	"gorm.io/gorm"
)

type NodeService struct {
	DB *gorm.DB
}

func NewNodeService(db *gorm.DB) *NodeService { return &NodeService{DB: db} }

func (s *NodeService) Create(node *models.Node) error {
	if node.Name == "" || node.URL == "" {
		return errors.New("name and url are required")
	}
	return s.DB.Create(node).Error
}

func (s *NodeService) List() ([]models.Node, error) {
	var nodes []models.Node
	if err := s.DB.Order("created_at desc").Find(&nodes).Error; err != nil {
		return nil, err
	}
	return nodes, nil
}

func (s *NodeService) Get(id uint) (*models.Node, error) {
	var node models.Node
	if err := s.DB.First(&node, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &node, nil
}
