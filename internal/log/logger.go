package log

type Logger interface {
	Info(msg string)
	Debug(msg string)
	Error(msg string, err error)
}
