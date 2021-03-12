/*
Package logger makes it easy to add a multilevel logging interface to objects, so that log entries associated with an object are annotated with the object's identity.  New loggers can inherit from existing loggers and add more specific
annotations.
*/
package logger

type nilRawLogger struct{}

func (l *nilRawLogger) Output(calldepth int, s string) error {
	return nil
}

func (l *nilRawLogger) GetLogLevel() LogLevel {
	return LogLevelFatal
}

// NilRawLogger is a precreated RawLogger that just throws away its output
var NilRawLogger = &nilRawLogger{}

// NilRawLogger is a precreated Logger that just throws away its output
var NilLogger = NewLogWrapper(NilRawLogger, "", LogLevelFatal)
