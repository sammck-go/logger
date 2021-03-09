/*
Package logger makes it easy to add a multilevel logging interface to objects, so that log entries associated with an object are annotated with the object's identity.  New loggers can inherit from existing loggers and add more specific
annotations.
*/
package logger

import (
	"log"
	"os"
)

// RawLogger is a minimal logging interface for an underlying logging component. A full-featured Logger implementation
// can be wrapped around anything that implements this interface.  Note that this is a subset of
// log.Logger, so any standard golang logger can be used.
type RawLogger interface {
	// Output writes the output for a logging event. The string s contains
	// the text to print after the prefix specified by the flags of the
	// Logger. A newline is appended if the last character of s is not
	// already a newline. Calldepth is used to recover the PC, filename,
	// etc.  If set to 1, it will display context for the immediate caller
	// of Output; if set to 2, the caller's caller, etc.
	Output(calldepth int, s string) error
}

// GetLogLeveler is an interface for a logger that supports GetLogLevel()
type GetLogLeveler interface {
	GetLogLevel() LogLevel
}

// Logger is an interface for a logging component that supports logging levels and prefix forking
type Logger interface {
	RawLogger
	GetLogLeveler

	// CdRawOutput is the lowest level log output method; it writes the output for a logging event
	// without any additional prefixes not supplied by the raw logger. The string s contains
	// the text to print after the prefix specified by the flags of the raw Logger. A newline
	// is appended if the last character of s is not already a newline. Calldepth is used to
	// recover the call frame, PC, filename, etc.  If set to 1, it will display context for the
	// immediate caller of CdRawOutput; if set to 2, the caller's caller, etc.
	CdRawOutput(calldepth int, s string)

	// CdPrint writes arguments to a Logger with a provided call depth. Arguments are formatted in the style of fmt.Sprint()
	CdPrint(calldepth int, args ...interface{})

	// Print outputs to a Logger in the style of fmt.Sprint()
	Print(args ...interface{})

	// CdPrintf outputs formatted text to a Logger with a provided call depth, in the style of fmt.Sprintf
	CdPrintf(calldepth int, f string, args ...interface{})

	// Printf outputs to a Logger in the style of fmt.Sprintf
	Printf(f string, args ...interface{})

	// CdLogStrNoPrefix outputs a single string to a Logger with provided call depth and without the prefix (beyond the raw
	// logger's prefix) if the given logLevel is enabled. Then,
	// if the given logLevel is LogLevelPanic or LogLevelFatal, exits appropriately.
	CdLogStrNoPrefix(calldepth int, logLevel LogLevel, s string)

	// LogStrNoPrefix outputs a single string to a Logger without the prefix (beyond the raw
	// logger's prefix) if the given logLevel is enabled. Then,
	// if the given logLevel is LogLevelPanic or LogLevelFatal, exits appropriately.
	LogStrNoPrefix(logLevel LogLevel, s string)

	// CdLogNoPrefix outputs to a Logger with provided call depth and without the prefix (beyond the raw
	// logger's prefix) if the given logLevel is enabled. Then,
	// if the given logLevel is LogLevelPanic or LogLevelFatal, exits appropriately.
	// Arguments are formatted in the style of fmt.Sprint
	CdLogNoPrefix(calldepth int, logLevel LogLevel, args ...interface{})

	// LogNoPrefix outputs to a Logger without the prefix if the given logLevel is enabled. Then,
	// if the given logLevel is LogLevelPanic or LogLevelFatal, exits appropriately
	// Arguments are formatted in the style of fmt.Sprint
	LogNoPrefix(logLevel LogLevel, args ...interface{})

	// CdLogfNoPrefix outputs to a Logger with a given call depth and without the prefix if the given logLevel is enabled. Then,
	// if the given logLevel is LogLevelPanic or LogLevelFatal, exits appropriately.
	// Arguments are formatted in the style of fmt.Sprintf
	CdLogfNoPrefix(calldepth int, logLevel LogLevel, f string, args ...interface{})

	// LogfNoPrefix outputs to a Logger without the prefix if the given logLevel is enabled. Then,
	// if the given logLevel is LogLevelPanic or LogLevelFatal, exits appropriately
	// Arguments are formatted in the style of fmt.Sprintf
	LogfNoPrefix(logLevel LogLevel, f string, args ...interface{})

	// CdLog outputs to a Logger with a given calldepth if the given logLevel is enabled. Then,
	// if the given logLevel is LogLevelPanic or LogLevelFatal, exits appropriately
	// Arguments are formatted in the style of fmt.Sprint
	CdLog(calldepth int, logLevel LogLevel, args ...interface{})

	// Log outputs to a Logger if the given logLevel is enabled. Then,
	// if the given logLevel is LogLevelPanic or LogLevelFatal, exits appropriately
	// Arguments are formatted in the style of fmt.Sprint
	Log(logLevel LogLevel, args ...interface{})

	// CdLogf outputs to a Logger with a given calldepth if the given logLevel is enabled. Then,
	// if the given logLevel is LogLevelPanic or LogLevelFatal, exits appropriately
	// Arguments are formatted in the style of fmt.Sprintf
	CdLogf(calldepth int, logLevel LogLevel, f string, args ...interface{})

	// Logf outputs to a Logger if the given logLevel is enabled. Then,
	// if the given logLevel is LogLevelPanic or LogLevelFatal, exits appropriately
	// Arguments are formatted in the style of fmt.Sprintf
	Logf(logLevel LogLevel, f string, args ...interface{})

	// CdLogErrorf outputs an error message with a given calldepth to a Logger iff logLevel is enabled,
	// then returns an error object with a description string that has the
	// logger's prefix. If the given logLevel is LogLevelPanic or LogLevelFatal, does not return.
	// Arguments are formatted in the style of fmt.Sprintf.
	CdLogErrorf(calldepth int, logLevel LogLevel, f string, args ...interface{}) error

	// LogErrorf outputs an error message to a Logger iff logLevel is enabled,
	// then returns an error object with a description string that has the
	// logger's prefix. If the given logLevel is LogLevelPanic or LogLevelFatal, does not return.
	// Arguments are formatted in the style of fmt.Sprintf.
	LogErrorf(logLevel LogLevel, f string, args ...interface{}) error

	// CdLogError outputs an error message with a given calldepth to a Logger iff logLevel is enabled,
	// then returns an error object with a description string that has the
	// logger's prefix. If the given logLevel is LogLevelPanic or LogLevelFatal, does not return.
	// Arguments are formatted in the style of fmt.Sprint.
	CdLogError(calldepth int, logLevel LogLevel, args ...interface{}) error

	// LogError outputs an error message to a Logger iff logLevel is enabled,
	// then returns an error object with a description string that has the
	// logger's prefix. If the given logLevel is LogLevelPanic or LogLevelFatal, does not return.
	// Arguments are formatted in the style of fmt.Sprint.
	LogError(logLevel LogLevel, args ...interface{}) error

	// CdPanic outputs a log message with a given call depth if LogLevelPanic is enabled, and then panics.
	// Arguments are formatted in the style of fmt.Sprint.
	CdPanic(calldepth int, args ...interface{})

	// Panic outputs a log message if LogLevelPanic is enabled, and then panics.
	// Arguments are formatted in the style of fmt.Sprint.
	Panic(args ...interface{})

	// CdPanicOnError does nothing if err is nil; otherwise
	// outputs a log message with a given call depth if LogLevelPanic is enabled, and then panics.
	CdPanicOnError(calldepth int, err error)

	// PanicOnError does nothing if err is nil; otherwise
	// outputs a log message if LogLevelPanic is enabled, and then panics.
	PanicOnError(err error)

	// Panicf outputs a formatted log message if LogLevelPanic is enabled, and then panics.
	// Arguments are formatted in the style of fmt.Sprintf.
	Panicf(f string, args ...interface{})

	// Fatal outputs a log message if LogLevelFatal is enabled, and then exits with error code 1.
	// Arguments are formatted in the style of fmt.Sprint.
	Fatal(args ...interface{})

	// Fatalf outputs a formatted log message if LogLevelFatal is enabled, and then exits with error code.
	// Arguments are formatted in the style of fmt.Sprintf.
	Fatalf(f string, args ...interface{})

	// ELog outputs a formatted log message if LogLevelError is enabled.
	// Arguments are formatted in the style of fmt.Sprint.
	ELog(args ...interface{})

	// ELogf outputs a formatted log message if LogLevelError is enabled.
	// Arguments are formatted in the style of fmt.Sprintf.
	ELogf(f string, args ...interface{})

	// WLog outputs a formatted log message if LogLevelWarning is enabled.
	// Arguments are formatted in the style of fmt.Sprint.
	WLog(args ...interface{})

	// WLogf outputs a formatted log message if LogLevelWarning is enabled.
	// Arguments are formatted in the style of fmt.Sprintf.
	WLogf(f string, args ...interface{})

	// ILog outputs a formatted log message if LogLevelInfo is enabled.
	// Arguments are formatted in the style of fmt.Sprint.
	ILog(args ...interface{})

	// ILogf outputs a formatted log message if LogLevelInfo is enabled.
	// Arguments are formatted in the style of fmt.Sprintf.
	ILogf(f string, args ...interface{})

	// DLog outputs a formatted log message if LogLevelDebug is enabled.
	// Arguments are formatted in the style of fmt.Sprint.
	DLog(args ...interface{})

	// DLogf outputs a formatted log message if LogLevelDebug is enabled.
	// Arguments are formatted in the style of fmt.Sprintf.
	DLogf(f string, args ...interface{})

	// TLog outputs a formatted log message if LogLevelTrace is enabled.
	// Arguments are formatted in the style of fmt.Sprint.
	TLog(args ...interface{})

	// TLogf outputs a formatted log message if LogLevelTrace is enabled.
	// Arguments are formatted in the style of fmt.Sprintf.
	TLogf(f string, args ...interface{})

	// CdError generates an error object with a given calldepth and this logger's prefix.
	// Arguments are formatted in the style of fmt.Sprint.
	// Note: The raw logger's prefix, if any, is not included.
	// Note: calldepth is not currently used but is included to allow future improvements
	//   that can put caller context into the error message.
	CdError(calldepth int, args ...interface{}) error

	// Error generates an error object with this logger's prefix.
	// Arguments are formatted in the style of fmt.Sprint.
	// Note: The raw logger's prefix, if any, is not included.
	Error(args ...interface{}) error

	// CdErrorf generates an error object with a given calldepth and this logger's prefix.
	// Arguments are formatted in the style of fmt.Sprintf.
	// Note: The raw logger's prefix, if any, is not included.
	// Note: calldepth is not currently used but is included to allow future improvements
	//   that can put caller context into the error message.
	CdErrorf(calldepth int, f string, args ...interface{}) error

	// Errorf generates an error object with this logger's prefix.
	// Arguments are formatted in the style of fmt.Sprintf.
	// Note: The raw logger's prefix, if any, is not included.
	Errorf(f string, args ...interface{}) error

	// CdSprintf returns a string that has the Logger's prefix, with a given call depth.
	// Arguments are formatted in the style of fmt.Sprintf.
	// Note: The raw logger's prefix, if any, is not included.
	// Note: calldepth is not currently used but is included to allow future improvements
	//   that can put caller context into the prefix.
	CdSprintf(calldepth int, f string, args ...interface{}) string

	// Sprintf returns a string that has the Logger's prefix
	// Arguments are formatted in the style of fmt.Sprintf.
	// Note: The raw logger's prefix, if any, is not included.
	Sprintf(f string, args ...interface{}) string

	// CdSprint returns a string that has the Logger's prefix, with a given call depth.
	// Arguments are formatted in the style of fmt.Sprint.
	// Note: The raw logger's prefix, if any, is not included.
	// Note: calldepth is not currently used but is included to allow future improvements
	//   that can put caller context into the prefix.
	CdSprint(calldepth int, args ...interface{}) string

	// Sprint returns a string that has the Logger's prefix
	// Arguments are formatted in the style of fmt.Sprint.
	// Note: The raw logger's prefix, if any, is not included.
	Sprint(args ...interface{}) string

	// ELogError outputs an error message to a Logger iff LogLevelError is enabled,
	// and returns an error object with a description string that has the
	// logger's prefix.
	// Arguments are formatted in the style of fmt.Sprint.
	ELogError(args ...interface{}) error

	// ELogErrorf outputs an error message to a Logger iff LogLevelError is enabled,
	// and returns an error object with a description string that has the
	// logger's prefix.
	// Arguments are formatted in the style of fmt.Sprintf.
	ELogErrorf(f string, args ...interface{}) error

	// WLogError outputs an error message to a Logger iff LogLevelWarning is enabled,
	// and returns an error object with a description string that has the
	// logger's prefix.
	// Arguments are formatted in the style of fmt.Sprint.
	WLogError(args ...interface{}) error

	// WLogErrorf outputs an error message to a Logger iff LogLevelWarning is enabled,
	// and returns an error object with a description string that has the
	// logger's prefix.
	// Arguments are formatted in the style of fmt.Sprintf.
	WLogErrorf(f string, args ...interface{}) error

	// ILogError outputs an error message to a Logger iff LogLevelInfo is enabled,
	// and returns an error object with a description string that has the
	// logger's prefix.
	// Arguments are formatted in the style of fmt.Sprint.
	ILogError(args ...interface{}) error

	// ILogErrorf outputs an error message to a Logger iff LogLevelInfo is enabled,
	// and returns an error object with a description string that has the
	// logger's prefix.
	// Arguments are formatted in the style of fmt.Sprintf.
	ILogErrorf(f string, args ...interface{}) error

	// DLogError outputs an error message to a Logger iff LogLevelDebug is enabled,
	// and returns an error object with a description string that has the
	// logger's prefix.
	// Arguments are formatted in the style of fmt.Sprint.
	DLogError(args ...interface{}) error

	// DLogErrorf outputs an error message to a Logger iff LogLevelDebug is enabled,
	// and returns an error object with a description string that has the
	// logger's prefix.
	// Arguments are formatted in the style of fmt.Sprintf.
	DLogErrorf(f string, args ...interface{}) error

	// TLogError outputs an error message to a Logger iff LogLevelTrace is enabled,
	// and returns an error object with a description string that has the
	// logger's prefix.
	// Arguments are formatted in the style of fmt.Sprint.
	TLogError(args ...interface{}) error

	// TLogErrorf outputs an error message to a Logger iff LogLevelTrace is enabled,
	// and returns an error object with a description string that has the
	// logger's prefix.
	// Arguments are formatted in the style of fmt.Sprintf.
	TLogErrorf(f string, args ...interface{}) error

	// ForkLogStr creates a new Logger that has an additional string appended onto
	// an existing logger's prefix (with ": " added between).
	ForkLogStr(prefix string) Logger

	// ForkLogf creates a new Logger that has an additional formatted string appended onto
	// an existing logger's prefix (with ": " added between).
	// Arguments are formatted in the style of fmt.Sprintf
	ForkLogf(prefixFmt string, args ...interface{}) Logger

	// ForkLog creates a new Logger that has an additional formatted string appended onto
	// an existing logger's prefix (with ": " added between).
	// Arguments are formatted in the style of fmt.Sprint
	ForkLog(args ...interface{}) Logger

	// Prefix returns the Logger's prefix string (does not include ": " trailer).
	// Does not include the raw logger's prefix, if any.
	Prefix() string

	// SetLogLevel sets the log level
	SetLogLevel(logLevel LogLevel)
}

