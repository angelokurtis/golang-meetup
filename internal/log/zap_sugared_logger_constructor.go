package log

import (
	"go.uber.org/zap"
	"log"
)

func NewSugaredLogger() *zap.SugaredLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	return logger.Sugar()
}
