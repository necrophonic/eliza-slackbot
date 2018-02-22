package log

import (
	"log"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	log.SetOutput(&OutWriter{})
	os.Exit(m.Run())
}

var testOutput string

func TestInitFromString(t *testing.T) {
	if err := InitFromString(""); err.Error() != "unable to init logger: unknown level ''" {
		t.Error("expected error when initialising with empty string but got none")
	}

	if err := InitFromString("unknown"); err.Error() != "unable to init logger: unknown level 'unknown'" {
		t.Errorf("expected error when initialising with 'unknown' but got: %s", err.Error())
	}

	if InitFromString("error"); level != LevelError {
		t.Error("failed to initialise to level 'error'")
	}

	if InitFromString("info"); level != LevelInfo {
		t.Error("failed to initialise to level 'info'")
	}

	if InitFromString("warn"); level != LevelWarn {
		t.Error("failed to initialise to level 'warn'")
	}

	if InitFromString("debug"); level != LevelDebug {
		t.Error("failed to initialise to level 'debug'")
	}

	if InitFromString("trace"); level != LevelTrace {
		t.Error("failed to initialise to level 'trace'")
	}

	if InitFromString("none"); level != LevelNone {
		t.Error("failed to initialise to level 'none'")
	}
}

func TestInit(t *testing.T) {
	for _, l := range []int{-1, 6} {
		if err := Init(l); err == nil {
			t.Errorf("expected error when level was '%d' but got none", l)
		}
	}
}

func TestTrace(t *testing.T) {
	testCases := map[int][]string{
		LevelNone:  []string{"Trace me %s", "", ""},
		LevelError: []string{"Trace me %s", "", ""},
		LevelInfo:  []string{"Trace me %s", "", ""},
		LevelWarn:  []string{"Trace me %s", "", ""},
		LevelDebug: []string{"Trace me %s", "", ""},
		LevelTrace: []string{"Trace me %s", "Trace me abc", "abc"},
	}
	testLogfMethod(t, func(f string, s string) { Tracef(f, s) }, testCases)
	testLogMethod(t, func(s string) { Trace(s) }, testCases)
}

func TestDebug(t *testing.T) {
	testCases := map[int][]string{
		LevelNone:  []string{"Debug me %s", "", ""},
		LevelError: []string{"Debug me %s", "", ""},
		LevelInfo:  []string{"Debug me %s", "", ""},
		LevelWarn:  []string{"Debug me %s", "", ""},
		LevelDebug: []string{"Debug me %s", "Debug me abc", "abc"},
		LevelTrace: []string{"Debug me %s", "Debug me abc", "abc"},
	}
	testLogfMethod(t, func(f string, s string) { Debugf(f, s) }, testCases)
	testLogMethod(t, func(s string) { Debug(s) }, testCases)
}

func TestWarn(t *testing.T) {
	testCases := map[int][]string{
		LevelNone:  []string{"Warn me %s", "", ""},
		LevelError: []string{"Warn me %s", "", ""},
		LevelInfo:  []string{"Warn me %s", "", ""},
		LevelWarn:  []string{"Warn me %s", "Warn me abc", "abc"},
		LevelTrace: []string{"Warn me %s", "Warn me abc", "abc"},
		LevelDebug: []string{"Warn me %s", "Warn me abc", "abc"},
	}
	testLogfMethod(t, func(f string, s string) { Warnf(f, s) }, testCases)
	testLogMethod(t, func(s string) { Warn(s) }, testCases)
}

func TestInfo(t *testing.T) {
	testCases := map[int][]string{
		LevelNone:  []string{"Info me %s", "", ""},
		LevelError: []string{"Info me %s", "", ""},
		LevelInfo:  []string{"Info me %s", "Info me abc", "abc"},
		LevelWarn:  []string{"Info me %s", "Info me abc", "abc"},
		LevelDebug: []string{"Info me %s", "Info me abc", "abc"},
		LevelTrace: []string{"Info me %s", "Info me abc", "abc"},
	}
	testLogfMethod(t, func(f string, s string) { Infof(f, s) }, testCases)
	testLogMethod(t, func(s string) { Info(s) }, testCases)
}

func TestError(t *testing.T) {
	testCases := map[int][]string{
		LevelNone:  []string{"Error me %s", "", ""},
		LevelError: []string{"Error me %s", "Error me abc", "abc"},
		LevelInfo:  []string{"Error me %s", "Error me abc", "abc"},
		LevelWarn:  []string{"Error me %s", "Error me abc", "abc"},
		LevelDebug: []string{"Error me %s", "Error me abc", "abc"},
		LevelTrace: []string{"Error me %s", "Error me abc", "abc"},
	}
	testLogfMethod(t, func(f string, s string) { Errorf(f, s) }, testCases)
	testLogMethod(t, func(s string) { Error(s) }, testCases)
}

func testLogfMethod(t *testing.T, f func(string, string), cases map[int][]string) {
	for lev, v := range cases {
		clear()
		Init(lev)
		f(v[0], v[2])
		if !strings.Contains(testOutput, v[1]) {
			t.Errorf("at level '%d', expected output '%s' but got '%s'", lev, v[1], testOutput)
		}
	}
}

func testLogMethod(t *testing.T, f func(s string), cases map[int][]string) {
	for lev, v := range cases {
		clear()
		Init(lev)
		f(v[1])
		if !strings.Contains(testOutput, v[1]) {
			t.Errorf("at level '%d', expected output '%s' but got '%s'", lev, v[1], testOutput)
		}
	}
}

// OutWriter implements the io.Writer interface. Used to capture test output
type OutWriter struct{}

func (o *OutWriter) Write(p []byte) (int, error) {
	testOutput = string(p)
	return 0, nil
}

func clear() {
	testOutput = ""
}
