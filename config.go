package logger

import (
	"io"
	"log"
)

// Config provides configuration options for contruction of a Logger.  The constructed object is immutable
// after it is constructed by NewConfig.
type Config struct {
	prefix       string
	flag         int
	logLevel     LogLevel
	parentLogger RawLogger
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

// WithClearLogFlags clears all log flags for the new logger. By default, log.Ldate and log.Ltime
// are set.  Note that flags are ignored if WithLogger() is provided. See log.Logger for a description of flag bits.
func WithClearLogFlags(flag int) ConfigOption {
	return func(cfg *Config) {
		cfg.flag = 0
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

// WithLdate adds log flag log.Ldate for the new logger, without affecting any other flags. This enables a date in the timestamp for
// the raw logger. By default, log.Ldate and log.Ltime are set. Note that flags are ignored if WithLogger() is provided.
// See log.Logger for a description of flag bits.
func WithLdate() ConfigOption {
	return WithLogFlags(log.Ldate)
}

// WithoutLdate removes log flag log.Ldate from the new logger, without affecting any other flags. This disables a date in the
// timestamp for the raw logger. By default, log.Ldate and log.Ltime
// are set. Note that flags are ignored if WithLogger() is provided. See log.Logger for a description of flag bits.
func WithoutLdate() ConfigOption {
	return WithoutLogFlags(log.Ldate)
}

// WithLtime adds log flag log.Ltime for the new logger, without affecting any other flags. This enables a time of day in
// the raw logger. By default, log.Ldate and log.Ltime
// are set. Note that flags are ignored if WithLogger() is provided. See log.Logger for a description of flag bits.
func WithLtime() ConfigOption {
	return WithLogFlags(log.Ltime)
}

// WithoutLtime removes log flag log.Ltime from the new logger, without affecting any other flags. This disables a time of
// day in the raw logger. By default, log.Ldate and log.Ltime
// are set. Note that flags are ignored if WithLogger() is provided. See log.Logger for a description of flag bits.
func WithoutLtime() ConfigOption {
	return WithoutLogFlags(log.Ltime)
}

// WithLUTC adds log flag log.LUTC for the new logger, without affecting any other flags. This causes dates and times in
// the raw logger to appear in UTC format. By default, local time is used. Note that flags are ignored if WithLogger()
// is provided. See log.Logger for a description of flag bits.
func WithLUTC() ConfigOption {
	return WithLogFlags(log.LUTC)
}

// WithoutLUTC removess log flag log.LUTC from the new logger, without affecting any other flags. This causes dates and times in
// the raw logger to appear in local timezone format. This is the default setting. Note that flags are ignored if WithLogger()
// is provided. See log.Logger for a description of flag bits.
func WithoutLUTC() ConfigOption {
	return WithoutLogFlags(log.LUTC)
}

// WithLlongfile adds log flag log.Llongfile for the new logger, without affecting any other flags. This causes a long
// source filename and line number to appear in log entries. By default, no file/line appears. Note that flags are ignored if WithLogger()
// is provided. See log.Logger for a description of flag bits.
func WithLlongfile() ConfigOption {
	return WithLogFlags(log.Llongfile)
}

// WithoutLlongfile removes log flag log.Llongfile from the new logger, without affecting any other flags. This disables a long
// source filename and line number from appearing in log entries. This is the default setting. Note that flags are ignored if WithLogger()
// is provided. See log.Logger for a description of flag bits.
func WithoutLlongfile() ConfigOption {
	return WithoutLogFlags(log.Llongfile)
}

// WithLshortfile adds log flag log.Lshortfile for the new logger, without affecting any other flags. This causes a short
// source filename and line number to appear in log entries. By default, no file/line appears. Note that flags are ignored if WithLogger()
// is provided. See log.Logger for a description of flag bits.
func WithLshortfile() ConfigOption {
	return WithLogFlags(log.Lshortfile)
}

// WithoutLshortfile removes log flag log.Lshortfile from the new logger, without affecting any other flags. This disables a short
// source filename and line number from appearing in log entries. This is the default setting. Note that flags are ignored if WithLogger()
// is provided. See log.Logger for a description of flag bits.
func WithoutLshortfile() ConfigOption {
	return WithoutLogFlags(log.Lshortfile)
}

// WithLmicroseconds adds log flag log.Lmicroseconds for the new logger, without affecting any other flags. This causes timestamps
// to appear with microsecond resolution in log entries. By default, one second resolution is used. Note that flags are ignored if WithLogger()
// is provided. See log.Logger for a description of flag bits.
func WithLmicroseconds() ConfigOption {
	return WithLogFlags(log.Lmicroseconds)
}

// WithoutLmicroseconds removes log flag log.Lmicroseconds from the new logger, without affecting any other flags. This causes a timestamps
// to appear with one-second resolution in log entries. This is the default setting. Note that flags are ignored if WithLogger()
// is provided. See log.Logger for a description of flag bits.
func WithoutLmicroseconds() ConfigOption {
	return WithoutLogFlags(log.Lmicroseconds)
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
func WithLogger(parentLogger RawLogger) ConfigOption {
	return func(cfg *Config) {
		cfg.parentLogger = parentLogger
		cfg.logWriter = nil
	}
}
