package main

import "github.com/xbmlz/guc/ulog"

func main() {
	ulog.StdLogger.Debug("This is a debug message.")
	ulog.StdLogger.Debugf("This is a debug message: %v", "debug")
	ulog.StdLogger.Info("This is a info message.")
	ulog.StdLogger.Infof("This is a info message: %v", "info")
	ulog.StdLogger.Warn("This is a warn message.")
	ulog.StdLogger.Warnf("This is a warn message: %v", "warn")
	ulog.StdLogger.Error("This is a error message.")
	ulog.StdLogger.Errorf("This is a error message: %v", "error")
}
