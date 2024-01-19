package ulog

import (
	"log"
	"testing"
)

func TestNewStdLogger(t *testing.T) {
	NewStdLogger(log.Writer())
}

func TestSetLevel(t *testing.T) {
	StdLogger.SetLevel(LevelDebug)
	StdLogger.SetLevel(LevelInfo)
	StdLogger.SetLevel(LevelWarn)
	StdLogger.SetLevel(LevelError)
}

func TestDebug(t *testing.T) {
	StdLogger.Debug("debug")
}

func TestDebugf(t *testing.T) {
	StdLogger.Debugf("debugf")
}

func TestInfo(t *testing.T) {
	StdLogger.Info("info")
}

func TestInfof(t *testing.T) {
	StdLogger.Infof("infof")
}

func TestWarn(t *testing.T) {
	StdLogger.Warn("warn")
}

func TestWarnf(t *testing.T) {
	StdLogger.Warnf("warnf")
}

func TestError(t *testing.T) {
	StdLogger.Error("error")
}

func TestErrorf(t *testing.T) {
	StdLogger.Errorf("errorf")
}
