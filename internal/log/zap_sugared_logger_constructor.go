package log

import (
	"go.uber.org/zap"
	defaultLogger "log"
)

func NewSugaredLogger() *zap.SugaredLogger { // Provider
	logger, err := zap.NewDevelopment()
	if err != nil {
		defaultLogger.Fatal(err)
	}
	log := logger.Sugar()
	return log
}
