package zlog

import (
	"os"
	"sync"

	"github.com/fatih/color"
)

const (
	head int = iota

	trace
	debug
	info
	warn
	err
	fatal

	max
)

var (
	once sync.Once
	line [max]struct {
		col   *color.Color
		space []interface{}
	}
)

func set_line(i int, col *color.Color, space []interface{}) {
	line[i].col = col
	line[i].space = space
}

func initialize() {
	set_line(head, color.New(color.FgHiBlue), []interface{}{""})
	set_line(trace, color.New(color.FgHiBlack), []interface{}{"[TRACE]"})
	set_line(debug, color.New(color.FgCyan), []interface{}{"[DEBUG]"})
	set_line(info, color.New(color.FgHiGreen), []interface{}{"[INFO]"})
	set_line(warn, color.New(color.FgHiYellow), []interface{}{"[WARN]"})
	set_line(err, color.New(color.FgHiMagenta), []interface{}{"[ERROR]"})
	set_line(fatal, color.New(color.FgHiRed), []interface{}{"[FATAL]"})
}

func Head(a ...interface{}) {
	once.Do(initialize)

	a = append(line[head].space, a...)
	line[head].col.Println(a...)
}

func Trace(a ...interface{}) {
	once.Do(initialize)

	a = append(line[trace].space, a...)
	line[trace].col.Println(a...)
}

func Debug(a ...interface{}) {
	once.Do(initialize)

	a = append(line[debug].space, a...)
	line[debug].col.Println(a...)
}

func Info(a ...interface{}) {
	once.Do(initialize)

	a = append(line[info].space, a...)
	line[info].col.Println(a...)
}

func Warn(a ...interface{}) {
	once.Do(initialize)

	a = append(line[warn].space, a...)
	line[warn].col.Println(a...)
}

func Error(a ...interface{}) {
	once.Do(initialize)

	a = append(line[err].space, a...)
	line[err].col.Println(a...)
}

func Fatal(a ...interface{}) {
	once.Do(initialize)

	a = append(line[fatal].space, a...)
	line[fatal].col.Println(a...)

	os.Exit(1)
}

func NoError(err error, a ...interface{}) {
	if err != nil {
		if len(a) > 0 {
			Fatal(append(a, []interface{}{", error:", err}...)...)
		} else {
			Fatal([]interface{}{err}...)
		}
	}
}

func MustTrue(ok bool, a ...interface{}) {
	if !ok {
		if len(a) > 0 {
			Fatal(append(a, []interface{}{", expected true:", ok}...)...)
		} else {
			Fatal([]interface{}{"expected true"}...)
		}
	}
}
