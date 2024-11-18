package log

import (
	"github.com/sirupsen/logrus"
)

type Log struct {
	entry *logrus.Entry
}

var origLog = logrus.New()
var baseLog = Log{logrus.NewEntry(origLog)}

const (
	FatalLevel Level = iota
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

type Level uint8

// LevelHooks is a collection of hooks that are synchronously
// triggered for each logging event.
type LevelHooks = logrus.LevelHooks

// F is a set of fields
type F map[string]interface{}

type Loggable interface {
	Loggable() map[string]interface{}
}

// Loggable allows Logger.With to consume an F.
func (f F) Loggable() map[string]interface{} {
	return f
}

type ILog interface {
	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})

	Error(...interface{})
	Errorf(string, ...interface{})
	Errorln(...interface{})

	Warn(...interface{})
	Warnf(string, ...interface{})
	Warnln(...interface{})

	Info(...interface{})
	Infof(string, ...interface{})
	Infoln(...interface{})

	Debug(...interface{})
	Debugf(string, ...interface{})
	Debugln(...interface{})

	Trace(...interface{})
	Tracef(string, ...interface{})
	Traceln(...interface{})

	With(Loggable) ILog
	WithError(error) ILog
	WithField(string, interface{}) ILog
	WithFields(logrus.Fields) ILog
}

func (l Log) Fatal(v ...interface{})              { l.entry.Fatal(v...) }
func (l Log) Fatalf(fmt string, v ...interface{}) { l.entry.Fatalf(fmt, v...) }
func (l Log) Fatalln(v ...interface{})            { l.entry.Fatalln(v...) }

func (l Log) Trace(v ...interface{})              { l.entry.Trace(v...) }
func (l Log) Tracef(fmt string, v ...interface{}) { l.entry.Tracef(fmt, v...) }
func (l Log) Traceln(v ...interface{})            { l.entry.Traceln(v...) }

func (l Log) Debug(v ...interface{})              { l.entry.Debug(v...) }
func (l Log) Debugf(fmt string, v ...interface{}) { l.entry.Debugf(fmt, v...) }
func (l Log) Debugln(v ...interface{})            { l.entry.Debugln(v...) }

func (l Log) Info(v ...interface{})              { l.entry.Info(v...) }
func (l Log) Infof(fmt string, v ...interface{}) { l.entry.Infof(fmt, v...) }
func (l Log) Infoln(v ...interface{})            { l.entry.Infoln(v...) }

func (l Log) Warn(v ...interface{})              { l.entry.Warn(v...) }
func (l Log) Warnf(fmt string, v ...interface{}) { l.entry.Warnf(fmt, v...) }
func (l Log) Warnln(v ...interface{})            { l.entry.Warnln(v...) }

func (l Log) Error(v ...interface{})              { l.entry.Error(v...) }
func (l Log) Errorf(fmt string, v ...interface{}) { l.entry.Errorf(fmt, v...) }
func (l Log) Errorln(v ...interface{})            { l.entry.Errorln(v...) }

func (l Log) With(v Loggable) ILog {
	return l.WithFields(v.Loggable())
}

func (l Log) WithError(err error) ILog {
	return Log{l.entry.WithError(err)}
}

func (l Log) WithField(k string, v interface{}) ILog {
	return Log{l.entry.WithField(k, v)}
}

func (l Log) WithFields(f logrus.Fields) ILog {
	return Log{l.entry.WithFields(f)}
}

func New(opt ...Option) ILog {
	for _, option := range withDefaults(opt) {
		option(origLog)
	}
	return Log{logrus.NewEntry(origLog)}
}

func WrapLogrus() ILog {
	return baseLog
}
