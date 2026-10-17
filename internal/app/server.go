package app

import (
	"fmt"

	"github.com/example/ember-gate/internal/config"
	"github.com/example/ember-gate/internal/handlers"
	"github.com/example/ember-gate/internal/middleware"
	"github.com/example/ember-gate/internal/services"
	"github.com/example/ember-gate/internal/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Run(cfg *config.Config, db *gorm.DB) {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// services
	authSvc := services.NewAuthService(db, cfg.JWTSecret)
	nodeSvc := services.NewNodeService(db)

	// handlers
	authH := handlers.NewAuthHandler(authSvc)
	nodeH := handlers.NewNodeHandler(nodeSvc)

	v1 := r.Group("/api/v1")
	{
		v1.POST("/register", authH.Register)
		v1.POST("/login", authH.Login)

		v1.GET("/nodes", nodeH.List)
		v1.GET("/nodes/:id", nodeH.Get)

		authGroup := v1.Group("")
		authGroup.Use(middleware.JWTAuth(cfg.JWTSecret))
		{
			authGroup.POST("/nodes", nodeH.Create)
		}
	}

	addr := fmt.Sprintf(":%s", cfg.Port)
	logger.L().Infof("starting ember-gate on %s", addr)
	r.Run(addr)
}
