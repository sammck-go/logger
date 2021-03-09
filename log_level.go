package logger

import (
	"fmt"
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

// NameToLogLevel is a map of log level names (lowercase) to LogLevel values
var NameToLogLevel = func() map[string]LogLevel {
	var result = make(map[string]LogLevel)
	for i, name := range logLevelNames {
		result[name] = LogLevel(i)
	}
	return result
}()

// LogLevelToName is a map of LogLevel values to log level names (lowercase)
var LogLevelToName = func() map[LogLevel]string {
	var result = make(map[LogLevel]string)
	for i, name := range logLevelNames {
		result[LogLevel(i)] = name
	}
	return result
}()

// StringToLogLevel converts a string to a LogLevel. LogLevelUnknown is returned for
// unrecognized strings.  upper/lowercase is accepted.
func StringToLogLevel(s string) LogLevel {
	result, ok := NameToLogLevel[strings.ToLower(s)]
	if !ok {
		result = LogLevelUnknown
	}
	return result
}

// String converts a LogLevel to a string (lowercase)
func (x *LogLevel) String() string {
	y := *x
	if y < LogLevelUnknown || y > LogLevelTrace {
		y = LogLevelUnknown
	}
	return logLevelNames[y]
}

// FromString initiales a LogLevel from a string.  Upper/lowercase are accepted.
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
