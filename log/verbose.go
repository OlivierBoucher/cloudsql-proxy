package log

import "log"

type verboseLogger struct {
	log *log.Logger
}

func (vl *verboseLogger) Debug(args ...interface{}) {
	vl.log.Println(args...)
}
func (vl *verboseLogger) Info(args ...interface{}) {
	vl.log.Println(args...)
}
func (vl *verboseLogger) Warn(args ...interface{}) {
	vl.log.Println(args...)
}
func (vl *verboseLogger) Error(args ...interface{}) {
	vl.log.Println(args...)
}
func (vl *verboseLogger) Panic(args ...interface{}) {
	vl.log.Panicln(args...)
}
func (vl *verboseLogger) Fatal(args ...interface{}) {
	vl.log.Fatalln(args...)
}

func (vl *verboseLogger) Debugf(t string, args ...interface{}) {
	vl.log.Printf(t, args...)
}
func (vl *verboseLogger) Infof(t string, args ...interface{}) {
	vl.log.Printf(t, args...)
}
func (vl *verboseLogger) Warnf(t string, args ...interface{}) {
	vl.log.Printf(t, args...)
}
func (vl *verboseLogger) Errorf(t string, args ...interface{}) {
	vl.log.Printf(t, args...)
}
func (vl *verboseLogger) Panicf(t string, args ...interface{}) {
	vl.log.Panicf(t, args...)
}
func (vl *verboseLogger) Fatalf(t string, args ...interface{}) {
	vl.log.Fatalf(t, args...)
}
