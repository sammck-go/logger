/*
Package logger makes it easy to add a multilevel logging interface to objects, so that log entries associated with an object are annotated with the object's identity.  New loggers can inherit from existing loggers and add more specific
annotations.
*/
package logger

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// LogLevel specifies the level of spew that shoud go to the log
type LogLevel int

const (
	// LogLevelUnknown is a default value for LogLevel. It's
	// behavior is undefined
	LogLevelUnknown LogLevel = iota

	// LogLevelPanic causes output of an error message followed by a panic
	LogLevelPanic LogLevel = iota

	// LogLevelFatal causes output of an error message followed by os.Exit(1)
	LogLevelFatal LogLevel = iota

	// LogLevelError is for unexpected error messages
	LogLevelError LogLevel = iota

	// LogLevelWarning is for Warning messages
	LogLevelWarning LogLevel = iota

	// LogLevelInfo is for Info messages
	LogLevelInfo LogLevel = iota

	// LogLevelDebug is for debug messaged
	LogLevelDebug LogLevel = iota

	// LogLevelTrace is for trace messages
	LogLevelTrace LogLevel = iota
)

var logLevelNames = [...]string{
	"unknown", "panic", "fatal", "error", "warning", "info", "debug", "trace",
}

var nameToLogError = func() map[string]LogLevel {
	var result = make(map[string]LogLevel)
	for i, name := range logLevelNames {
		result[name] = LogLevel(i)
	}
	return result
}()

// StringToLogLevel converts a string to a LogLevel
func StringToLogLevel(s string) LogLevel {
	result, ok := nameToLogError[strings.ToLower(s)]
	if !ok {
		result = LogLevelUnknown
	}
	return result
}

func (x *LogLevel) String() string {
	y := *x
	if y < LogLevelUnknown || y > LogLevelTrace {
		y = LogLevelUnknown
	}
	return logLevelNames[y]
}

// FromString initiales a LogLevel from a string
func (x *LogLevel) FromString(s string) error {
	result := StringToLogLevel(s)
	var err error
	if result == LogLevelUnknown {
		err = fmt.Errorf("Unknown log level: \"%s\"", s)
	} else {
		*x = result
	}
	return err
}

// MinLogger is a minimal logging interface for a logging component. A full-featured Logger implemntation
// can be wrapped areount anything that implements this interface.  Note that this is a subset of
// log.Logger, so any standard golang logger can be used.
type MinLogger interface {
	Print(args ...interface{})
	Prefix() string
}

// GetLogLeveler is an interface for a logger that supports GetLogLevel()
type GetLogLeveler interface {
	GetLogLevel() LogLevel
}

