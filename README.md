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

#### type BasicLogger

```go
type BasicLogger struct {
}
```

BasicLogger is a logical log output stream with a level filter and a prefix
added to each output record.

#### func (*BasicLogger) DLog

```go
func (l *BasicLogger) DLog(args ...interface{})
```
DLog outputs a formatted log message if logLevel permits

#### func (*BasicLogger) DLogError

```go
func (l *BasicLogger) DLogError(args ...interface{}) error
```
DLogError outputs an error message to a Logger iff DEBUG logging level is
enabled, and returns an error object with a description string that has the
logger's prefix

#### func (*BasicLogger) DLogErrorf

```go
func (l *BasicLogger) DLogErrorf(f string, args ...interface{}) error
```
DLogErrorf outputs an error message to a Logger iff DEBUG logging level is
enabled, and returns an error object with a description string that has the
logger's prefix

#### func (*BasicLogger) DLogf

```go
func (l *BasicLogger) DLogf(f string, args ...interface{})
```
DLogf outputs a formatted log message if logLevel permits

#### func (*BasicLogger) ELog

```go
func (l *BasicLogger) ELog(args ...interface{})
```
ELog outputs a formatted log message if logLevel permits

#### func (*BasicLogger) ELogError

```go
func (l *BasicLogger) ELogError(args ...interface{}) error
```
ELogError outputs an error message to a Logger iff logging level is enabled, and
returns an error object with a description string that has the logger's prefix

#### func (*BasicLogger) ELogErrorf

```go
func (l *BasicLogger) ELogErrorf(f string, args ...interface{}) error
```
ELogErrorf outputs an error message to a Logger iff logging level is enabled,
and returns an error object with a description string that has the logger's
prefix

#### func (*BasicLogger) ELogf

```go
func (l *BasicLogger) ELogf(f string, args ...interface{})
```
ELogf outputs a formatted log message if logLevel permits

#### func (*BasicLogger) Error

```go
func (l *BasicLogger) Error(args ...interface{}) error
```
Error generates an error object with this logger's prefix

#### func (*BasicLogger) Errorf

```go
func (l *BasicLogger) Errorf(f string, args ...interface{}) error
```
Errorf returns an error object with a description string that has the Logger's
prefix

#### func (*BasicLogger) Fatal

```go
func (l *BasicLogger) Fatal(args ...interface{})
```
Fatal outputs a log message if logLevel permits, and then exits with error code
1

#### func (*BasicLogger) Fatalf

```go
func (l *BasicLogger) Fatalf(f string, args ...interface{})
```
Fatalf outputs a formatted log message if logLevel permits, and then exists with
error code

#### func (*BasicLogger) Flags

```go
func (l *BasicLogger) Flags() int
```
Flags returns the logger flags bits

#### func (*BasicLogger) Fork

```go
func (l *BasicLogger) Fork(prefix string, args ...interface{}) Logger
```
Fork creates a new Logger that has an additional formatted string appended onto
an existing logger's prefix (with ": " added between)

#### func (*BasicLogger) GetLogLevel

```go
func (l *BasicLogger) GetLogLevel() LogLevel
```
GetLogLevel returns the log level

#### func (*BasicLogger) ILog

```go
func (l *BasicLogger) ILog(args ...interface{})
```
ILog outputs a formatted log message if logLevel permits

#### func (*BasicLogger) ILogError

```go
func (l *BasicLogger) ILogError(args ...interface{}) error
```
ILogError outputs an error message to a Logger iff logging level is enabled, and
returns an error object with a description string that has the logger's prefix

#### func (*BasicLogger) ILogErrorf

```go
func (l *BasicLogger) ILogErrorf(f string, args ...interface{}) error
```
ILogErrorf outputs an error message to a Logger iff logging level is enabled,
and returns an error object with a description string that has the logger's
prefix

#### func (*BasicLogger) ILogf

```go
func (l *BasicLogger) ILogf(f string, args ...interface{})
```
ILogf outputs a formatted log message if logLevel permits

#### func (*BasicLogger) Log

```go
func (l *BasicLogger) Log(logLevel LogLevel, args ...interface{})
```
Log outputs to a Logger if the given logLevel is enabled. Then, if the given
logLevel is LogLevelPanic or LogLevelFatal, exits appropriately

#### func (*BasicLogger) LogError

```go
func (l *BasicLogger) LogError(logLevel LogLevel, args ...interface{}) error
```
LogError outputs an error message to a Logger iff logging level is enabled, and
returns an error object with a description string that has the logger's prefix

#### func (*BasicLogger) LogErrorf

```go
func (l *BasicLogger) LogErrorf(logLevel LogLevel, f string, args ...interface{}) error
```
LogErrorf outputs an error message to a Logger iff logging level is enabled, and
returns an error object with a description string that has the logger's prefix

#### func (*BasicLogger) LogNoPrefix

```go
func (l *BasicLogger) LogNoPrefix(logLevel LogLevel, args ...interface{})
```
LogNoPrefix outputs to a Logger without the prefix if the given logLevel is
enabled. Then, if the given logLevel is LogLevelPanic or LogLevelFatal, exits
appropriately

#### func (*BasicLogger) Logf

```go
func (l *BasicLogger) Logf(logLevel LogLevel, f string, args ...interface{})
```
Logf outputs to a Logger if the given logLevel is enabled. Then, if the given
logLevel is LogLevelPanic or LogLevelFatal, exits appropriately

#### func (*BasicLogger) LogfNoPrefix

