package log

import "log"

type silentLogger struct {
	log *log.Logger
}

func (vl *silentLogger) Debug(args ...interface{}) {}
func (vl *silentLogger) Info(args ...interface{})  {}
func (vl *silentLogger) Warn(args ...interface{})  {}

func (vl *silentLogger) Error(args ...interface{}) {
	vl.log.Println(args...)
}
func (vl *silentLogger) Panic(args ...interface{}) {
	vl.log.Panicln(args...)
}
func (vl *silentLogger) Fatal(args ...interface{}) {
	vl.log.Fatalln(args...)
}

func (vl *silentLogger) Debugf(t string, args ...interface{}) {}
func (vl *silentLogger) Infof(t string, args ...interface{})  {}
func (vl *silentLogger) Warnf(t string, args ...interface{})  {}

func (vl *silentLogger) Errorf(t string, args ...interface{}) {
	vl.log.Printf(t, args...)
}
func (vl *silentLogger) Panicf(t string, args ...interface{}) {
	vl.log.Panicf(t, args...)
}
func (vl *silentLogger) Fatalf(t string, args ...interface{}) {
	vl.log.Fatalf(t, args...)
}
