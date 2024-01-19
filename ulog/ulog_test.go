package ulog

import (
	"testing"
)

func TestSetLevel(t *testing.T) {
	StdLogger.SetLevel(LevelDebug)
	StdLogger.SetLevel(LevelInfo)
	StdLogger.SetLevel(LevelWarn)
	StdLogger.SetLevel(LevelError)
}

func TestDebug(t *testing.T) {
	StdLogger.Debug("debug")
	StdLogger.Debugf("debugf: %v", "debug")
}

func TestInfo(t *testing.T) {
	StdLogger.Info("info")
	StdLogger.Infof("infof: %v", "info")
}

func TestWarn(t *testing.T) {
	StdLogger.Warn("warn")
	StdLogger.Warnf("warnf: %v", "warn")
}

func TestError(t *testing.T) {
	StdLogger.Error("error")
	StdLogger.Errorf("errorf: %v", "error")
}
