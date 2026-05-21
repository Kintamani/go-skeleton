package logger

import (
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func New(environment string) *logrus.Logger {
	log := logrus.New()

	logDir := "logs"
	if err := os.MkdirAll(logDir, 0o755); err != nil {
		log.WithError(err).Warn("failed to create log directory")
	}

	logFilePath := filepath.Join(logDir, time.Now().Format("2006-01-02")+".log")
	rollingFile := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	}

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC1123Z,
	})

	if environment == "production" {
		log.SetLevel(logrus.WarnLevel)
		log.SetOutput(io.MultiWriter(os.Stdout, rollingFile))
	} else {
		log.SetLevel(logrus.DebugLevel)
		log.SetOutput(io.MultiWriter(os.Stdout, rollingFile))
	}

	log.WithFields(logrus.Fields{
		"runtime_version": runtime.Version(),
		"cpu_count":       runtime.NumCPU(),
		"architecture":    runtime.GOARCH,
	}).Info("application initializing")

	quitSignals := make(chan os.Signal, 1)
	rotateSignals := make(chan os.Signal, 1)

	signal.Notify(quitSignals, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	signal.Notify(rotateSignals, syscall.SIGHUP)

	go func() {
		for range quitSignals {
			_ = rollingFile.Close()
			log.Info("closing logs")
		}
	}()

	go func() {
		for range rotateSignals {
			if err := rollingFile.Rotate(); err != nil {
				log.WithError(err).Error("failed to rotate logs")
				continue
			}
			log.Info("rotating logs")
		}
	}()

	return log
}
