package ulog

import (
	"fmt"
	"io"
	"log"
)

// Level is a logger level.
type Level int8

const LevelKey = "level"

const (
	// LevelDebug is logger debug level.
	LevelDebug Level = iota
	// LevelInfo is logger info level.
	LevelInfo
	// LevelWarn is logger warn level.
	LevelWarn
	// LevelError is logger error level.
	LevelError
)

const (
	// LevelDebugPrefix is logger debug prefix. color: blue
	DebugPrefix = "\033[34m[DEBUG]\033[0m "
	// LevelInfoPrefix is logger info prefix. color: green
	InfoPrefix = "\033[32m[INFO ]\033[0m "
	// LevelWarnPrefix is logger warn prefix. color: yellow
	WarnPrefix = "\033[33m[WARN ]\033[0m "
	// LevelErrorPrefix is logger error prefix. color: red
	ErrorPrefix = "\033[31m[ERROR]\033[0m "
)

type Logger interface {
	SetLevel(level Level)
	Debug(v ...any)
	Debugf(format string, v ...any)
	Info(v ...any)
	Infof(format string, v ...any)
	Warn(v ...any)
	Warnf(format string, v ...any)
	Error(v ...any)
	Errorf(format string, v ...any)
}

type stdLogger struct {
	log   *log.Logger
	level Level
}

var _ Logger = (*stdLogger)(nil)

// StdLogger is default logger.
var StdLogger = NewStdLogger(log.Writer())

func NewStdLogger(w io.Writer) Logger {
	return &stdLogger{
		log: log.New(w, "", log.LstdFlags),
	}
}

func (l *stdLogger) SetLevel(level Level) {
	l.level = level
}

func (l *stdLogger) Debug(v ...interface{}) {
	if LevelDebug < l.level {
		return
	}
	l.log.SetPrefix(DebugPrefix)
	l.log.Output(2, fmt.Sprintln(v...))
}

func (l *stdLogger) Debugf(format string, v ...interface{}) {
	if LevelDebug < l.level {
		return
	}
	l.log.SetPrefix(DebugPrefix)
	l.log.Output(2, fmt.Sprintf(format, v...))
}

func (l *stdLogger) Info(v ...interface{}) {
	if LevelInfo < l.level {
		return
	}
	l.log.SetPrefix(InfoPrefix)
	l.log.Output(2, fmt.Sprintln(v...))
}

func (l *stdLogger) Infof(format string, v ...interface{}) {
	if LevelInfo < l.level {
		return
	}
	l.log.SetPrefix(InfoPrefix)
	l.log.Output(2, fmt.Sprintf(format, v...))
}

func (l *stdLogger) Warn(v ...interface{}) {
	if LevelWarn < l.level {
		return
	}
	l.log.SetPrefix(WarnPrefix)
	l.log.Output(2, fmt.Sprintln(v...))
}

func (l *stdLogger) Warnf(format string, v ...interface{}) {
	if LevelWarn < l.level {
		return
	}
	l.log.SetPrefix(WarnPrefix)
	l.log.Output(2, fmt.Sprintf(format, v...))
}

func (l *stdLogger) Error(v ...interface{}) {
	if LevelError < l.level {
		return
	}
	l.log.SetPrefix(ErrorPrefix)
	l.log.Output(2, fmt.Sprintln(v...))
}

func (l *stdLogger) Errorf(format string, v ...interface{}) {
	if LevelError < l.level {
		return
	}
	l.log.SetPrefix(ErrorPrefix)
	l.log.Output(2, fmt.Sprintf(format, v...))
}
