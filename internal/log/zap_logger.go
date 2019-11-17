package log

import (
	"fmt"
	"go.uber.org/zap"
)

type ZapLogger struct {
	zap.SugaredLogger
}

func NewZapLogger(sugaredLogger zap.SugaredLogger) *ZapLogger {
	return &ZapLogger{SugaredLogger: sugaredLogger}
}

func (z *ZapLogger) Info(msg string) {
	z.SugaredLogger.Info(msg)
}

func (z *ZapLogger) Debug(msg string) {
	z.SugaredLogger.Debug(msg)
}

func (z *ZapLogger) Error(msg string, err error) {
	z.Errorw(msg, "error", fmt.Sprintf("%+v", err))
}