// Logger is an interface for a logging component that supports logging levels and prefix forking
type Logger interface {
	MinLogger
	GetLogLeveler

	// Panic outputs a log message and then exits with error status
	Panic(args ...interface{})

	// Panicf outputs a log message and then exits with error status
	Panicf(f string, args ...interface{})

	// PanicOnError does nothing if err is nil; otherwise
	// outputs a log message if logLevel permits, and then panics
	PanicOnError(err error)

	// Panicf outputs a log message and then exits with error status
	Fatalf(f string, args ...interface{})

	// Panic outputs a log message and then exits with error status
	Fatal(args ...interface{})

	// Log outputs to a Logger iff logging level is enabled
	Log(logLevel LogLevel, args ...interface{})

	// ELog outputs to a Logger iff logging level is enabled
	Logf(logLevel LogLevel, f string, args ...interface{})

	// ELog outputs to a Logger iff ERROR logging level is enabled
	ELog(args ...interface{})

	// ELogf outputs to a Logger iff ERROR logging level is enabled
	ELogf(f string, args ...interface{})

	// WLog outputs to a Logger iff WARNING logging level is enabled
	WLog(args ...interface{})

	// WLogf outputs to a Logger iff WARNING logging level is enabled
	WLogf(f string, args ...interface{})

	// ILog outputs to a Logger iff INFO logging level is enabled
	ILog(args ...interface{})

	// ILogf outputs to a Logger iff INFO logging level is enabled
	ILogf(f string, args ...interface{})

	// DLog outputs to a Logger iff DEBUG logging level is enabled
	DLog(args ...interface{})

	// DLogf outputs to a Logger iff DEBUG logging level is enabled
	DLogf(f string, args ...interface{})

	// TLog outputs to a Logger iff TRACE logging level is enabled
	TLog(args ...interface{})

	// TLogf outputs to a Logger iff TRACE logging level is enabled
	TLogf(f string, args ...interface{})

	// Error returns an error object with a description string that has the
	// Logger's prefix
	Error(args ...interface{}) error

	// Errorf returns an error object with a description string that has the
	// Logger's prefix
	Errorf(f string, args ...interface{}) error

	// Sprintf returns a string that has the Logger's prefix
	Sprintf(f string, args ...interface{}) string

	// Sprint returns a string that has the Logger's prefix
	Sprint(args ...interface{}) string

	// ELogError outputs an error message to a Logger iff logging level is enabled,
	// and returns an error object with a description string that has the
	// logger's prefix
	ELogError(args ...interface{}) error

	// ELogErrorf outputs an error message to a Logger iff logging level is enabled,
	// and returns an error object with a description string that has the
	// logger's prefix
	ELogErrorf(f string, args ...interface{}) error

	// WLogError outputs an error message to a Logger iff logging level is enabled,
	// and returns an error object with a description string that has the
	// logger's prefix
	WLogError(args ...interface{}) error

	// WLogErrorf outputs an error message to a Logger iff logging level is enabled,
	// and returns an error object with a description string that has the
	// logger's prefix
	WLogErrorf(f string, args ...interface{}) error

	// WLogError outputs an error message to a Logger iff logging level is enabled,
	// and returns an error object with a description string that has the
	// logger's prefix
	ILogError(args ...interface{}) error

	// WLogErrorf outputs an error message to a Logger iff logging level is enabled,
	// and returns an error object with a description string that has the
	// logger's prefix
	ILogErrorf(f string, args ...interface{}) error

	// DLogError outputs an error message to a Logger iff logging level is enabled,
	// and returns an error object with a description string that has the
	// logger's prefix
	DLogError(args ...interface{}) error

	// DLogErrorf outputs an error message to a Logger iff logging level is enabled,
	// and returns an error object with a description string that has the
	// logger's prefix
	DLogErrorf(f string, args ...interface{}) error

	// TLogError outputs an error message to a Logger iff logging level is enabled,
	// and returns an error object with a description string that has the
	// logger's prefix
	TLogError(args ...interface{}) error

	// TLogErrorf outputs an error message to a Logger iff logging level is enabled,
	// and returns an error object with a description string that has the
	// logger's prefix
	TLogErrorf(f string, args ...interface{}) error

	// ForkLog creates a new Logger that has an additional string appended onto
	// an existing logger's prefix (with ": " added between).
	ForkLog(prefix string) Logger

	// ForkLogf creates a new Logger that has an additional formatted string appended onto
	// an existing logger's prefix (with ": " added between).
	ForkLogf(prefixFmt string, args ...interface{}) Logger

	// SetLogLevel updates the logging level for the Logger.
	SetLogLevel(logLevel LogLevel)
}

// Config provides configuration options for contruction of a Logger.  The constructed object is immutable
// after it is constructed by NewConfig.
type Config struct {
	prefix       string
	flag         int
	logLevel     LogLevel
	parentLogger MinLogger
	logWriter    io.Writer
}

// ConfigOption is an opaque configuration option setter created by one of the With functions.
// It follows the Golang "options" pattern.
type ConfigOption func(*Config)

// See log.Logger
const (
	defaultLogFlags = log.Ldate | log.Ltime
	allLogFlags     = log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile | log.Lshortfile | log.LUTC | log.Lmsgprefix
	defaultLogLevel = LogLevelWarning
)

