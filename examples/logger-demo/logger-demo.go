package main

import (
	"fmt"
	"os"

	"github.com/sammck-go/logger"
	flag "github.com/spf13/pflag"
	"github.com/thediveo/enumflag"
)

type testObj struct {
	logger.Logger
	counter int
}

func newTestObj(logger logger.Logger, id int) *testObj {
	t := &testObj{
		Logger:  logger.ForkLogf("TestObj %d", id),
		counter: 0,
	}
	return t
}

const (
	nto     = 3
	niter   = 10
	dateLen = 10
	timeLen = 8
)

// LogLevelVarP adds a custom pflag that can be set to one of the known LogLevels by name
func LogLevelVarP(logLevel *logger.LogLevel, name string, shorthand string, value logger.LogLevel, usage string) {
	*logLevel = value
	logLevelOpts := map[logger.LogLevel][]string{}
	for ll, s := range logger.LogLevelToName {
		if s != "unknown" {
			logLevelOpts[ll] = []string{s}
		}
	}
	flag.VarP(
		enumflag.New(logLevel, "LogLevel", logLevelOpts, enumflag.EnumCaseInsensitive),
		name,
		shorthand,
		usage,
	)
}

func run() error {

	noDate := false
	noTime := false
	withLongFile := false
	withShortFile := false
	withMicroseconds := false
	withUTC := false
	logLevel := logger.LogLevelUnknown
	flag.BoolVarP(&noDate, "no-date", "", false, "Do not include a date in the timestamp.")
	flag.BoolVarP(&noTime, "no-time", "", false, "Do not include a time of date in the timestamp.")
	flag.BoolVarP(&withLongFile, "with-long-file", "", false, "include long filename and line number.")
	flag.BoolVarP(&withShortFile, "with-short-file", "", false, "include short filename and line number.")
	flag.BoolVarP(&withMicroseconds, "microseconds", "", false, "Display time with microsecond resolution.")
	flag.BoolVarP(&withUTC, "utc", "u", false, "Display date/time as UTC.")
	LogLevelVarP(&logLevel, "loglevel", "", logger.LogLevelDebug, "Set the log level.")

	flag.Parse()

	cfg := logger.NewConfig(logger.WithLogLevel(logLevel))
	if withShortFile {
		cfg = cfg.Refine(logger.WithLshortfile())
	}
	if withLongFile {
		cfg = cfg.Refine(logger.WithLlongfile())
	}
	if withUTC {
		cfg = cfg.Refine(logger.WithLUTC())
	}
	if withMicroseconds {
		cfg = cfg.Refine(logger.WithLmicroseconds())
	}
	if noDate {
		cfg = cfg.Refine(logger.WithoutLdate())
	}
	if noTime {
		cfg = cfg.Refine(logger.WithoutLtime())
	}

	lg, err := logger.New(logger.WithConfig(cfg))

	if err != nil {
		return fmt.Errorf("logger.New() returned error: %s", err)
	}

	lg.DLogf("Top level log entry")

	tos := []*testObj{}

	for i := 0; i < nto; i++ {
		to := newTestObj(lg, i)
		tos = append(tos, to)
	}

	to := newTestObj(tos[0], 10)
	tos = append(tos, to)

	for i := 0; i < niter; i++ {
		for j := 0; j < len(tos); j++ {
			tos[j].DLogf("Log Message %d", i)
		}
	}

	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "logger-demo: %s\n", err)
		os.Exit(1)
	}
}
