package klogger

type (
	Logger interface {
		Debug(format string, args ...interface{})
		Info(format string, args ...interface{})
		Warning(format string, args ...interface{})
		Error(format string, args ...interface{})
		Critical(format string, args ...interface{})
	}

	loggerFactory       func() Logger
	prefixLoggerFactory func(string) Logger
)

var (
	New           loggerFactory
	NewWithPrefix prefixLoggerFactory
)
