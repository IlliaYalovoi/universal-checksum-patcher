package main

import "log"

type logger struct {
	*log.Logger
}

func newLogger() *logger {
	return &logger{log.Default()}
}

const (
	traceLevel = iota
	debugLevel
	infoLevel
	warnLevel
	errorLevel
	fatalLevel

	lnFormat = "%v\n"
)

func addPrefix(format string, level int) string {
	switch level {
	case traceLevel:
		return "[TRACE] " + format
	case debugLevel:
		return "[DEBUG] " + format
	case infoLevel:
		return "[INFO] " + format
	case warnLevel:
		return "[WARN] " + format
	case errorLevel:
		return "[ERROR] " + format
	case fatalLevel:
		return "[FATAL] " + format
	default:
		return "" + format
	}
}

func (l *logger) Trace(v ...any) {
	l.Println(addPrefix(lnFormat, traceLevel), v)
}

func (l *logger) Tracef(format string, v ...any) {
	l.Printf(addPrefix(format, traceLevel), v...)
}

func (l *logger) Debug(v ...any) {
	l.Printf(addPrefix(lnFormat, debugLevel), v)
}

func (l *logger) Debugf(format string, v ...any) {
	l.Printf(addPrefix(format, debugLevel), v...)
}

func (l *logger) Info(v ...any) {
	l.Printf(addPrefix(lnFormat, infoLevel), v)
}

func (l *logger) Infof(format string, v ...any) {
	l.Printf(addPrefix(format, infoLevel), v...)
}

func (l *logger) Warn(v ...any) {
	l.Printf(addPrefix(lnFormat, warnLevel), v)
}

func (l *logger) Warnf(format string, v ...any) {
	l.Printf(addPrefix(format, warnLevel), v...)
}

func (l *logger) Error(v ...any) {
	l.Printf(addPrefix(lnFormat, errorLevel), v)
}

func (l *logger) Errorf(format string, v ...any) {
	l.Printf(addPrefix(format, errorLevel), v...)
}

func (l *logger) Fatal(v ...any) {
	l.Fatalf(addPrefix(lnFormat, fatalLevel), v)
}

func (l *logger) Fatalf(format string, v ...any) {
	l.Fatalf(addPrefix(format, fatalLevel), v...)
}
