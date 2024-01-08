package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

const (
	defaultLogLevel = logrus.DebugLevel
	red             = 31
	green           = 32
	yellow          = 33
	blue            = 36
	gray            = 37
)

var (
	entry *logrus.Entry
)

type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger {
	return &Logger{entry}
}

func init() {
	logger := logrus.New()
	switch os.Getenv("MODE") {
	case "dev":
		entry = setupFormatter(logger, FormatJSON)
		logger.SetLevel(defaultLogLevel)
	case "prod":
		entry = setupFormatter(logger, FormatJSON)
		logger.SetLevel(logrus.InfoLevel)
	default:
		entry = setupFormatter(logger, FormatText)
		logger.SetLevel(defaultLogLevel)
	}
}

type customFormatter struct {
	logrus.Formatter
}

func (f *customFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	filePath := ""
	if entry.HasCaller() {
		caller := entry.Caller
		filePath = fmt.Sprintf("%s:%d", caller.File, caller.Line)
	}

	color := getColorByLevel(entry.Level)
	message := fmt.Sprintf("\033[%dm%s [%s] [%s] %s\033[0m\n", color, entry.Time.Format("2006-01-02 15:04:05"), strings.ToUpper(entry.Level.String()), filePath, entry.Message)
	return []byte(message), nil
}

func getColorByLevel(level logrus.Level) int {
	switch level {
	case logrus.DebugLevel, logrus.TraceLevel:
		return gray
	case logrus.InfoLevel:
		return blue
	case logrus.WarnLevel:
		return yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		return red
	default:
		return blue
	}
}
