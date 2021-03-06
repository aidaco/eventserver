package log

import (
	"io"
	"log"
	"os"
	path "path/filepath"
)

type DefaultLogger struct {
	logger *log.Logger
}

func (l *DefaultLogger) log(prefix string, msgs ...interface{}) {
	l.logger.SetPrefix(prefix)
	l.logger.Println(msgs)
}

func (l *DefaultLogger) Info(msgs ...interface{}) {
	l.log("INFO:", msgs)
}

func (l *DefaultLogger) Warn(msg ...interface{}) {
	l.log("WARN:", msg)
}

func (l *DefaultLogger) Error(msg ...interface{}) {
	l.log("ERROR:", msg)
}

func NewDefaultLogger() *DefaultLogger {
	eslogdir := os.Getenv("esLOGDIR")
	if eslogdir == "" {
		eslogdir = esLOGDIR
	}

	logdir, err := path.Abs(eslogdir)
	if err != nil {
		log.Fatalln("Could not start logger: unable to open log file. \nDetails:", err)
	}
	logpath := path.Join(logdir, "LOG.txt")
	logfile, err := os.OpenFile(logpath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("Could not start logger: unable to open log file. \nDetails:", err)
	}

	writer := io.MultiWriter(os.Stdout, logfile)
	l := log.New(writer, "", log.LstdFlags|log.Lshortfile)

	return &DefaultLogger{l}
}
