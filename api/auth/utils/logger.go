package utils

import (
	"go.uber.org/zap"
)

// logger setup here:
func NewLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	return logger
}
