package logger

import "go.uber.org/zap"

// type Logger interface {
// 	Info(str string)
// 	Infof(str string, args interface{})
// }

type Logging struct {
	log *zap.Logger
}

func (l *Logging) Info(str string) {
	l.log.Info(str)
}

func (l *Logging) Infof(str string, args interface{}) {
	l.log.Sugar().Infof(str, args)
}