// NewWithConfig creates a new Logger object from a configuration
func NewWithConfig(cfg *Config) (Logger, error) {
	parentLogger := cfg.parentLogger
	if parentLogger == nil {
		lw := cfg.logWriter
		if lw == nil {
			lw = os.Stderr
		}
		parentLogger = log.New(lw, "", cfg.flag)
	}

	lg := NewLogWrapper(parentLogger, cfg.prefix, cfg.logLevel)

	return lg, nil
}

// New creates a new Logger object from supplied configuration options
func New(opts ...ConfigOption) (Logger, error) {
	cfg := NewConfig(opts...)
	lg, err := NewWithConfig(cfg)
	return lg, err
}

// NewLogWrapper creates a new Logger that wraps an existing base RawLogger with an optional additional
// prefix and an adjusted loglevel. Note that the logging level cannot be effectively increased from the
// logging level of the base logger, since the base logger since the base logger will filter items passed
// to it. As an optimization, the wrapper will limit the loglevel to the base logger's level if
// it implements GetLogLevel(). If the base logger's loglevel subsequently changes, it is the caller's
// responsibility to adjust the new wrapper's loglevel if desired.
func NewLogWrapper(logger RawLogger, prefix string, logLevel LogLevel) Logger {
	if logLevel > LogLevelFatal {
		gll, ok := logger.(GetLogLeveler)
		if ok {
			gllLevel := gll.GetLogLevel()
			if gllLevel < logLevel {
				logLevel = gllLevel
			}
		}
	}

	prefixC := prefix
	if prefixC != "" {
		prefixC += ": "
	}

	l := &BasicLogger{
		prefix:   prefix,
		prefixC:  prefixC,
		logger:   logger,
		logLevel: logLevel,
	}
	return l
}
