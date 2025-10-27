package glog

import (
	"os"
	"time"

	"github.com/phuslu/log"
)

type Logger struct {
	log.Logger
}

func New() *Logger {
	return &Logger{
		Logger: log.Logger{
			TimeLocation: time.UTC,
			Writer: &log.IOWriter{
				Writer: os.Stdout,
			},
		},
	}
}
