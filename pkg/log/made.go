package log

import (
	"os"
	"runtime/debug"
)

func Debug(args ...interface{}) {
	mainLogger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	mainLogger.Debugf(format, args...)
}

func Info(args ...interface{}) {
	mainLogger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	mainLogger.Infof(format, args...)
}

func Warn(args ...interface{}) {
	mainLogger.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	mainLogger.Warnf(format, args...)
}

func Error(args ...interface{}) {
	mainLogger.Error(args...)
	os.Exit(1)
}

func Errorf(format string, args ...interface{}) {
	mainLogger.Errorf(format, args...)
	os.Exit(1)
}

func Panic(args ...interface{}) {
	mainLogger.Panic(args...)
}

func Fatal(args ...interface{}) {
	mainLogger.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	mainLogger.Fatalf(format, args...)
}

func LogPanicHandler(err interface{}) {
	Errorf("Error in Go routine: %s", err)
	Errorf("Stack: %s", debug.Stack())
}
