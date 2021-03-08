package logger

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

type TestObj struct {
	Logger

	counter int
}

func NewTestObj(logger Logger, id int) *TestObj {
	t := &TestObj{
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

func TestLogging(t *testing.T) {

	d := t.TempDir()

	lfname := filepath.Join(d, "test.log")

	lf, err := os.Create(lfname)
	if err != nil {
		t.Fatalf("os.Create(\"%s\") returned error: %s", lfname, err)
	}

	cleanup := func() {
		if lf != nil {
			lf.Close()
		}
	}
	defer cleanup()

	lg, err := New(
		WithWriter(lf),
		WithLogLevel(LogLevelDebug),
		WithPrefix("TestLogging"),
	)

	if err != nil {
		t.Fatalf("logger.New() returned error: %s", err)
	}

	expectedLines := []string{}

	lg.DLogf("Top level")
	expectedLines = append(expectedLines, fmt.Sprintf("TestLogging: Top level"))

	tos := []*TestObj{}

	for i := 0; i < nto; i++ {
		to := NewTestObj(lg, i)
		tos = append(tos, to)
	}

	for i := 0; i < niter; i++ {
		for j := 0; j < nto; j++ {
			tos[j].DLogf("Log Message %d", i)
			expectedLines = append(expectedLines, fmt.Sprintf("TestLogging: TestObj %d: Log Message %d", j, i))
		}
	}

	err = lf.Close()
	if err != nil {
		t.Fatalf("lf.Close() of write file returned error: %s", err)
	}

	lf, err = os.Open(lfname)
	if err != nil {
		t.Fatalf("os.Open(\"%s\") returned error: %s", lfname, err)
	}

	scanner := bufio.NewScanner(lf)

	for scanner.Scan() {
		line := scanner.Text()
		if len(expectedLines) == 0 {
			t.Errorf("scanner.Scan() found extraneous line \"%s\" at expected EOF", line)
		} else {
			expectedLine := expectedLines[0]
			expectedLines = expectedLines[1:]
			if len(line) < dateLen+1+timeLen+1 {
				t.Errorf("Log line too short to contain date and timestamp: [%s]", line)
			} else {
				lineTail := line[dateLen+1+timeLen+1:]
				if lineTail != expectedLine {
					t.Errorf("Expected log line [<date> <time> %s]; got [%s]", expectedLine, line)
				}
			}
		}
	}

	err = scanner.Err()
	if err != nil {
		t.Errorf("scanner.Err() returned error: %s", err)
	}

	for len(expectedLines) > 0 {
		expectedLine := expectedLines[0]
		expectedLines = expectedLines[1:]
		t.Errorf("scanner.Scan() got EOF before expected line \"%s\"", expectedLine)
	}

	err = lf.Close()
	if err != nil {
		t.Fatalf("lf.Close() of read file returned error: %s", err)
	}
	lf = nil
}
