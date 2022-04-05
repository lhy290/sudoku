package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var defaultLogger = logrus.New()

func InitLog(level int, file string) {
	defaultLogger.SetLevel(logrus.DebugLevel)
	defaultLogger.SetOutput(os.Stdout)
}

func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args)
}

func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args)
}

func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args)
}