// NewConfig creates a Config object from provided options. The resulting object
// can be passed to New using WithConfig, or directly to NewWithConfig.
func NewConfig(opts ...ConfigOption) *Config {
	cfg := &Config{
		prefix:       "",
		flag:         defaultLogFlags,
		logLevel:     defaultLogLevel,
		parentLogger: nil,
		logWriter:    nil,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return cfg
}

// WithConfig allows initialization of a new configuration object starting with an existing one,
// and incremental initialization of configuration separately from initialization of the PidFile.
// If provided, this option should be appear first in the option list, since it replaces all
// configuration values.
func WithConfig(other *Config) ConfigOption {
	return func(cfg *Config) {
		cfg.prefix = other.prefix
		cfg.flag = other.flag
		cfg.logLevel = other.logLevel
		cfg.parentLogger = other.parentLogger
		cfg.logWriter = other.logWriter
	}
}

// Refine creates a new Config object by applying ConfigOptions to an existing config.
func (cfg *Config) Refine(opts ...ConfigOption) *Config {
	newOpts := append([]ConfigOption{WithConfig(cfg)}, opts...)
	newConfig := NewConfig(newOpts...)
	return newConfig
}

// WithPrefix sets the prefix for the new logger. By default, no prefix is added.
func WithPrefix(prefix string) ConfigOption {
	return func(cfg *Config) {
		cfg.prefix = prefix
	}
}

// WithReplaceLogFlags replaces all log flags for the new logger. By default, log.Ldate and log.Ltime
// are set.  Note that flags are ignored if WithLogger() is provided. See log.Logger for a description of flag bits.
func WithReplaceLogFlags(flag int) ConfigOption {
	return func(cfg *Config) {
		cfg.flag = flag
	}
}

// WithLogFlags adds log flags for the new logger, without removing any existing flags. By default, log.Ldate and log.Ltime
// are set. Note that flags are ignored if WithLogger() is provided. See log.Logger for a description of flag bits.
func WithLogFlags(flag int) ConfigOption {
	return func(cfg *Config) {
		cfg.flag |= flag
	}
}

// WithoutLogFlags clears specified log flags for the new logger, without affecting any other flags. By default, log.Ldate and log.Ltime
// are set. Note that flags are ignored if WithLogger() is provided. See log.Logger for a description of flag bits.
func WithoutLogFlags(flag int) ConfigOption {
	return func(cfg *Config) {
		cfg.flag &= ^flag
	}
}

// WithLogLevel sets the prefix for the new logger. By default, no prefix is added.
func WithLogLevel(logLevel LogLevel) ConfigOption {
	return func(cfg *Config) {
		cfg.logLevel = logLevel
	}
}

// WithWriter sets the io,Writer to which log output will be sent. By default, log output will be sent to stderr.
// This setting replaces any prior effect of WithLogger().
func WithWriter(logWriter io.Writer) ConfigOption {
	return func(cfg *Config) {
		cfg.logWriter = logWriter
		cfg.parentLogger = nil
	}
}

// WithLogger sets the parent logger for new logger. By default, a new logger to stderr
// is created. This setting replaces any prior effect of WithWriter().
func WithLogger(parentLogger MinLogger) ConfigOption {
	return func(cfg *Config) {
		cfg.parentLogger = parentLogger
		cfg.logWriter = nil
	}
}

// BasicLogger is a logical log output stream with a level filter
// and a prefix added to each output record.
type BasicLogger struct {
	// prefix includes all of the inherited prefixes, delimited with ": ". It
	// is an empty string if this is the top level.
	prefix string
	// prefixC is prefix if prefix is empty; otherwise prefix + ": "
	prefixC string
	// logger is the raw logger
	logger   MinLogger
	logLevel LogLevel
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

// NewLogWrapper creates a new Logger that wraps an existing base MinLogger with an optional additional
// prefix and an adjusted loglevel. Note that the logging level cannot be effectively increased from the
// logging level of the base logger, since the base logger since the base logger will filter items passed
// to it. As an optimization, the wrapper will limit the loglevel to the base logger's level if
// it implements GetLogLevel(). If the base logger's loglevel subsequently changes, it is the caller's
// responsibility to adjust the new wrapper's loglevel if desired.
func NewLogWrapper(logger MinLogger, prefix string, logLevel LogLevel) Logger {
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

// Print outputs to a Logger
func (l *BasicLogger) Print(args ...interface{}) {
	l.logger.Print(l.Sprint(args...))
}

// Printf outputs to a Logger
func (l *BasicLogger) Printf(f string, args ...interface{}) {
	l.logger.Print(l.Sprintf(f, args...))
}

// LogNoPrefix outputs to a Logger without the prefix if the given logLevel is enabled. Then,
// if the given logLevel is LogLevelPanic or LogLevelFatal, exits appropriately
func (l *BasicLogger) LogNoPrefix(logLevel LogLevel, args ...interface{}) {
	if logLevel <= l.logLevel || logLevel <= LogLevelFatal {
		msg := fmt.Sprint(args...)
		if logLevel >= LogLevelPanic {
			l.logger.Print(msg)
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
func (l *BasicLogger) LogfNoPrefix(logLevel LogLevel, f string, args ...interface{}) {
	if logLevel <= l.logLevel || logLevel <= LogLevelFatal {
		msg := fmt.Sprintf(f, args...)
		if logLevel <= LogLevelPanic {
			l.logger.Print(msg)
		}
		if logLevel == LogLevelFatal {
			os.Exit(1)
		}
		if logLevel == LogLevelPanic {
			panic(msg)
		}
	}
}

// Log outputs to a Logger if the given logLevel is enabled. Then,
// if the given logLevel is LogLevelPanic or LogLevelFatal, exits appropriately
func (l *BasicLogger) Log(logLevel LogLevel, args ...interface{}) {
	if logLevel <= l.logLevel || logLevel <= LogLevelFatal {
		msg := l.Sprint(args...)
		l.LogNoPrefix(logLevel, msg)
	}
}

// Logf outputs to a Logger if the given logLevel is enabled. Then,
// if the given logLevel is LogLevelPanic or LogLevelFatal, exits appropriately
func (l *BasicLogger) Logf(logLevel LogLevel, f string, args ...interface{}) {
	if logLevel <= l.logLevel || logLevel <= LogLevelFatal {
		msg := l.Sprintf(f, args...)
		l.LogNoPrefix(logLevel, msg)
	}
}

// LogErrorf outputs an error message to a Logger iff logging level is enabled,
// and returns an error object with a description string that has the
// logger's prefix
func (l *BasicLogger) LogErrorf(logLevel LogLevel, f string, args ...interface{}) error {
	msg := l.Sprintf(f, args...)
	l.LogNoPrefix(logLevel, msg)
	return errors.New(msg)
}

// LogError outputs an error message to a Logger iff logging level is enabled,
// and returns an error object with a description string that has the
// logger's prefix
func (l *BasicLogger) LogError(logLevel LogLevel, args ...interface{}) error {
	msg := l.Sprint(args...)
	l.LogNoPrefix(logLevel, msg)
	return errors.New(msg)
}

// Panic outputs a log message if logLevel permits, and then panics
func (l *BasicLogger) Panic(args ...interface{}) {
	l.Log(LogLevelPanic, args...)
}

// PanicOnError does nothing if err is nil; otherwise
// outputs a log message if logLevel permits, and then panics
func (l *BasicLogger) PanicOnError(err error) {
	if err != nil {
		l.Panic(err)
	}
}

// Panicf outputs a formatted log message if logLevel permits, and then panics
func (l *BasicLogger) Panicf(f string, args ...interface{}) {
	l.Logf(LogLevelPanic, f, args...)
}

// Fatal outputs a log message if logLevel permits, and then exits with error code 1
func (l *BasicLogger) Fatal(args ...interface{}) {
	l.Log(LogLevelFatal, args...)
}

// Fatalf outputs a formatted log message if logLevel permits, and then exists with error code
func (l *BasicLogger) Fatalf(f string, args ...interface{}) {
	l.Logf(LogLevelFatal, f, args...)
}

// ELog outputs a formatted log message if logLevel permits
func (l *BasicLogger) ELog(args ...interface{}) {
	l.Log(LogLevelError, args...)
}

// ELogf outputs a formatted log message if logLevel permits
func (l *BasicLogger) ELogf(f string, args ...interface{}) {
	l.Logf(LogLevelError, f, args...)
}

// WLog outputs a formatted log message if logLevel permits
func (l *BasicLogger) WLog(args ...interface{}) {
	l.Log(LogLevelWarning, args...)
}

// WLogf outputs a formatted log message if logLevel permits
func (l *BasicLogger) WLogf(f string, args ...interface{}) {
	l.Logf(LogLevelWarning, f, args...)
}

// ILog outputs a formatted log message if logLevel permits
func (l *BasicLogger) ILog(args ...interface{}) {
	l.Log(LogLevelInfo, args...)
}

// ILogf outputs a formatted log message if logLevel permits
func (l *BasicLogger) ILogf(f string, args ...interface{}) {
	l.Logf(LogLevelInfo, f, args...)
}

// DLog outputs a formatted log message if logLevel permits
func (l *BasicLogger) DLog(args ...interface{}) {
	l.Log(LogLevelDebug, args...)
}

// DLogf outputs a formatted log message if logLevel permits
func (l *BasicLogger) DLogf(f string, args ...interface{}) {
	l.Logf(LogLevelDebug, f, args...)
}

// TLog outputs a formatted log message if logLevel permits
func (l *BasicLogger) TLog(args ...interface{}) {
	l.Log(LogLevelTrace, args...)
}

// TLogf outputs a formatted log message if logLevel permits
func (l *BasicLogger) TLogf(f string, args ...interface{}) {
	l.Logf(LogLevelTrace, f, args...)
}

// Error generates an error object with this logger's prefix
func (l *BasicLogger) Error(args ...interface{}) error {
	return errors.New(l.Sprint(args...))
}

// Errorf returns an error object with a description string that has the
// Logger's prefix
func (l *BasicLogger) Errorf(f string, args ...interface{}) error {
	return errors.New(l.Sprintf(f, args...))
}

// Sprintf returns a string that has the Logger's prefix
func (l *BasicLogger) Sprintf(f string, args ...interface{}) string {
	return l.prefixC + fmt.Sprintf(f, args...)
}

// Sprint returns a string that has the Logger's prefix
func (l *BasicLogger) Sprint(args ...interface{}) string {
	return l.prefixC + fmt.Sprint(args...)
}

// ELogError outputs an error message to a Logger iff logging level is enabled,
// and returns an error object with a description string that has the
// logger's prefix
func (l *BasicLogger) ELogError(args ...interface{}) error {
	return l.LogError(LogLevelError, args...)
}

// ELogErrorf outputs an error message to a Logger iff logging level is enabled,
// and returns an error object with a description string that has the
// logger's prefix
func (l *BasicLogger) ELogErrorf(f string, args ...interface{}) error {
	return l.LogErrorf(LogLevelError, f, args...)
}

// WLogError outputs an error message to a Logger iff logging level is enabled,
// and returns an error object with a description string that has the
// logger's prefix
func (l *BasicLogger) WLogError(args ...interface{}) error {
	return l.LogError(LogLevelWarning, args...)
}

// WLogErrorf outputs an error message to a Logger iff logging level is enabled,
// and returns an error object with a description string that has the
// logger's prefix
func (l *BasicLogger) WLogErrorf(f string, args ...interface{}) error {
	return l.LogErrorf(LogLevelWarning, f, args...)
}

// ILogError outputs an error message to a Logger iff logging level is enabled,
// and returns an error object with a description string that has the
// logger's prefix
func (l *BasicLogger) ILogError(args ...interface{}) error {
	return l.LogError(LogLevelInfo, args...)
}

// ILogErrorf outputs an error message to a Logger iff logging level is enabled,
// and returns an error object with a description string that has the
// logger's prefix
func (l *BasicLogger) ILogErrorf(f string, args ...interface{}) error {
	return l.LogErrorf(LogLevelInfo, f, args...)
}

// DLogError outputs an error message to a Logger iff DEBUG logging level is enabled,
// and returns an error object with a description string that has the
// logger's prefix
func (l *BasicLogger) DLogError(args ...interface{}) error {
	return l.LogError(LogLevelDebug, args...)
}

// DLogErrorf outputs an error message to a Logger iff DEBUG logging level is enabled,
// and returns an error object with a description string that has the
// logger's prefix
func (l *BasicLogger) DLogErrorf(f string, args ...interface{}) error {
	return l.LogErrorf(LogLevelDebug, f, args...)
}

// TLogError outputs an error message to a Logger iff logging level is enabled,
// and returns an error object with a description string that has the
// logger's prefix
func (l *BasicLogger) TLogError(args ...interface{}) error {
	return l.LogError(LogLevelTrace, args...)
}

// TLogErrorf outputs an error message to a Logger iff logging level is enabled,
// and returns an error object with a description string that has the
// logger's prefix
func (l *BasicLogger) TLogErrorf(f string, args ...interface{}) error {
	return l.LogErrorf(LogLevelTrace, f, args...)
}

// FlagsLogger is an interface for a logger that supports Flags() api. Note that log.Logger meets
// this requirement.
type FlagsLogger interface {
	Flags() int
}

// Flags returns the logger flags bits
func (l *BasicLogger) Flags() int {
	flagsLogger := l.logger.(FlagsLogger)

	var logFlags int
	if flagsLogger != nil {
		logFlags = flagsLogger.Flags()
	} else {
		logFlags = defaultLogFlags
	}
	return logFlags
}

// ForkLogf creates a new Logger that has an additional formatted string appended onto
// an existing logger's prefix (with ": " added between).
func (l *BasicLogger) ForkLogf(prefixFmt string, args ...interface{}) Logger {
	prefix := fmt.Sprintf(prefixFmt, args...)
	return l.ForkLog(prefix)
}

// ForkLog creates a new Logger that has an additional string appended onto
// an existing logger's prefix (with ": " added between).
func (l *BasicLogger) ForkLog(prefix string) Logger {
	var newPrefix string
	if prefix == "" {
		newPrefix = l.prefix
	} else {
		newPrefix = l.prefix + ": " + prefix
	}
	ll := NewLogWrapper(l.logger, newPrefix, l.GetLogLevel())
	return ll
}

// Prefix returns the Logger's prefix string (does not include ": " trailer)
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
