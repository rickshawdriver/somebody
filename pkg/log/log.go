package log

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"strings"
)

type Log struct {
	LogLevel string `toml:"log_level"`
}

var (
	log = logrus.New()
)

func init() {
	formatter := new(prefixed.TextFormatter)
	formatter.FullTimestamp = true

	log.Formatter = formatter
}

func Get() *logrus.Logger {
	return log
}

func SetLogLevel(level string) {
	switch strings.ToLower(level) {
	case "error":
		log.Level = logrus.ErrorLevel
	case "warn":
		log.Level = logrus.WarnLevel
	case "debug":
		log.Level = logrus.DebugLevel
	default:
		log.Level = logrus.InfoLevel
	}
}
