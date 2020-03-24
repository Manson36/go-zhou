package mylog

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type LogLevel int8

const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

type Logger struct {
	level LogLevel
}

func parseLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("不存在这个级别")
		return UNKNOWN, err
	}
}

func NewLogger(levelStr string) Logger {
	level, err := parseLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		level: level,
	}
}

func (l Logger) Debug(msg string) {
	if l.level > DEBUG {
		now := time.Now()
		fmt.Printf("[%s] [DEBUG] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}

func (l Logger) Info(msg string) {
	fmt.Println(msg)
}

func (l Logger) Warning(msg string) {
	fmt.Println(msg)
}

func (l Logger) Error(msg string) {
	fmt.Println(msg)
}

func (l Logger) Fatal(msg string) {
	fmt.Println(msg)
}
