package logging

import (
	"github.com/sirupsen/logrus"
	"os"
)

func setupFormatter(logger *logrus.Logger, typeFormat Format) *logrus.Entry {
	var entry *logrus.Entry

	switch typeFormat {
	case FormatJSON:
		logger.SetOutput(os.Stdout)
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})

		entry = logrus.NewEntry(logger)
	case FormatText:
		logger.SetOutput(os.Stdout)
		logger.SetFormatter(&logrus.TextFormatter{
			ForceColors:     true,
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			DisableQuote:    true,
		})
		logger.SetReportCaller(true)
		entry = logrus.NewEntry(logger)

		entry.Logger.Formatter = &customFormatter{entry.Logger.Formatter}
	}

	return entry
}
