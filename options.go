package log

import (
	"github.com/sirupsen/logrus"
	"io"
)

type Option func(logger *logrus.Logger)

func WithLevel(lvl Level) Option {
	return func(log *logrus.Logger) {
		switch lvl {
		case InfoLevel:
			log.Level = logrus.InfoLevel
		case FatalLevel:
			log.Level = logrus.FatalLevel
		case DebugLevel:
			log.Level = logrus.DebugLevel
		case ErrorLevel:
			log.Level = logrus.ErrorLevel
		case TraceLevel:
			log.Level = logrus.TraceLevel
		case WarnLevel:
			log.Level = logrus.WarnLevel
		default:
			panic(lvl)
		}
	}
}
func WithFormatter(f logrus.Formatter) Option {
	if f == nil {
		f = &logrus.TextFormatter{}
	}
	return func(log *logrus.Logger) {
		log.Formatter = f
	}
}
func WithLevelHooks(hs LevelHooks) Option {
	if hs == nil {
		hs = make(LevelHooks)
	}
	return func(log *logrus.Logger) {
		log.Hooks = hs
	}
}
func WithOutput(writers ...io.Writer) Option {
	return func(log *logrus.Logger) {
		log.SetOutput(io.MultiWriter(writers...))
	}
}
func withDefaults(opt []Option) []Option {
	return append([]Option{
		WithLevel(InfoLevel),
	}, opt...)
}
