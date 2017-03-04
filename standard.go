package golog

var (
	std = NewLogger()
)

func SetOutputAndRotate(rotate LogRotate) {
	std.mu.Lock()
	defer std.mu.Unlock()
	std.SetOutPutAndRotate(rotate)
}

func SetLogLevel(level LogLevel) {
	std.mu.Lock()
	defer std.mu.Unlock()
	std.SetLogLevel(level)
}

func SetFormatter(fm string) {
	std.mu.Lock()
	defer std.mu.Unlock()
	std.SetFormatter(fm)
}

func GetLogLevel() LogLevel {
	return std.GetLogLevel()
}

func Debug(args ...interface{}) {
	std.Debug(args...)
}

func Print(args ...interface{}) {
	std.Info(args...)
}

func Info(args ...interface{}) {
	std.Info(args...)
}

func Warn(args ...interface{}) {
	std.Warn(args...)
}

func Error(args ...interface{}) {
	std.Error(args...)
}

func Fatal(args ...interface{}) {
	std.Fatal(args...)
}

func Panic(args ...interface{}) {
	std.Panic(args...)
}

func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}

func Printf(format string, args ...interface{}) {
	std.Infof(format, args...)
}

func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	std.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	std.Panicf(format, args...)
}

func Debugln(args ...interface{}) {
	std.Debugln(args...)
}

func Println(args ...interface{}) {
	std.Infoln(args...)
}

func Infoln(args ...interface{}) {
	std.Infoln(args...)
}

func Warnln(args ...interface{}) {
	std.Warnln(args...)
}

func Errorln(args ...interface{}) {
	std.Errorln(args...)
}

func Fatalln(args ...interface{}) {
	std.Fatalln(args...)
}

func Panicln(args ...interface{}) {
	std.Panicln(args...)
}
