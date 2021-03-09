package logger

import (
	"errors"
	"fmt"
	"os"
)

// BasicLogger is a logical log output stream with a level filter
// and a prefix added to each output record.
type BasicLogger struct {
	// prefix includes all of the inherited prefixes, delimited with ": ". It
	// is an empty string if this is the top level.
	prefix string
	// prefixC is prefix if prefix is empty; otherwise prefix + ": "
	prefixC string
	// logger is the raw logger
	logger   RawLogger
	logLevel LogLevel
}

// CdRawOutput is the lowest level log output method; it writes the output for a logging event
// without any additional prefixes not supplied by the raw logger. The string s contains
// the text to print after the prefix specified by the flags of the raw Logger. A newline
// is appended if the last character of s is not already a newline. Calldepth is used to
// recover the call frame, PC, filename, etc.  If set to 1, it will display context for the
// immediate caller of CdRawOutput; if set to 2, the caller's caller, etc.
func (l *BasicLogger) CdRawOutput(calldepth int, s string) {
	l.logger.Output(calldepth+1, s)
}

// Output is the compatible with log.Logger.Output, to make this a RawLogger
func (l *BasicLogger) Output(calldepth int, s string) error {
	return l.logger.Output(calldepth+1, l.Sprint(s))
}

// CdPrint writes arguments to a Logger with a provided call depth. Arguments are formatted in the style of fmt.Sprint()
func (l *BasicLogger) CdPrint(calldepth int, args ...interface{}) {
	l.CdRawOutput(calldepth+1, l.Sprint(args...))
}

// Print outputs to a Logger in the style of fmt.Sprint()
func (l *BasicLogger) Print(args ...interface{}) {
	l.CdPrint(2, args...)
}

// CdPrintf outputs formatted text to a Logger with a provided call depth, in the style of fmt.Sprintf
func (l *BasicLogger) CdPrintf(calldepth int, f string, args ...interface{}) {
	l.CdRawOutput(calldepth+1, l.Sprintf(f, args...))
}

// Printf outputs to a Logger in the style of fmt.Sprintf
func (l *BasicLogger) Printf(f string, args ...interface{}) {
	l.CdPrintf(2, f, args...)
}

// CdLogStrNoPrefix outputs a single string to a Logger with provided call depth and without the prefix (beyond the raw
// logger's prefix) if the given logLevel is enabled. Then,
// if the given logLevel is LogLevelPanic or LogLevelFatal, exits appropriately.
func (l *BasicLogger) CdLogStrNoPrefix(calldepth int, logLevel LogLevel, s string) {
	if logLevel <= l.logLevel || logLevel <= LogLevelFatal {
		if logLevel >= LogLevelPanic {
			l.CdRawOutput(calldepth+1, s)
		}
		if logLevel == LogLevelFatal {
			os.Exit(1)
		}
		if logLevel == LogLevelPanic {
			panic(s)
		}
	}
}

// LogStrNoPrefix outputs a single string to a Logger without the prefix (beyond the raw
// logger's prefix) if the given logLevel is enabled. Then,
// if the given logLevel is LogLevelPanic or LogLevelFatal, exits appropriately.
func (l *BasicLogger) LogStrNoPrefix(logLevel LogLevel, s string) {
	l.CdLogStrNoPrefix(2, logLevel, s)
}

// CdLogNoPrefix outputs to a Logger with provided call depth and without the prefix (beyond the raw
// logger's prefix) if the given logLevel is enabled. Then,
// if the given logLevel is LogLevelPanic or LogLevelFatal, exits appropriately.
// Arguments are formatted in the style of fmt.Sprint
func (l *BasicLogger) CdLogNoPrefix(calldepth int, logLevel LogLevel, args ...interface{}) {
	if logLevel <= l.logLevel || logLevel <= LogLevelFatal {
		msg := fmt.Sprint(args...)
		if logLevel >= LogLevelPanic {
			l.CdRawOutput(calldepth+1, msg)
		}
		if logLevel == LogLevelFatal {
			os.Exit(1)
		}
		if logLevel == LogLevelPanic {
			panic(msg)
		}
	}
}

// LogNoPrefix outputs to a Logger without the prefix if the given logLevel is enabled. Then,
// if the given logLevel is LogLevelPanic or LogLevelFatal, exits appropriately
// Arguments are formatted in the style of fmt.Sprint
func (l *BasicLogger) LogNoPrefix(logLevel LogLevel, args ...interface{}) {
	l.CdLogNoPrefix(2, logLevel, args...)
}

