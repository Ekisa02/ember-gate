package logger

import (
	"github.com/example/ember-gate/internal/config"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func Init(cfg *config.Config) {
	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		level = logrus.InfoLevel
	}
	log.SetLevel(level)
	log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
}

func L() *logrus.Logger { return log }
