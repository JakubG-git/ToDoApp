package logging

import "go.uber.org/zap"

var Logger *zap.Logger

func ConfigureLogger() {
	// Configure logger
	logger, _ := zap.NewProduction()
	Logger = logger
}