// CdLogfNoPrefix outputs to a Logger with a given call depth and without the prefix if the given logLevel is enabled. Then,
// if the given logLevel is LogLevelPanic or LogLevelFatal, exits appropriately.
// Arguments are formatted in the style of fmt.Sprintf
func (l *BasicLogger) CdLogfNoPrefix(calldepth int, logLevel LogLevel, f string, args ...interface{}) {
	if logLevel <= l.logLevel || logLevel <= LogLevelFatal {
		msg := fmt.Sprintf(f, args...)
		if logLevel <= LogLevelPanic {
			l.CdRawOutput(calldepth+1, msg)
		}
		if logLevel == LogLevelFatal {
			os.Exit(1)
		}
		if logLevel == LogLevelPanic {
			panic(msg)
		}
	}
}

// LogfNoPrefix outputs to a Logger without the prefix if the given logLevel is enabled. Then,
// if the given logLevel is LogLevelPanic or LogLevelFatal, exits appropriately
// Arguments are formatted in the style of fmt.Sprintf
func (l *BasicLogger) LogfNoPrefix(logLevel LogLevel, f string, args ...interface{}) {
	l.CdLogfNoPrefix(2, logLevel, f, args...)
}

// CdLog outputs to a Logger with a given calldepth if the given logLevel is enabled. Then,
// if the given logLevel is LogLevelPanic or LogLevelFatal, exits appropriately
// Arguments are formatted in the style of fmt.Sprint
func (l *BasicLogger) CdLog(calldepth int, logLevel LogLevel, args ...interface{}) {
	if logLevel <= l.logLevel || logLevel <= LogLevelFatal {
		msg := l.Sprint(args...)
		l.CdLogStrNoPrefix(calldepth+1, logLevel, msg)
	}
}

// Log outputs to a Logger if the given logLevel is enabled. Then,
// if the given logLevel is LogLevelPanic or LogLevelFatal, exits appropriately
// Arguments are formatted in the style of fmt.Sprint
func (l *BasicLogger) Log(logLevel LogLevel, args ...interface{}) {
	l.CdLog(2, logLevel, args...)
}

// CdLogf outputs to a Logger with a given calldepth if the given logLevel is enabled. Then,
// if the given logLevel is LogLevelPanic or LogLevelFatal, exits appropriately
// Arguments are formatted in the style of fmt.Sprintf
func (l *BasicLogger) CdLogf(calldepth int, logLevel LogLevel, f string, args ...interface{}) {
	if logLevel <= l.logLevel || logLevel <= LogLevelFatal {
		msg := l.Sprintf(f, args...)
		l.CdLogStrNoPrefix(calldepth+1, logLevel, msg)
	}
}

// Logf outputs to a Logger if the given logLevel is enabled. Then,
// if the given logLevel is LogLevelPanic or LogLevelFatal, exits appropriately
// Arguments are formatted in the style of fmt.Sprintf
func (l *BasicLogger) Logf(logLevel LogLevel, f string, args ...interface{}) {
	l.CdLogf(2, logLevel, f, args...)
}

// CdLogErrorf outputs an error message with a given calldepth to a Logger iff logLevel is enabled,
// then returns an error object with a description string that has the
// logger's prefix. If the given logLevel is LogLevelPanic or LogLevelFatal, does not return.
// Arguments are formatted in the style of fmt.Sprintf.
func (l *BasicLogger) CdLogErrorf(calldepth int, logLevel LogLevel, f string, args ...interface{}) error {
	msg := l.Sprintf(f, args...)
	l.CdLogStrNoPrefix(calldepth+1, logLevel, msg)
	return errors.New(msg)
}

// LogErrorf outputs an error message to a Logger iff logLevel is enabled,
// then returns an error object with a description string that has the
// logger's prefix. If the given logLevel is LogLevelPanic or LogLevelFatal, does not return.
// Arguments are formatted in the style of fmt.Sprintf.
func (l *BasicLogger) LogErrorf(logLevel LogLevel, f string, args ...interface{}) error {
	return l.CdLogErrorf(2, logLevel, f, args...)
}

// CdLogError outputs an error message with a given calldepth to a Logger iff logLevel is enabled,
// then returns an error object with a description string that has the
// logger's prefix. If the given logLevel is LogLevelPanic or LogLevelFatal, does not return.
// Arguments are formatted in the style of fmt.Sprint.
func (l *BasicLogger) CdLogError(calldepth int, logLevel LogLevel, args ...interface{}) error {
	msg := l.Sprint(args...)
	l.CdLogStrNoPrefix(calldepth+1, logLevel, msg)
	return errors.New(msg)
}

