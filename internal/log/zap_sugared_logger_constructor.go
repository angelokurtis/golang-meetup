package log

import (
	"go.uber.org/zap"
	defaultLogger "log"
)

func NewSugaredLogger() *zap.SugaredLogger { // Wire's Provider
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeCaller = nil
	logger, err := config.Build()
	if err != nil {
		defaultLogger.Fatal(err)
	}
	log := logger.Sugar()
	return log
}
