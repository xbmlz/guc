package ulog

import (
	"io"
	"log"
)

var _ Logger = (*stdLogger)(nil)

type stdLogger struct {
	log *log.Logger
}

func (l *stdLogger) Debug(v ...any) {
	l.log.Print(v...)
}

func (l *stdLogger) Debugf(format string, v ...any) {
	l.log.Printf(format, v...)
}

func (l *stdLogger) Info(v ...any) {
	l.log.Print(v...)
}

func (l *stdLogger) Infof(format string, v ...any) {
	l.log.Printf(format, v...)
}

func (l *stdLogger) Warn(v ...any) {
	l.log.Print(v...)
}

func (l *stdLogger) Warnf(format string, v ...any) {
	l.log.Printf(format, v...)
}

func (l *stdLogger) Error(v ...any) {
	l.log.Print(v...)
}

func (l *stdLogger) Errorf(format string, v ...any) {
	l.log.Printf(format, v...)
}

// NewStdLogger new std logger
func NewStdLogger(w io.Writer) Logger {
	return &stdLogger{
		log: log.New(w, "", log.LstdFlags),
	}
}
