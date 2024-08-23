package Log

import (
	"github.com/sirupsen/logrus"
	"os"
)

var (
	LogFile *os.File
	Logger  *logrus.Logger
)

func InitLogger() {
	var err error
	LogFile, err = os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("Failed to open log file: %v", err)
	}

	Logger = logrus.New()
	Logger.SetOutput(LogFile)
	Logger.SetFormatter(&logrus.JSONFormatter{})
	Logger.SetLevel(logrus.InfoLevel)
}

func LogInfo(message string, fields logrus.Fields) {
	Logger.WithFields(fields).Info(message)
}

func LogWarn(message string, fields logrus.Fields) {
	Logger.WithFields(fields).Warn(message)
}

func LogError(message string, fields logrus.Fields) {
	Logger.WithFields(fields).Error(message)
}
