package log

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	New(WithLevel(ErrorLevel),
		WithFormatter(&logrus.JSONFormatter{}),
		WithOutput(logrus.StandardLogger().Out)).
		With(F{"key": "value"}).
		WithFields(logrus.Fields{"key1": "value1", "key3": "value3"}).
		Error("hello")

	New(WithLevel(InfoLevel),
		WithFormatter(&logrus.JSONFormatter{})).
		With(F{"key": "value"}).
		Error("hello")

	logTofile1 := &lumberjack.Logger{
		Filename:   "./testlogrus.log",
		MaxSize:    500,  // 日志文件大小,单位是 MB
		MaxBackups: 3,    // 最大过期日志保留个数
		MaxAge:     28,   // 保留过期文件最大时间,单位 天
		Compress:   true, // 是否压缩日志,默认是不压缩,这里设置为true,压缩日志
	}
	logTofile2 := &lumberjack.Logger{
		Filename:   "./testlogrus.log",
		MaxSize:    500,  // 日志文件大小,单位是 MB
		MaxBackups: 3,    // 最大过期日志保留个数
		MaxAge:     28,   // 保留过期文件最大时间,单位 天
		Compress:   true, // 是否压缩日志,默认是不压缩,这里设置为true,压缩日志
	}
	New(WithLevel(InfoLevel),
		WithFormatter(&logrus.JSONFormatter{}), WithOutput(logTofile1, os.Stdout, logTofile2)).
		With(F{"key": "value"}).
		Error("hello")
}
