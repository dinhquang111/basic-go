package logger

import "log"

type Logger interface {
	Info(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Debug(msg string, fields ...Field)
}

type Field struct {
	Key   string
	Value interface{}
}

func NewLogger() Logger {
	logger, err := newZapLogger()
	if err != nil {
		log.Fatal("Can not create logger")
	}
	return logger
}
