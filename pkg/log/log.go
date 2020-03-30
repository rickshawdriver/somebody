package log

import (
	"github.com/sirupsen/logrus"
	stdlog "log"
	"os"
	"path/filepath"
	"strings"
)

type Logger struct {
	LogLevel  string `toml:"log_level"`
	LogPath   string `toml:"log_path"`
	LogFormat string `toml:"log_format"`
}

var (
	mainLogger *logrus.Logger
)

func init() {
	mainLogger = logrus.New()
	logrus.SetOutput(os.Stdout)
}

func Initialize(logger *Logger) {
	stdlog.SetFlags(stdlog.Lshortfile | stdlog.LstdFlags)

	var levelStr string = "debug"
	if logger.LogLevel != "" {
		levelStr = strings.ToLower(logger.LogLevel)
	}

	level, err := logrus.ParseLevel(levelStr)
	if err != nil {
		mainLogger.Errorf("error set level: %v", err)
	}
	mainLogger.SetLevel(level)

	if logger.LogPath != "" && err != SetOutPut(logger.LogPath) {
		mainLogger.Errorf("set log path error: %s", err)
	}

	mainLogger.SetFormatter(&logrus.TextFormatter{DisableColors: false, FullTimestamp: true, DisableSorting: true})
	if strings.ToLower(logger.LogFormat) == "json" {
		mainLogger.SetFormatter(&logrus.JSONFormatter{})
	}
}

func SetOutPut(path string) error {
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		mainLogger.Errorf("create log path err : %s", err)
	}

	logFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	mainLogger.SetOutput(logFile)
	return nil
}
