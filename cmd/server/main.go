package main

import (
	"github.com/example/ember-gate/internal/app"
	"github.com/example/ember-gate/internal/config"
	"github.com/example/ember-gate/internal/db"
	"github.com/example/ember-gate/internal/logger"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logger.Init(cfg)
	dbConn, err := db.New(cfg)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	app.Run(cfg, dbConn)
}