// LogError outputs an error message to a Logger iff logLevel is enabled,
// then returns an error object with a description string that has the
// logger's prefix. If the given logLevel is LogLevelPanic or LogLevelFatal, does not return.
// Arguments are formatted in the style of fmt.Sprint.
func (l *BasicLogger) LogError(logLevel LogLevel, args ...interface{}) error {
	return l.CdLogError(2, logLevel, args...)
}

// CdPanic outputs a log message with a given call depth if LogLevelPanic is enabled, and then panics.
// Arguments are formatted in the style of fmt.Sprint.
func (l *BasicLogger) CdPanic(calldepth int, args ...interface{}) {
	l.CdLog(calldepth+1, LogLevelPanic, args...)
}

// Panic outputs a log message if LogLevelPanic is enabled, and then panics.
// Arguments are formatted in the style of fmt.Sprint.
func (l *BasicLogger) Panic(args ...interface{}) {
	l.CdPanic(2, args...)
}

// CdPanicOnError does nothing if err is nil; otherwise
// outputs a log message with a given call depth if LogLevelPanic is enabled, and then panics.
func (l *BasicLogger) CdPanicOnError(calldepth int, err error) {
	if err != nil {
		l.CdPanic(calldepth+1, err)
	}
}

// PanicOnError does nothing if err is nil; otherwise
// outputs a log message if LogLevelPanic is enabled, and then panics.
func (l *BasicLogger) PanicOnError(err error) {
	if err != nil {
		l.CdPanic(2, err)
	}
}

// Panicf outputs a formatted log message if LogLevelPanic is enabled, and then panics.
// Arguments are formatted in the style of fmt.Sprintf.
func (l *BasicLogger) Panicf(f string, args ...interface{}) {
	l.CdLogf(2, LogLevelPanic, f, args...)
}

// Fatal outputs a log message if LogLevelFatal is enabled, and then exits with error code 1.
// Arguments are formatted in the style of fmt.Sprint.
func (l *BasicLogger) Fatal(args ...interface{}) {
	l.CdLog(2, LogLevelFatal, args...)
}

// Fatalf outputs a formatted log message if LogLevelFatal is enabled, and then exits with error code.
// Arguments are formatted in the style of fmt.Sprintf.
func (l *BasicLogger) Fatalf(f string, args ...interface{}) {
	l.CdLogf(2, LogLevelFatal, f, args...)
}

// ELog outputs a formatted log message if LogLevelError is enabled.
// Arguments are formatted in the style of fmt.Sprint.
func (l *BasicLogger) ELog(args ...interface{}) {
	l.CdLog(2, LogLevelError, args...)
}

// ELogf outputs a formatted log message if LogLevelError is enabled.
// Arguments are formatted in the style of fmt.Sprintf.
func (l *BasicLogger) ELogf(f string, args ...interface{}) {
	l.CdLogf(2, LogLevelError, f, args...)
}

// WLog outputs a formatted log message if LogLevelWarning is enabled.
// Arguments are formatted in the style of fmt.Sprint.
func (l *BasicLogger) WLog(args ...interface{}) {
	l.CdLog(2, LogLevelWarning, args...)
}

// WLogf outputs a formatted log message if LogLevelWarning is enabled.
// Arguments are formatted in the style of fmt.Sprintf.
func (l *BasicLogger) WLogf(f string, args ...interface{}) {
	l.CdLogf(2, LogLevelWarning, f, args...)
}

// ILog outputs a formatted log message if LogLevelInfo is enabled.
// Arguments are formatted in the style of fmt.Sprint.
func (l *BasicLogger) ILog(args ...interface{}) {
	l.CdLog(2, LogLevelInfo, args...)
}

// ILogf outputs a formatted log message if LogLevelInfo is enabled.
// Arguments are formatted in the style of fmt.Sprintf.
func (l *BasicLogger) ILogf(f string, args ...interface{}) {
	l.CdLogf(2, LogLevelInfo, f, args...)
}

// DLog outputs a formatted log message if LogLevelDebug is enabled.
// Arguments are formatted in the style of fmt.Sprint.
func (l *BasicLogger) DLog(args ...interface{}) {
	l.CdLog(2, LogLevelDebug, args...)
}

