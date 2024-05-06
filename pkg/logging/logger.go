package logging

import (
	"ToDoApp/pkg/config"
	"go.uber.org/zap"
)

var Logger *zap.Logger

func ConfigureLogger(configMode config.Mode) {
	// Configure logger
	if configMode == config.Development {
		Logger, _ = zap.NewDevelopment()
	} else {
		Logger, _ = zap.NewProduction()
	}
}
