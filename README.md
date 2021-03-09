# logger
A simple golang package to make logging easier.

[![GoDoc](https://godoc.org/github.com/sammck-go/logger?status.svg)](https://godoc.org/github.com/sammck-go/logger)

**logger** makes it easy to add a multilevel logging interface to objects, so that log entries associated with an object are annotated with the object's identity.  New loggers can inherit from existing loggers and add more specific
annotations.

**logger** is derived from logging code in [chisel](https://github.com/jpillora/chisel) and inherits from its license. However, it is is published here as an independent module and as
such is not a proper fork, and does not track changes in the **chisel** project.

### Features

- Easy to use
- Multiple logging levels
- Drop-in to objects to implement logging

**Source**

```sh
$ go get -v github.com/sammck-go/logger
```

### Package Usage

<!-- render these help texts by hand,
  or use https://github.com/jpillora/md-tmpl
    with $ md-tmpl -w README.md -->

<!--tmpl:echo && godocdown -template ./.godocdown.template -->
```go
import "github.com/sammck-go/logger"
```

Package logger makes it easy to add a multilevel logging interface to objects,
so that log entries associated with an object are annotated with the object's
identity. New loggers can inherit from existing loggers and add more specific
annotations.

## Usage

```go
var LogLevelToName = func() map[LogLevel]string {
	var result = make(map[LogLevel]string)
	for i, name := range logLevelNames {
		result[LogLevel(i)] = name
	}
	return result
}()
```
LogLevelToName is a map of LogLevel values to log level names (lowercase)

```go
var NameToLogLevel = func() map[string]LogLevel {
	var result = make(map[string]LogLevel)
	for i, name := range logLevelNames {
		result[name] = LogLevel(i)
	}
	return result
}()
```
NameToLogLevel is a map of log level names (lowercase) to LogLevel values

#### type BasicLogger

```go
type BasicLogger struct {
}
```

BasicLogger is a logical log output stream with a level filter and a prefix
added to each output record.

#### func (*BasicLogger) CdError

```go
func (l *BasicLogger) CdError(calldepth int, args ...interface{}) error
```
CdError generates an error object with a given calldepth and this logger's
prefix. Arguments are formatted in the style of fmt.Sprint. Note: The raw
logger's prefix, if any, is not included. Note: calldepth is not currently used
but is included to allow future improvements

    that can put caller context into the error message.

#### func (*BasicLogger) CdErrorf

```go
func (l *BasicLogger) CdErrorf(calldepth int, f string, args ...interface{}) error
```
CdErrorf generates an error object with a given calldepth and this logger's
prefix. Arguments are formatted in the style of fmt.Sprintf. Note: The raw
logger's prefix, if any, is not included. Note: calldepth is not currently used
but is included to allow future improvements

    that can put caller context into the error message.

#### func (*BasicLogger) CdLog

```go
func (l *BasicLogger) CdLog(calldepth int, logLevel LogLevel, args ...interface{})
```
CdLog outputs to a Logger with a given calldepth if the given logLevel is
enabled. Then, if the given logLevel is LogLevelPanic or LogLevelFatal, exits
appropriately Arguments are formatted in the style of fmt.Sprint

#### func (*BasicLogger) CdLogError

```go
func (l *BasicLogger) CdLogError(calldepth int, logLevel LogLevel, args ...interface{}) error
```
CdLogError outputs an error message with a given calldepth to a Logger iff
logLevel is enabled, then returns an error object with a description string that
has the logger's prefix. If the given logLevel is LogLevelPanic or
LogLevelFatal, does not return. Arguments are formatted in the style of
fmt.Sprint.

#### func (*BasicLogger) CdLogErrorf

```go
func (l *BasicLogger) CdLogErrorf(calldepth int, logLevel LogLevel, f string, args ...interface{}) error
```
CdLogErrorf outputs an error message with a given calldepth to a Logger iff
logLevel is enabled, then returns an error object with a description string that
has the logger's prefix. If the given logLevel is LogLevelPanic or
LogLevelFatal, does not return. Arguments are formatted in the style of
fmt.Sprintf.

#### func (*BasicLogger) CdLogNoPrefix

```go
func (l *BasicLogger) CdLogNoPrefix(calldepth int, logLevel LogLevel, args ...interface{})
```
CdLogNoPrefix outputs to a Logger with provided call depth and without the
prefix (beyond the raw logger's prefix) if the given logLevel is enabled. Then,
if the given logLevel is LogLevelPanic or LogLevelFatal, exits appropriately.
Arguments are formatted in the style of fmt.Sprint

#### func (*BasicLogger) CdLogStrNoPrefix

```go
func (l *BasicLogger) CdLogStrNoPrefix(calldepth int, logLevel LogLevel, s string)
```
CdLogStrNoPrefix outputs a single string to a Logger with provided call depth
and without the prefix (beyond the raw logger's prefix) if the given logLevel is
enabled. Then, if the given logLevel is LogLevelPanic or LogLevelFatal, exits
appropriately.

#### func (*BasicLogger) CdLogf

```go
func (l *BasicLogger) CdLogf(calldepth int, logLevel LogLevel, f string, args ...interface{})
```
CdLogf outputs to a Logger with a given calldepth if the given logLevel is
enabled. Then, if the given logLevel is LogLevelPanic or LogLevelFatal, exits
appropriately Arguments are formatted in the style of fmt.Sprintf

#### func (*BasicLogger) CdLogfNoPrefix

```go
func (l *BasicLogger) CdLogfNoPrefix(calldepth int, logLevel LogLevel, f string, args ...interface{})
```
CdLogfNoPrefix outputs to a Logger with a given call depth and without the
prefix if the given logLevel is enabled. Then, if the given logLevel is
LogLevelPanic or LogLevelFatal, exits appropriately. Arguments are formatted in
the style of fmt.Sprintf

#### func (*BasicLogger) CdPanic

```go
func (l *BasicLogger) CdPanic(calldepth int, args ...interface{})
```
CdPanic outputs a log message with a given call depth if LogLevelPanic is
enabled, and then panics. Arguments are formatted in the style of fmt.Sprint.

#### func (*BasicLogger) CdPanicOnError

```go
func (l *BasicLogger) CdPanicOnError(calldepth int, err error)
```
CdPanicOnError does nothing if err is nil; otherwise outputs a log message with
a given call depth if LogLevelPanic is enabled, and then panics.

#### func (*BasicLogger) CdPrint

```go
func (l *BasicLogger) CdPrint(calldepth int, args ...interface{})
```
CdPrint writes arguments to a Logger with a provided call depth. Arguments are
formatted in the style of fmt.Sprint()

#### func (*BasicLogger) CdPrintf

```go
func (l *BasicLogger) CdPrintf(calldepth int, f string, args ...interface{})
```
CdPrintf outputs formatted text to a Logger with a provided call depth, in the
style of fmt.Sprintf

#### func (*BasicLogger) CdRawOutput

```go
func (l *BasicLogger) CdRawOutput(calldepth int, s string)
```
CdRawOutput is the lowest level log output method; it writes the output for a
logging event without any additional prefixes not supplied by the raw logger.
The string s contains the text to print after the prefix specified by the flags
of the raw Logger. A newline is appended if the last character of s is not
already a newline. Calldepth is used to recover the call frame, PC, filename,
etc. If set to 1, it will display context for the immediate caller of
CdRawOutput; if set to 2, the caller's caller, etc.

#### func (*BasicLogger) CdSprint

```go
func (l *BasicLogger) CdSprint(calldepth int, args ...interface{}) string
```
CdSprint returns a string that has the Logger's prefix, with a given call depth.
Arguments are formatted in the style of fmt.Sprint. Note: The raw logger's
prefix, if any, is not included. Note: calldepth is not currently used but is
included to allow future improvements

    that can put caller context into the prefix.

#### func (*BasicLogger) CdSprintf

```go
func (l *BasicLogger) CdSprintf(calldepth int, f string, args ...interface{}) string
```
CdSprintf returns a string that has the Logger's prefix, with a given call
depth. Arguments are formatted in the style of fmt.Sprintf. Note: The raw
logger's prefix, if any, is not included. Note: calldepth is not currently used
but is included to allow future improvements

    that can put caller context into the prefix.

#### func (*BasicLogger) DLog

```go
func (l *BasicLogger) DLog(args ...interface{})
```
DLog outputs a formatted log message if LogLevelDebug is enabled. Arguments are
formatted in the style of fmt.Sprint.

#### func (*BasicLogger) DLogError

```go
func (l *BasicLogger) DLogError(args ...interface{}) error
```
DLogError outputs an error message to a Logger iff LogLevelDebug is enabled, and
returns an error object with a description string that has the logger's prefix.
Arguments are formatted in the style of fmt.Sprint.

#### func (*BasicLogger) DLogErrorf

```go
func (l *BasicLogger) DLogErrorf(f string, args ...interface{}) error
```
DLogErrorf outputs an error message to a Logger iff LogLevelDebug is enabled,
and returns an error object with a description string that has the logger's
prefix. Arguments are formatted in the style of fmt.Sprintf.

#### func (*BasicLogger) DLogf

```go
func (l *BasicLogger) DLogf(f string, args ...interface{})
```
DLogf outputs a formatted log message if LogLevelDebug is enabled. Arguments are
formatted in the style of fmt.Sprintf.

#### func (*BasicLogger) ELog

```go
func (l *BasicLogger) ELog(args ...interface{})
```
ELog outputs a formatted log message if LogLevelError is enabled. Arguments are
formatted in the style of fmt.Sprint.

#### func (*BasicLogger) ELogError

```go
func (l *BasicLogger) ELogError(args ...interface{}) error
```
ELogError outputs an error message to a Logger iff LogLevelError is enabled, and
returns an error object with a description string that has the logger's prefix.
Arguments are formatted in the style of fmt.Sprint.

#### func (*BasicLogger) ELogErrorf

```go
func (l *BasicLogger) ELogErrorf(f string, args ...interface{}) error
```
ELogErrorf outputs an error message to a Logger iff LogLevelError is enabled,
and returns an error object with a description string that has the logger's
prefix. Arguments are formatted in the style of fmt.Sprintf.

#### func (*BasicLogger) ELogf

```go
func (l *BasicLogger) ELogf(f string, args ...interface{})
```
ELogf outputs a formatted log message if LogLevelError is enabled. Arguments are
formatted in the style of fmt.Sprintf.

#### func (*BasicLogger) Error

```go
func (l *BasicLogger) Error(args ...interface{}) error
```
Error generates an error object with this logger's prefix. Arguments are
formatted in the style of fmt.Sprint. Note: The raw logger's prefix, if any, is
not included.

#### func (*BasicLogger) Errorf

```go
func (l *BasicLogger) Errorf(f string, args ...interface{}) error
```
Errorf generates an error object with this logger's prefix. Arguments are
formatted in the style of fmt.Sprintf. Note: The raw logger's prefix, if any, is
not included.

#### func (*BasicLogger) Fatal

```go
func (l *BasicLogger) Fatal(args ...interface{})
```
Fatal outputs a log message if LogLevelFatal is enabled, and then exits with
error code 1. Arguments are formatted in the style of fmt.Sprint.

#### func (*BasicLogger) Fatalf

```go
func (l *BasicLogger) Fatalf(f string, args ...interface{})
```
Fatalf outputs a formatted log message if LogLevelFatal is enabled, and then
exits with error code. Arguments are formatted in the style of fmt.Sprintf.

#### func (*BasicLogger) ForkLog

```go
func (l *BasicLogger) ForkLog(args ...interface{}) Logger
```
ForkLog creates a new Logger that has an additional formatted string appended
onto an existing logger's prefix (with ": " added between). Arguments are
formatted in the style of fmt.Sprint

#### func (*BasicLogger) ForkLogStr

```go
func (l *BasicLogger) ForkLogStr(prefix string) Logger
```
ForkLogStr creates a new Logger that has an additional string appended onto an
existing logger's prefix (with ": " added between).

#### func (*BasicLogger) ForkLogf

```go
func (l *BasicLogger) ForkLogf(prefixFmt string, args ...interface{}) Logger
```
ForkLogf creates a new Logger that has an additional formatted string appended
onto an existing logger's prefix (with ": " added between). Arguments are
formatted in the style of fmt.Sprintf

#### func (*BasicLogger) GetLogLevel

```go
func (l *BasicLogger) GetLogLevel() LogLevel
```
GetLogLevel returns the log level

#### func (*BasicLogger) ILog

```go
func (l *BasicLogger) ILog(args ...interface{})
```
ILog outputs a formatted log message if LogLevelInfo is enabled. Arguments are
formatted in the style of fmt.Sprint.

#### func (*BasicLogger) ILogError

```go
func (l *BasicLogger) ILogError(args ...interface{}) error
```
ILogError outputs an error message to a Logger iff LogLevelInfo is enabled, and
returns an error object with a description string that has the logger's prefix.
Arguments are formatted in the style of fmt.Sprint.

#### func (*BasicLogger) ILogErrorf

```go
func (l *BasicLogger) ILogErrorf(f string, args ...interface{}) error
```
ILogErrorf outputs an error message to a Logger iff LogLevelInfo is enabled, and
returns an error object with a description string that has the logger's prefix.
Arguments are formatted in the style of fmt.Sprintf.

#### func (*BasicLogger) ILogf

```go
func (l *BasicLogger) ILogf(f string, args ...interface{})
```
ILogf outputs a formatted log message if LogLevelInfo is enabled. Arguments are
formatted in the style of fmt.Sprintf.

#### func (*BasicLogger) Log

```go
func (l *BasicLogger) Log(logLevel LogLevel, args ...interface{})
```
Log outputs to a Logger if the given logLevel is enabled. Then, if the given
logLevel is LogLevelPanic or LogLevelFatal, exits appropriately Arguments are
formatted in the style of fmt.Sprint

#### func (*BasicLogger) LogError

```go
func (l *BasicLogger) LogError(logLevel LogLevel, args ...interface{}) error
```
LogError outputs an error message to a Logger iff logLevel is enabled, then
returns an error object with a description string that has the logger's prefix.
If the given logLevel is LogLevelPanic or LogLevelFatal, does not return.
Arguments are formatted in the style of fmt.Sprint.

#### func (*BasicLogger) LogErrorf

```go
func (l *BasicLogger) LogErrorf(logLevel LogLevel, f string, args ...interface{}) error
```
LogErrorf outputs an error message to a Logger iff logLevel is enabled, then
returns an error object with a description string that has the logger's prefix.
If the given logLevel is LogLevelPanic or LogLevelFatal, does not return.
Arguments are formatted in the style of fmt.Sprintf.

#### func (*BasicLogger) LogNoPrefix

```go
func (l *BasicLogger) LogNoPrefix(logLevel LogLevel, args ...interface{})
```
LogNoPrefix outputs to a Logger without the prefix if the given logLevel is
enabled. Then, if the given logLevel is LogLevelPanic or LogLevelFatal, exits
appropriately Arguments are formatted in the style of fmt.Sprint

#### func (*BasicLogger) LogStrNoPrefix

```go
func (l *BasicLogger) LogStrNoPrefix(logLevel LogLevel, s string)
```
LogStrNoPrefix outputs a single string to a Logger without the prefix (beyond
the raw logger's prefix) if the given logLevel is enabled. Then, if the given
logLevel is LogLevelPanic or LogLevelFatal, exits appropriately.

#### func (*BasicLogger) Logf

```go
func (l *BasicLogger) Logf(logLevel LogLevel, f string, args ...interface{})
```
Logf outputs to a Logger if the given logLevel is enabled. Then, if the given
logLevel is LogLevelPanic or LogLevelFatal, exits appropriately Arguments are
formatted in the style of fmt.Sprintf

#### func (*BasicLogger) LogfNoPrefix

```go
func (l *BasicLogger) LogfNoPrefix(logLevel LogLevel, f string, args ...interface{})
```
LogfNoPrefix outputs to a Logger without the prefix if the given logLevel is
enabled. Then, if the given logLevel is LogLevelPanic or LogLevelFatal, exits
appropriately Arguments are formatted in the style of fmt.Sprintf

#### func (*BasicLogger) Output

```go
func (l *BasicLogger) Output(calldepth int, s string) error
```
Output is the compatible with log.Logger.Output, to make this a RawLogger

#### func (*BasicLogger) Panic

```go
func (l *BasicLogger) Panic(args ...interface{})
```
Panic outputs a log message if LogLevelPanic is enabled, and then panics.
Arguments are formatted in the style of fmt.Sprint.

#### func (*BasicLogger) PanicOnError

```go
func (l *BasicLogger) PanicOnError(err error)
```
PanicOnError does nothing if err is nil; otherwise outputs a log message if
LogLevelPanic is enabled, and then panics.

#### func (*BasicLogger) Panicf

```go
func (l *BasicLogger) Panicf(f string, args ...interface{})
```
Panicf outputs a formatted log message if LogLevelPanic is enabled, and then
panics. Arguments are formatted in the style of fmt.Sprintf.

#### func (*BasicLogger) Prefix

```go
func (l *BasicLogger) Prefix() string
```
Prefix returns the Logger's prefix string (does not include ": " trailer) Does
not include the raw logger's prefix, if any.

#### func (*BasicLogger) Print

```go
func (l *BasicLogger) Print(args ...interface{})
```
Print outputs to a Logger in the style of fmt.Sprint()

#### func (*BasicLogger) Printf

```go
func (l *BasicLogger) Printf(f string, args ...interface{})
```
Printf outputs to a Logger in the style of fmt.Sprintf

#### func (*BasicLogger) SetLogLevel

```go
func (l *BasicLogger) SetLogLevel(logLevel LogLevel)
```
SetLogLevel sets the log level

#### func (*BasicLogger) Sprint

```go
func (l *BasicLogger) Sprint(args ...interface{}) string
```
Sprint returns a string that has the Logger's prefix Arguments are formatted in
the style of fmt.Sprint. Note: The raw logger's prefix, if any, is not included.

#### func (*BasicLogger) Sprintf

```go
func (l *BasicLogger) Sprintf(f string, args ...interface{}) string
```
Sprintf returns a string that has the Logger's prefix Arguments are formatted in
the style of fmt.Sprintf. Note: The raw logger's prefix, if any, is not
included.

#### func (*BasicLogger) TLog

```go
func (l *BasicLogger) TLog(args ...interface{})
```
TLog outputs a formatted log message if LogLevelTrace is enabled. Arguments are
formatted in the style of fmt.Sprint.

#### func (*BasicLogger) TLogError

```go
func (l *BasicLogger) TLogError(args ...interface{}) error
```
TLogError outputs an error message to a Logger iff LogLevelTrace is enabled, and
returns an error object with a description string that has the logger's prefix.
Arguments are formatted in the style of fmt.Sprint.

#### func (*BasicLogger) TLogErrorf

```go
func (l *BasicLogger) TLogErrorf(f string, args ...interface{}) error
```
TLogErrorf outputs an error message to a Logger iff LogLevelTrace is enabled,
and returns an error object with a description string that has the logger's
prefix. Arguments are formatted in the style of fmt.Sprintf.

#### func (*BasicLogger) TLogf

```go
func (l *BasicLogger) TLogf(f string, args ...interface{})
```
TLogf outputs a formatted log message if LogLevelTrace is enabled. Arguments are
formatted in the style of fmt.Sprintf.

#### func (*BasicLogger) WLog

```go
func (l *BasicLogger) WLog(args ...interface{})
```
WLog outputs a formatted log message if LogLevelWarning is enabled. Arguments
are formatted in the style of fmt.Sprint.

#### func (*BasicLogger) WLogError

```go
func (l *BasicLogger) WLogError(args ...interface{}) error
```
WLogError outputs an error message to a Logger iff LogLevelWarning is enabled,
and returns an error object with a description string that has the logger's
prefix. Arguments are formatted in the style of fmt.Sprint.

#### func (*BasicLogger) WLogErrorf

```go
func (l *BasicLogger) WLogErrorf(f string, args ...interface{}) error
```
WLogErrorf outputs an error message to a Logger iff LogLevelWarning is enabled,
and returns an error object with a description string that has the logger's
prefix. Arguments are formatted in the style of fmt.Sprintf.

#### func (*BasicLogger) WLogf

```go
func (l *BasicLogger) WLogf(f string, args ...interface{})
```
WLogf outputs a formatted log message if LogLevelWarning is enabled. Arguments
are formatted in the style of fmt.Sprintf.

#### type Config

```go
type Config struct {
}
```

Config provides configuration options for contruction of a Logger. The
constructed object is immutable after it is constructed by NewConfig.

#### func  NewConfig

```go
func NewConfig(opts ...ConfigOption) *Config
```
NewConfig creates a Config object from provided options. The resulting object
can be passed to New using WithConfig, or directly to NewWithConfig.

#### func (*Config) Refine

```go
func (cfg *Config) Refine(opts ...ConfigOption) *Config
```
Refine creates a new Config object by applying ConfigOptions to an existing
config.

#### type ConfigOption

```go
type ConfigOption func(*Config)
```

ConfigOption is an opaque configuration option setter created by one of the With
functions. It follows the Golang "options" pattern.

#### func  WithClearLogFlags

```go
func WithClearLogFlags(flag int) ConfigOption
```
WithClearLogFlags clears all log flags for the new logger. By default, log.Ldate
and log.Ltime are set. Note that flags are ignored if WithLogger() is provided.
See log.Logger for a description of flag bits.

#### func  WithConfig

```go
func WithConfig(other *Config) ConfigOption
```
WithConfig allows initialization of a new configuration object starting with an
existing one, and incremental initialization of configuration separately from
initialization of the PidFile. If provided, this option should be appear first
in the option list, since it replaces all configuration values.

#### func  WithLUTC

```go
func WithLUTC() ConfigOption
```
WithLUTC adds log flag log.LUTC for the new logger, without affecting any other
flags. This causes dates and times in the raw logger to appear in UTC format. By
default, local time is used. Note that flags are ignored if WithLogger() is
provided. See log.Logger for a description of flag bits.

#### func  WithLdate

```go
func WithLdate() ConfigOption
```
WithLdate adds log flag log.Ldate for the new logger, without affecting any
other flags. This enables a date in the timestamp for the raw logger. By
default, log.Ldate and log.Ltime are set. Note that flags are ignored if
WithLogger() is provided. See log.Logger for a description of flag bits.

#### func  WithLlongfile

```go
func WithLlongfile() ConfigOption
```
WithLlongfile adds log flag log.Llongfile for the new logger, without affecting
any other flags. This causes a long source filename and line number to appear in
log entries. By default, no file/line appears. Note that flags are ignored if
WithLogger() is provided. See log.Logger for a description of flag bits.

#### func  WithLmicroseconds

```go
func WithLmicroseconds() ConfigOption
```
WithLmicroseconds adds log flag log.Lmicroseconds for the new logger, without
affecting any other flags. This causes timestamps to appear with microsecond
resolution in log entries. By default, one second resolution is used. Note that
flags are ignored if WithLogger() is provided. See log.Logger for a description
of flag bits.

#### func  WithLogFlags

```go
func WithLogFlags(flag int) ConfigOption
```
WithLogFlags adds log flags for the new logger, without removing any existing
flags. By default, log.Ldate and log.Ltime are set. Note that flags are ignored
if WithLogger() is provided. See log.Logger for a description of flag bits.

#### func  WithLogLevel

```go
func WithLogLevel(logLevel LogLevel) ConfigOption
```
WithLogLevel sets the prefix for the new logger. By default, no prefix is added.

#### func  WithLogger

```go
func WithLogger(parentLogger RawLogger) ConfigOption
```
WithLogger sets the parent logger for new logger. By default, a new logger to
stderr is created. This setting replaces any prior effect of WithWriter().

#### func  WithLshortfile

```go
func WithLshortfile() ConfigOption
```
WithLshortfile adds log flag log.Lshortfile for the new logger, without
affecting any other flags. This causes a short source filename and line number
to appear in log entries. By default, no file/line appears. Note that flags are
ignored if WithLogger() is provided. See log.Logger for a description of flag
bits.

#### func  WithLtime

```go
func WithLtime() ConfigOption
```
WithLtime adds log flag log.Ltime for the new logger, without affecting any
other flags. This enables a time of day in the raw logger. By default, log.Ldate
and log.Ltime are set. Note that flags are ignored if WithLogger() is provided.
See log.Logger for a description of flag bits.

#### func  WithPrefix

```go
func WithPrefix(prefix string) ConfigOption
```
WithPrefix sets the prefix for the new logger. By default, no prefix is added.

#### func  WithReplaceLogFlags

```go
func WithReplaceLogFlags(flag int) ConfigOption
```
WithReplaceLogFlags replaces all log flags for the new logger. By default,
log.Ldate and log.Ltime are set. Note that flags are ignored if WithLogger() is
provided. See log.Logger for a description of flag bits.

#### func  WithWriter

```go
func WithWriter(logWriter io.Writer) ConfigOption
```
WithWriter sets the io,Writer to which log output will be sent. By default, log
output will be sent to stderr. This setting replaces any prior effect of
WithLogger().

#### func  WithoutLUTC

```go
func WithoutLUTC() ConfigOption
```
WithoutLUTC removess log flag log.LUTC from the new logger, without affecting
any other flags. This causes dates and times in the raw logger to appear in
local timezone format. This is the default setting. Note that flags are ignored
if WithLogger() is provided. See log.Logger for a description of flag bits.

#### func  WithoutLdate

```go
func WithoutLdate() ConfigOption
```
WithoutLdate removes log flag log.Ldate from the new logger, without affecting
any other flags. This disables a date in the timestamp for the raw logger. By
default, log.Ldate and log.Ltime are set. Note that flags are ignored if
WithLogger() is provided. See log.Logger for a description of flag bits.

#### func  WithoutLlongfile

```go
func WithoutLlongfile() ConfigOption
```
WithoutLlongfile removes log flag log.Llongfile from the new logger, without
affecting any other flags. This disables a long source filename and line number
from appearing in log entries. This is the default setting. Note that flags are
ignored if WithLogger() is provided. See log.Logger for a description of flag
bits.

#### func  WithoutLmicroseconds

```go
func WithoutLmicroseconds() ConfigOption
```
WithoutLmicroseconds removes log flag log.Lmicroseconds from the new logger,
without affecting any other flags. This causes a timestamps to appear with
one-second resolution in log entries. This is the default setting. Note that
flags are ignored if WithLogger() is provided. See log.Logger for a description
of flag bits.

#### func  WithoutLogFlags

```go
func WithoutLogFlags(flag int) ConfigOption
```
WithoutLogFlags clears specified log flags for the new logger, without affecting
any other flags. By default, log.Ldate and log.Ltime are set. Note that flags
are ignored if WithLogger() is provided. See log.Logger for a description of
flag bits.

#### func  WithoutLshortfile

```go
func WithoutLshortfile() ConfigOption
```
WithoutLshortfile removes log flag log.Lshortfile from the new logger, without
affecting any other flags. This disables a short source filename and line number
from appearing in log entries. This is the default setting. Note that flags are
ignored if WithLogger() is provided. See log.Logger for a description of flag
bits.

#### func  WithoutLtime

```go
func WithoutLtime() ConfigOption
```
WithoutLtime removes log flag log.Ltime from the new logger, without affecting
any other flags. This disables a time of day in the raw logger. By default,
log.Ldate and log.Ltime are set. Note that flags are ignored if WithLogger() is
provided. See log.Logger for a description of flag bits.

#### type GetLogLeveler

```go
type GetLogLeveler interface {
	GetLogLevel() LogLevel
}
```

GetLogLeveler is an interface for a logger that supports GetLogLevel()

#### type LogLevel

```go
type LogLevel int
```

LogLevel specifies the level of spew that shoud go to the log

```go
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
```

#### func  StringToLogLevel

```go
func StringToLogLevel(s string) LogLevel
```
StringToLogLevel converts a string to a LogLevel. LogLevelUnknown is returned
for unrecognized strings. upper/lowercase is accepted.

#### func (*LogLevel) FromString

```go
func (x *LogLevel) FromString(s string) error
```
FromString initiales a LogLevel from a string. Upper/lowercase are accepted.

#### func (*LogLevel) String

```go
func (x *LogLevel) String() string
```
String converts a LogLevel to a string (lowercase)

#### type Logger

```go
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
```

Logger is an interface for a logging component that supports logging levels and
prefix forking

#### func  New

```go
func New(opts ...ConfigOption) (Logger, error)
```
New creates a new Logger object from supplied configuration options

#### func  NewLogWrapper

```go
func NewLogWrapper(logger RawLogger, prefix string, logLevel LogLevel) Logger
```
NewLogWrapper creates a new Logger that wraps an existing base RawLogger with an
optional additional prefix and an adjusted loglevel. Note that the logging level
cannot be effectively increased from the logging level of the base logger, since
the base logger since the base logger will filter items passed to it. As an
optimization, the wrapper will limit the loglevel to the base logger's level if
it implements GetLogLevel(). If the base logger's loglevel subsequently changes,
it is the caller's responsibility to adjust the new wrapper's loglevel if
desired.

#### func  NewWithConfig

```go
func NewWithConfig(cfg *Config) (Logger, error)
```
NewWithConfig creates a new Logger object from a configuration

#### type RawLogger

```go
type RawLogger interface {
	// Output writes the output for a logging event. The string s contains
	// the text to print after the prefix specified by the flags of the
	// Logger. A newline is appended if the last character of s is not
	// already a newline. Calldepth is used to recover the PC, filename,
	// etc.  If set to 1, it will display context for the immediate caller
	// of Output; if set to 2, the caller's caller, etc.
	Output(calldepth int, s string) error
}
```

RawLogger is a minimal logging interface for an underlying logging component. A
full-featured Logger implementation can be wrapped around anything that
implements this interface. Note that this is a subset of log.Logger, so any
standard golang logger can be used.
<!--/tmpl-->

### Contributing

- http://golang.org/doc/code.html
- http://golang.org/doc/effective_go.html

### Changelog

- `1.0` - Initial release.
- `1.1` - Source refactoring. Fixed log.Lshortfile and log.Llongfile to display the correct source file. Better example.

### Todo

- Better tests

#### MIT License

Copyright © 2017 Jaime Pillora &lt;dev@jpillora.com&gt;\
Copyright © 2021 Sam McKelvie &lt;dev@mckelvie.org&gt;

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
'Software'), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