// DLogf outputs a formatted log message if LogLevelDebug is enabled.
// Arguments are formatted in the style of fmt.Sprintf.
func (l *BasicLogger) DLogf(f string, args ...interface{}) {
	l.CdLogf(2, LogLevelDebug, f, args...)
}

// TLog outputs a formatted log message if LogLevelTrace is enabled.
// Arguments are formatted in the style of fmt.Sprint.
func (l *BasicLogger) TLog(args ...interface{}) {
	l.CdLog(2, LogLevelTrace, args...)
}

// TLogf outputs a formatted log message if LogLevelTrace is enabled.
// Arguments are formatted in the style of fmt.Sprintf.
func (l *BasicLogger) TLogf(f string, args ...interface{}) {
	l.CdLogf(2, LogLevelTrace, f, args...)
}

// CdError generates an error object with a given calldepth and this logger's prefix.
// Arguments are formatted in the style of fmt.Sprint.
// Note: The raw logger's prefix, if any, is not included.
// Note: calldepth is not currently used but is included to allow future improvements
//   that can put caller context into the error message.
func (l *BasicLogger) CdError(calldepth int, args ...interface{}) error {
	return errors.New(l.Sprint(args...))
}

// Error generates an error object with this logger's prefix.
// Arguments are formatted in the style of fmt.Sprint.
// Note: The raw logger's prefix, if any, is not included.
func (l *BasicLogger) Error(args ...interface{}) error {
	return l.CdError(2, args...)
}

// CdErrorf generates an error object with a given calldepth and this logger's prefix.
// Arguments are formatted in the style of fmt.Sprintf.
// Note: The raw logger's prefix, if any, is not included.
// Note: calldepth is not currently used but is included to allow future improvements
//   that can put caller context into the error message.
func (l *BasicLogger) CdErrorf(calldepth int, f string, args ...interface{}) error {
	return errors.New(l.Sprintf(f, args...))
}

// Errorf generates an error object with this logger's prefix.
// Arguments are formatted in the style of fmt.Sprintf.
// Note: The raw logger's prefix, if any, is not included.
func (l *BasicLogger) Errorf(f string, args ...interface{}) error {
	return l.CdErrorf(2, f, args...)
}

// CdSprintf returns a string that has the Logger's prefix, with a given call depth.
// Arguments are formatted in the style of fmt.Sprintf.
// Note: The raw logger's prefix, if any, is not included.
// Note: calldepth is not currently used but is included to allow future improvements
//   that can put caller context into the prefix.
func (l *BasicLogger) CdSprintf(calldepth int, f string, args ...interface{}) string {
	return l.prefixC + fmt.Sprintf(f, args...)
}

// Sprintf returns a string that has the Logger's prefix
// Arguments are formatted in the style of fmt.Sprintf.
// Note: The raw logger's prefix, if any, is not included.
func (l *BasicLogger) Sprintf(f string, args ...interface{}) string {
	return l.prefixC + fmt.Sprintf(f, args...)
}

// CdSprint returns a string that has the Logger's prefix, with a given call depth.
// Arguments are formatted in the style of fmt.Sprint.
// Note: The raw logger's prefix, if any, is not included.
// Note: calldepth is not currently used but is included to allow future improvements
//   that can put caller context into the prefix.
func (l *BasicLogger) CdSprint(calldepth int, args ...interface{}) string {
	return l.prefixC + fmt.Sprint(args...)
}

// Sprint returns a string that has the Logger's prefix
// Arguments are formatted in the style of fmt.Sprint.
// Note: The raw logger's prefix, if any, is not included.
func (l *BasicLogger) Sprint(args ...interface{}) string {
	return l.prefixC + fmt.Sprint(args...)
}

// ELogError outputs an error message to a Logger iff LogLevelError is enabled,
// and returns an error object with a description string that has the
// logger's prefix.
// Arguments are formatted in the style of fmt.Sprint.
func (l *BasicLogger) ELogError(args ...interface{}) error {
	return l.CdLogError(2, LogLevelError, args...)
}

// ELogErrorf outputs an error message to a Logger iff LogLevelError is enabled,
// and returns an error object with a description string that has the
// logger's prefix.
// Arguments are formatted in the style of fmt.Sprintf.
func (l *BasicLogger) ELogErrorf(f string, args ...interface{}) error {
	return l.CdLogErrorf(2, LogLevelError, f, args...)
}

