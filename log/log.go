package log

import (
	"fmt"
	"io"
	"os"
	"time"
	"runtime/debug"
)

// Level defines current log level
type Level int32

const (
	LFatal Level = iota
	LError
	LWarning
	LInfo
	LDebug
)

var levelPrefix = []string{
	LFatal:   "FATA",
	LError:   "ERRO",
	LWarning: "WARN",
	LInfo:    "INFO",
	LDebug:   "DEBU",
}

var logLevel = LInfo
var out io.Writer = os.Stdout
var buf []byte

// Setup log level and output writer
func Setup(l Level, w io.Writer) {
	logLevel = l
	out = w
}

func itoa(b *[]byte, i int, wid int) {
	// Assemble decimal in reverse order.
	var tmp [20]byte
	bp := len(tmp) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		tmp[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	tmp[bp] = byte('0' + i)
	*b = append(*b, tmp[bp:]...)
}

func formatHeader(b *[]byte, l Level, t time.Time) {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()

	itoa(b, year, 4)
	*b = append(*b, '/')
	itoa(b, int(month), 2)
	*b = append(*b, '/')
	itoa(b, day, 2)
	*b = append(*b, ' ')

	itoa(b, hour, 2)
	*b = append(*b, ':')
	itoa(b, min, 2)
	*b = append(*b, ':')
	itoa(b, sec, 2)

	*b = append(*b, ' ', '[')
	*b = append(*b, levelPrefix[l]...)
	*b = append(*b, ']', ' ')
}

func output(l Level, s string) {
	now := time.Now()
	buf = buf[:0]
	formatHeader(&buf, l, now)
	buf = append(buf, s...)
	if len(s) == 0 || s[len(s)-1] != '\n' {
		buf = append(buf, '\n')
	}
	out.Write(buf)
}

func outputI(l Level, v ...interface{}) {
	if l <= logLevel {
		output(l, fmt.Sprint(v...))
	}
}

func outputF(l Level, f string, v ...interface{}) {
	if l <= logLevel {
		output(l, fmt.Sprintf(f, v...))
	}
}

// Fatal message and exit 1
func Fatal(v ...interface{}) {
	outputI(LFatal, v...)
	debug.PrintStack()
	os.Exit(1)
}

// FatalF ouptputs formatted message and exit 1
func FatalF(f string, v ...interface{}) {
	outputF(LFatal, f, v...)
	debug.PrintStack()
	os.Exit(1)
}

// Error message
func Error(v ...interface{}) {
	outputI(LError, v...)
}

// ErrorF outputs formatted message
func ErrorF(f string, v ...interface{}) {
	outputF(LError, f, v...)
}

// Warning message
func Warning(v ...interface{}) {
	outputI(LWarning, v...)
}

// WarningF outputs formatted message
func WarningF(f string, v ...interface{}) {
	outputF(LWarning, f, v...)
}

// Info message
func Info(v ...interface{}) {
	outputI(LInfo, v...)
}

// InfoF outputs formatted message
func InfoF(f string, v ...interface{}) {
	outputF(LInfo, f, v...)
}

// Debug message
func Debug(v ...interface{}) {
	outputI(LDebug, v...)
}

// DebugF outputs formatted message
func DebugF(f string, v ...interface{}) {
	outputF(LDebug, f, v...)
}
