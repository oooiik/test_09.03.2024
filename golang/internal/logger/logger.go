package logger

import (
	"github.com/oooiik/test_09.03.2024/internal/config"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func confUp() {
	if Logger != nil {
		return
	}
	cfg := config.Load().Log

	Logger = logrus.New()
	Logger.SetOutput(cfg.Writer())

	level := logrus.InfoLevel
	switch cfg.Level {
	case "panic":
		level = logrus.PanicLevel
	case "fatal":
		level = logrus.FatalLevel
	case "error":
		level = logrus.ErrorLevel
	case "warn":
		level = logrus.WarnLevel
	case "info":
		level = logrus.InfoLevel
	case "debug":
		level = logrus.DebugLevel

	}
	Logger.SetLevel(level)
}

func Panic(args ...interface{}) {
	confUp()
	Logger.Panicln(args...)
}

func Fatal(args ...interface{}) {
	confUp()
	Logger.Fatalln(args...)
}

func Error(args ...interface{}) {
	confUp()
	Logger.Errorln(args...)
}
func Warn(args ...interface{}) {
	confUp()
	Logger.Warnln(args...)
}

func Info(args ...interface{}) {
	confUp()
	Logger.Infoln(args...)
}

func Debug(args ...interface{}) {
	confUp()
	Logger.Debug(args...)
}
