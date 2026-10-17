package handlers

import (
	"net/http"
	"strconv"

	"github.com/example/ember-gate/internal/middleware"
	"github.com/example/ember-gate/internal/models"
	"github.com/example/ember-gate/internal/services"
	"github.com/gin-gonic/gin"
)

type NodeHandler struct {
	Svc *services.NodeService
}

func NewNodeHandler(svc *services.NodeService) *NodeHandler { return &NodeHandler{Svc: svc} }

func (h *NodeHandler) Create(c *gin.Context) {
	var req models.Node
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if uid, ok := c.Get(middleware.ContextUserID); ok {
		req.OwnerID = uid.(uint)
	}
	if err := h.Svc.Create(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, req)
}

func (h *NodeHandler) List(c *gin.Context) {
	nodes, err := h.Svc.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nodes)
}

func (h *NodeHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	node, err := h.Svc.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if node == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, node)
}
