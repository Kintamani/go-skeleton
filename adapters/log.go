package adapters

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

// NewLogger membuat logger baru
func NewLogger(stage string) *logrus.Logger {
	var log = logrus.New()

	logFilePath := func() string {
		now := time.Now()
		return filepath.Join("infrastructures/loggers/", now.Format("2006-01-02")+".log")
	}

	var (
		lummberjackLog = &lumberjack.Logger{
			Filename:   logFilePath(),
			MaxSize:    10,   // dalam megabytes
			MaxBackups: 3,    // jumlah file log lama yang dipertahankan
			MaxAge:     28,   // dalam hari
			Compress:   true, // kompres file log lama
		}

		writers = []io.Writer{log.Out, lummberjackLog}
		mw      = io.MultiWriter(writers...)
	)

	// menggunakan json format untuk production
	if stage == "production" {
		// log = log.New(lumberjackLogger).With().Timestamp().Caller().Logger().Level(zerolog.InfoLevel)

		log.SetLevel(logrus.WarnLevel)
		log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, TimestampFormat: time.RFC1123Z})
		logMultiWriter := io.MultiWriter(os.Stdout, lummberjackLog)
		log.SetOutput(logMultiWriter)
	} else {
		// log = zerolog.New(mw).With().Timestamp().Caller().Logger().Level(logLevel)

		log.SetLevel(logrus.Level(logrus.DebugLevel))
		log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, TimestampFormat: time.RFC1123Z})
		logMultiWriter := io.MultiWriter(os.Stdout, mw)
		log.SetOutput(logMultiWriter)
	}

	logrus.WithFields(logrus.Fields{
		"Runtime Version": runtime.Version(),
		"Number of CPUs":  runtime.NumCPU(),
		"Arch":            runtime.GOARCH,
	}).Info("Application Initializing")

	q := make(chan os.Signal, 1)
	c := make(chan os.Signal, 1)
	signal.Notify(q, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	signal.Notify(c, syscall.SIGHUP)
	go func() {
		for {
			<-q
			lummberjackLog.Close()
			logrus.Info("Closing logs ...")
		}
	}()
	go func() {
		for {
			<-c
			if err := lummberjackLog.Rotate(); err != nil {
				logrus.Error("Error while rotating logs")
			}
			logrus.Info("Rotating logs ...")
		}
	}()

	return log
}