// WLogError outputs an error message to a Logger iff LogLevelWarning is enabled,
// and returns an error object with a description string that has the
// logger's prefix.
// Arguments are formatted in the style of fmt.Sprint.
func (l *BasicLogger) WLogError(args ...interface{}) error {
	return l.CdLogError(2, LogLevelWarning, args...)
}

// WLogErrorf outputs an error message to a Logger iff LogLevelWarning is enabled,
// and returns an error object with a description string that has the
// logger's prefix.
// Arguments are formatted in the style of fmt.Sprintf.
func (l *BasicLogger) WLogErrorf(f string, args ...interface{}) error {
	return l.CdLogErrorf(2, LogLevelWarning, f, args...)
}

// ILogError outputs an error message to a Logger iff LogLevelInfo is enabled,
// and returns an error object with a description string that has the
// logger's prefix.
// Arguments are formatted in the style of fmt.Sprint.
func (l *BasicLogger) ILogError(args ...interface{}) error {
	return l.CdLogError(2, LogLevelInfo, args...)
}

// ILogErrorf outputs an error message to a Logger iff LogLevelInfo is enabled,
// and returns an error object with a description string that has the
// logger's prefix.
// Arguments are formatted in the style of fmt.Sprintf.
func (l *BasicLogger) ILogErrorf(f string, args ...interface{}) error {
	return l.CdLogErrorf(2, LogLevelInfo, f, args...)
}

// DLogError outputs an error message to a Logger iff LogLevelDebug is enabled,
// and returns an error object with a description string that has the
// logger's prefix.
// Arguments are formatted in the style of fmt.Sprint.
func (l *BasicLogger) DLogError(args ...interface{}) error {
	return l.CdLogError(2, LogLevelDebug, args...)
}

// DLogErrorf outputs an error message to a Logger iff LogLevelDebug is enabled,
// and returns an error object with a description string that has the
// logger's prefix.
// Arguments are formatted in the style of fmt.Sprintf.
func (l *BasicLogger) DLogErrorf(f string, args ...interface{}) error {
	return l.CdLogErrorf(2, LogLevelDebug, f, args...)
}

// TLogError outputs an error message to a Logger iff LogLevelTrace is enabled,
// and returns an error object with a description string that has the
// logger's prefix.
// Arguments are formatted in the style of fmt.Sprint.
func (l *BasicLogger) TLogError(args ...interface{}) error {
	return l.CdLogError(2, LogLevelTrace, args...)
}

// TLogErrorf outputs an error message to a Logger iff LogLevelTrace is enabled,
// and returns an error object with a description string that has the
// logger's prefix.
// Arguments are formatted in the style of fmt.Sprintf.
func (l *BasicLogger) TLogErrorf(f string, args ...interface{}) error {
	return l.CdLogErrorf(2, LogLevelTrace, f, args...)
}

// ForkLogStr creates a new Logger that has an additional string appended onto
// an existing logger's prefix (with ": " added between).
func (l *BasicLogger) ForkLogStr(prefix string) Logger {
	var newPrefix string
	if prefix == "" {
		newPrefix = l.prefix
	} else if l.prefix == "" {
		newPrefix = prefix
	} else {
		newPrefix = l.prefix + ": " + prefix
	}
	ll := NewLogWrapper(l.logger, newPrefix, l.GetLogLevel())
	return ll
}

// ForkLogf creates a new Logger that has an additional formatted string appended onto
// an existing logger's prefix (with ": " added between).
// Arguments are formatted in the style of fmt.Sprintf
func (l *BasicLogger) ForkLogf(prefixFmt string, args ...interface{}) Logger {
	prefix := fmt.Sprintf(prefixFmt, args...)
	return l.ForkLogStr(prefix)
}

// ForkLog creates a new Logger that has an additional formatted string appended onto
// an existing logger's prefix (with ": " added between).
// Arguments are formatted in the style of fmt.Sprint
func (l *BasicLogger) ForkLog(args ...interface{}) Logger {
	prefix := fmt.Sprint(args...)
	return l.ForkLogStr(prefix)
}

// Prefix returns the Logger's prefix string (does not include ": " trailer)
// Does not include the raw logger's prefix, if any.
func (l *BasicLogger) Prefix() string {
	return l.prefix
}

// GetLogLevel returns the log level
func (l *BasicLogger) GetLogLevel() LogLevel {
	return l.logLevel
}

// SetLogLevel sets the log level
func (l *BasicLogger) SetLogLevel(logLevel LogLevel) {
	l.logLevel = logLevel
}
