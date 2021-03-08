package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sammck-go/logger"
	flag "github.com/spf13/pflag"
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

func run() error {
	noDate := false
	noTime := false
	flag.BoolVarP(&noDate, "no-date", "", false, "Do not include a date in the timestamp.")
	flag.BoolVarP(&noTime, "no-time", "", false, "Do not include a time of date in the timestamp.")

	flag.Parse()

	cfg := logger.NewConfig(logger.WithLogLevel(logger.LogLevelDebug))
	if noDate {
		cfg.Refine(logger.WithoutLogFlags(log.Ldate))
	}
	if noTime {
		cfg.Refine(logger.WithoutLogFlags(log.Ltime))
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
