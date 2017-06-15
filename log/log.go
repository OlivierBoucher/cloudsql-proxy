package log

import (
	"log"
	"os"
)

// Use golang's standard logger by default.
// Access is not mutex-protected: do not modify except in init()
// functions.
var logger Logger = &silentLogger{log: log.New(os.Stderr, "", log.LstdFlags)}

// Logger mimics most leveled logging libraries' Logger as an interface.
type Logger interface {
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Panic(...interface{})
	Fatal(...interface{})
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	Panicf(string, ...interface{})
	Fatalf(string, ...interface{})
}

// SetVerbose sets the verbose logger. Call only from
// init() functions.
func SetVerbose() {
	logger = &verboseLogger{log: log.New(os.Stderr, "", log.LstdFlags)}
}

// SetLogger sets the logger that is used in cloudsql-proxy. Call only from
// init() functions.
func SetLogger(l Logger) {
	logger = l
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Debugf(t string, args ...interface{}) {
	logger.Debugf(t, args...)
}

func Infof(t string, args ...interface{}) {
	logger.Infof(t, args...)
}

func Warnf(t string, args ...interface{}) {
	logger.Warnf(t, args...)
}

func Errorf(t string, args ...interface{}) {
	logger.Errorf(t, args...)
}

func Panicf(t string, args ...interface{}) {
	logger.Panicf(t, args...)
}

func Fatalf(t string, args ...interface{}) {
	logger.Fatalf(t, args...)
}
