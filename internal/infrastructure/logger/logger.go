package logger

import (
	"os"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

func New(environment string) *logrus.Logger {
	log := logrus.New()
	log.SetOutput(os.Stdout)

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC1123Z,
	})

	if environment == "production" {
		log.SetLevel(logrus.WarnLevel)
	} else {
		log.SetLevel(logrus.DebugLevel)
	}

	log.WithFields(logrus.Fields{
		"runtime_version": runtime.Version(),
		"cpu_count":       runtime.NumCPU(),
		"architecture":    runtime.GOARCH,
	}).Info("application initializing")

	return log
}
