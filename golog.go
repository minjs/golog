package golog

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
)

type LogLevel uint8

const (
	PanicLevel LogLevel = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

type LogRotate struct {
	Filename  string
	MaxSize   int		 // megabytes
	MaxBackup int
	MaxAge    int		 //days
	DupToStd  bool
}

type Logger interface {
	SetLogLevel(level LogLevel)
	SetOutPut(out io.Writer)

	Debug(...interface{})
	Debugf(string, ...interface{})
	Debugln(...interface{})

	Info(...interface{})
	Infof(string, ...interface{})
	Infoln(...interface{})

	Warn(...interface{})
	Warnf(string, ...interface{})
	Warnln(...interface{})

	Error(...interface{})
	Errorf(string, ...interface{})
	Errorln(...interface{})

	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})

	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicln(...interface{})
}

type logger struct {
	logEntry *logrus.Entry
	mu       sync.Mutex
}

func NewLogger() logger {
	logrusEntry := logrus.NewEntry(logrus.New())
	l := logger{
		logEntry: logrusEntry,
	}
	return l
}

func (l logger) SetLogLevel(level LogLevel) {
	l.logEntry.Logger.Level = logrus.Level(level)
}

func (l logger) SetFormatter(fm string) {
	if fm == "json" {
		l.logEntry.Logger.Formatter = new(logrus.JSONFormatter)
		fmt.Println("setup log json format", fm)
	} else {
		customFormatter := new(logrus.TextFormatter)
		customFormatter.DisableColors = true
		l.logEntry.Logger.Formatter = customFormatter
		fmt.Println("setup log text format", fm)
	}
}

func (l logger) GetLogLevel() LogLevel {
	return LogLevel(l.logEntry.Logger.Level)
}

func (l logger) SetOutPutAndRotate(rotate LogRotate) {
	if rotate.DupToStd {
		l.logEntry.Logger.Out = io.MultiWriter(&lumberjack.Logger{
			Filename:   rotate.Filename,
			MaxSize:    rotate.MaxSize, // megabytes
			MaxBackups: rotate.MaxBackup,
			MaxAge:     rotate.MaxAge, //days
		}, os.Stdout)
	} else {
		l.logEntry.Logger.Out = &lumberjack.Logger{
			Filename:   rotate.Filename,
			MaxSize:    rotate.MaxSize, // megabytes
			MaxBackups: rotate.MaxBackup,
			MaxAge:     rotate.MaxAge, //days
		}
	}
}

func (l logger) sourced() *logrus.Entry {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "<???>"
		line = 0
	}
	slash := strings.LastIndex(file, "/")
	file = file[slash+1:]
	return l.logEntry.WithField("source", fmt.Sprintf("%s:%d", file, line))
}

func (l logger) Debug(args ...interface{}) {
	l.sourced().Debug(args...)
}

func (l logger) Warn(args ...interface{}) {
	l.sourced().Warn(args...)
}

func (l logger) Info(args ...interface{}) {
	l.sourced().Info(args...)
}

func (l logger) Error(args ...interface{}) {
	l.sourced().Error(args...)
}

func (l logger) Fatal(args ...interface{}) {
	l.sourced().Fatal(args...)
}

func (l logger) Panic(args ...interface{}) {
	l.sourced().Panic(args...)
}

func (l logger) Debugln(args ...interface{}) {
	l.sourced().Debugln(args...)
}

func (l logger) Infoln(args ...interface{}) {
	l.sourced().Infoln(args...)
}

func (l logger) Warnln(args ...interface{}) {
	l.sourced().Warnln(args...)
}

func (l logger) Errorln(args ...interface{}) {
	l.sourced().Errorln(args...)
}

func (l logger) Fatalln(args ...interface{}) {
	l.sourced().Fatalln(args...)
}

func (l logger) Panicln(args ...interface{}) {
	l.sourced().Panicln(args...)
}

func (l logger) Debugf(format string, args ...interface{}) {
	l.sourced().Debugf(format, args...)
}

func (l logger) Infof(format string, args ...interface{}) {
	l.sourced().Infof(format, args...)
}

func (l logger) Warnf(format string, args ...interface{}) {
	l.sourced().Warnf(format, args...)
}

func (l logger) Errorf(format string, args ...interface{}) {
	l.sourced().Errorf(format, args...)
}

func (l logger) Fatalf(format string, args ...interface{}) {
	l.sourced().Fatalf(format, args...)
}

func (l logger) Panicf(format string, args ...interface{}) {
	l.sourced().Panicf(format, args...)
}
