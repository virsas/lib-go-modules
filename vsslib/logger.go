package vsslib

import (
	"go.uber.org/zap"
)

type LoggerHandler interface {
	Panic(msg string)
	Error(msg string)
}

type logger struct {
	zap *zap.Logger
}

func NewZapLogger() (LoggerHandler, error) {
	log := &logger{}

	zap, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	log.zap = zap

	return log, err
}

func (l *logger) Panic(msg string) {
	l.zap.Panic(msg)
}
func (l *logger) Error(msg string) {
	l.zap.Error(msg)
}
