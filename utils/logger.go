package utils

import (
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func LogInit()  {
	log := logrus.New()
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetReportCaller(true)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	  })
	Log = log
}