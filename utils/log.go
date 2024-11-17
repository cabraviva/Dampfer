package utils

import "github.com/sirupsen/logrus"

// Register Logger
var Log = logrus.New()

func InitLogger() {
	Log.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	Log.SetLevel(logrus.InfoLevel)
}