```go
func (l *BasicLogger) LogfNoPrefix(logLevel LogLevel, f string, args ...interface{})
```
LogfNoPrefix outputs to a Logger without the prefix if the given logLevel is
enabled. Then, if the given logLevel is LogLevelPanic or LogLevelFatal, exits
appropriately

#### func (*BasicLogger) Panic

```go
func (l *BasicLogger) Panic(args ...interface{})
```
Panic outputs a log message if logLevel permits, and then panics

#### func (*BasicLogger) PanicOnError

```go
func (l *BasicLogger) PanicOnError(err error)
```
PanicOnError does nothing if err is nil; otherwise outputs a log message if
logLevel permits, and then panics

#### func (*BasicLogger) Panicf

```go
func (l *BasicLogger) Panicf(f string, args ...interface{})
```
Panicf outputs a formatted log message if logLevel permits, and then panics

#### func (*BasicLogger) Prefix

```go
func (l *BasicLogger) Prefix() string
```
Prefix returns the Logger's prefix string (does not include ": " trailer)

#### func (*BasicLogger) Print

```go
func (l *BasicLogger) Print(args ...interface{})
```
Print outputs to a Logger

#### func (*BasicLogger) Printf

```go
func (l *BasicLogger) Printf(f string, args ...interface{})
```
Printf outputs to a Logger

#### func (*BasicLogger) SetLogLevel

```go
func (l *BasicLogger) SetLogLevel(logLevel LogLevel)
```
SetLogLevel sets the log level

#### func (*BasicLogger) Sprint

```go
func (l *BasicLogger) Sprint(args ...interface{}) string
```
Sprint returns a string that has the Logger's prefix

#### func (*BasicLogger) Sprintf

```go
func (l *BasicLogger) Sprintf(f string, args ...interface{}) string
```
Sprintf returns a string that has the Logger's prefix

#### func (*BasicLogger) TLog

```go
func (l *BasicLogger) TLog(args ...interface{})
```
TLog outputs a formatted log message if logLevel permits

#### func (*BasicLogger) TLogError

```go
func (l *BasicLogger) TLogError(args ...interface{}) error
```
TLogError outputs an error message to a Logger iff logging level is enabled, and
returns an error object with a description string that has the logger's prefix

#### func (*BasicLogger) TLogErrorf

```go
func (l *BasicLogger) TLogErrorf(f string, args ...interface{}) error
```
TLogErrorf outputs an error message to a Logger iff logging level is enabled,
and returns an error object with a description string that has the logger's
prefix

#### func (*BasicLogger) TLogf

```go
func (l *BasicLogger) TLogf(f string, args ...interface{})
```
TLogf outputs a formatted log message if logLevel permits

#### func (*BasicLogger) WLog

```go
func (l *BasicLogger) WLog(args ...interface{})
```
WLog outputs a formatted log message if logLevel permits

#### func (*BasicLogger) WLogError

```go
func (l *BasicLogger) WLogError(args ...interface{}) error
```
WLogError outputs an error message to a Logger iff logging level is enabled, and
returns an error object with a description string that has the logger's prefix

#### func (*BasicLogger) WLogErrorf

```go
func (l *BasicLogger) WLogErrorf(f string, args ...interface{}) error
```
WLogErrorf outputs an error message to a Logger iff logging level is enabled,
and returns an error object with a description string that has the logger's
prefix

#### func (*BasicLogger) WLogf

```go
func (l *BasicLogger) WLogf(f string, args ...interface{})
```
WLogf outputs a formatted log message if logLevel permits

#### type FlagsLogger

```go
type FlagsLogger interface {
	Flags() int
}
```

FlagsLogger is an interface for a logger that supports Flags() api

#### type GetLogLeveler

```go
type GetLogLeveler interface {
	GetLogLevel() LogLevel
}
```

GetLogLeveler is An interface for a logger that supports GetLogLevel()

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
StringToLogLevel converts a string to a LogLevel

#### func (*LogLevel) FromString

```go
func (x *LogLevel) FromString(s string) error
```
FromString initiales a LogLevel from a string

#### func (*LogLevel) String

```go
func (x *LogLevel) String() string
```

#### type Logger

```go
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

	// Fork creates a new Logger that has an additional formatted string appended onto
	// an existing logger's prefix (with ": " added between)
	Fork(prefix string, args ...interface{}) Logger

	SetLogLevel(logLevel LogLevel)
}
```

Logger is an interface for a logging component that supports logging levels and
prefix forking

#### func  NewLogWrapper

```go
func NewLogWrapper(logger MinLogger, prefix string, logLevel LogLevel) Logger
```
NewLogWrapper creates a new Logger that wraps an existing MinLogger

#### func  NewLogger

```go
func NewLogger(prefix string, logLevel LogLevel) Logger
```
NewLogger creates a new Logger with a given prefix and Default flags, emitting
output to os.Stderr

#### func  NewLoggerWithFlags

```go
func NewLoggerWithFlags(prefix string, flag int, logLevel LogLevel) Logger
```
NewLoggerWithFlags creates a new Logger with a given prefix flags, emitting
output to os.Stderr

#### type MinLogger

```go
type MinLogger interface {
	Print(args ...interface{})
	Prefix() string
}
```

MinLogger is a minimal logging interface for a logging component
<!--/tmpl-->

### Contributing

- http://golang.org/doc/code.html
- http://golang.org/doc/effective_go.html
- `github.com/sammck-go/wstunnel/share` contains the shared package
- `github.com/sammck-go/wstunnel/server` contains the server package
- `github.com/sammck-go/wstunnel/client` contains the client package

### Changelog

- `1.0` - Initial release.

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
