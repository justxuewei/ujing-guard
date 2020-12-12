package util

import (
	"fmt"
	"github.com/fatih/color"
	"sync"
	"time"
)

// Log Level
const (
	Error = iota
	Warning
	Info
	Debug
)

var GlobalLogger *Logger
var Level = Debug

type Logger struct {
	level int
	mu    sync.Mutex
}

var colors = map[string]func(a ...interface{}) string{
	"Warning": color.New(color.FgYellow).Add(color.Bold).SprintFunc(),
	"Panic":   color.New(color.BgRed).Add(color.Bold).SprintFunc(),
	"Error":   color.New(color.FgRed).Add(color.Bold).SprintFunc(),
	"Info":    color.New(color.FgCyan).Add(color.Bold).SprintFunc(),
	"Debug":   color.New(color.FgWhite).Add(color.Bold).SprintFunc(),
}

// keep the width equivalent
var spaces = map[string]string{
	"Warning": "",
	"Panic":   "  ",
	"Error":   "  ",
	"Info":    "   ",
	"Debug":   "  ",
}

func (l *Logger) Println(prefix string, msg string) {
	c := color.New()

	l.mu.Lock()
	defer l.mu.Unlock()

	_, _ = c.Printf("%s%s %s %s\n", colors[prefix]("["+prefix+"]"),
		spaces[prefix],
		time.Now().Format("\"2006-01-02 15:04:05\""),
		msg,
	)
}

func (l *Logger) Panic(format string, v ...interface{}) {
	if Error > l.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	l.Println("Panic", msg)
	panic(msg)
}

// Error 错误
func (l *Logger) Error(format string, v ...interface{}) {
	if Error > l.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	l.Println("Error", msg)
}

// Warning 警告
func (l *Logger) Warning(format string, v ...interface{}) {
	if Warning > l.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	l.Println("Warning", msg)
}

// Info 信息
func (l *Logger) Info(format string, v ...interface{}) {
	if Info > l.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	l.Println("Info", msg)
}

// Debug 校验
func (l *Logger) Debug(format string, v ...interface{}) {
	if Debug > l.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	l.Println("Debug", msg)
}

func BuildLogger(level string) {
	intLevel := Error
	switch level {
	case "error":
		intLevel = Error
	case "warning":
		intLevel = Warning
	case "info":
		intLevel = Info
	case "debug":
		intLevel = Debug
	}
	l := Logger{
		level: intLevel,
	}
	GlobalLogger = &l
}

func Log() *Logger {
	if GlobalLogger == nil {
		l := Logger{
			level: Level,
		}
		GlobalLogger = &l
	}
	return GlobalLogger
}
