package logger

type Logger interface {
	Info(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Debug(msg string, fields ...Field)
}

type Field struct {
	Key   string
	Value interface{}
}

func SetupLogger() {
	SetupZeroLogger()
}
