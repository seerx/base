package logs

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

// Log 默认日志
var Log *logrus.Logger

// InitLog 初始化默认日志
func InitLog(b *Builder) {
	if b == nil {
		Log = NewBuilder().
			Build()
	}
	Log = b.Build()
}

// WithField Log.WithField
func WithField(key string, value interface{}) *logrus.Entry {
	return Log.WithField(key, value)
}

// WithFields Log.WithFields
func WithFields(fields logrus.Fields) *logrus.Entry {
	return Log.WithFields(fields)
}

// WithError Log.WithError
func WithError(err error) *logrus.Entry {
	return Log.WithError(err)
}

// WithContext Add a context to the log entry.
func WithContext(ctx context.Context) *logrus.Entry {
	return Log.WithContext(ctx)
}

// WithTime Overrides the time of the log entry.
func WithTime(t time.Time) *logrus.Entry {
	return Log.WithTime(t)
}

// Logf Log.Logf
func Logf(level logrus.Level, format string, args ...interface{}) {
	Log.Logf(level, format, args...)
}

// Tracef Log.Tracef
func Tracef(format string, args ...interface{}) {
	Log.Tracef(format, args...)
}

// Debugf Log.Debugf
func Debugf(format string, args ...interface{}) {
	Log.Debugf(format, args...)
}

// Infof Log.Infof
func Infof(format string, args ...interface{}) {
	Log.Infof(format, args...)
}

// Printf Log.Printf
func Printf(format string, args ...interface{}) {
	Log.Printf(format, args...)
}

// Warnf Log.Warnf
func Warnf(format string, args ...interface{}) {
	Log.Warnf(format, args...)
}

// Warningf Log.Warningf
func Warningf(format string, args ...interface{}) {
	Log.Warningf(format, args...)
}

// Errorf Log.Errorf
func Errorf(format string, args ...interface{}) {
	Log.Errorf(format, args...)
}

// Fatalf Log.Fatalf
func Fatalf(format string, args ...interface{}) {
	Log.Fatalf(format, args...)
}

// Panicf Log.Panicf
func Panicf(format string, args ...interface{}) {
	Log.Panicf(format, args...)
}

// LogLog Log.Log
func LogLog(level logrus.Level, args ...interface{}) {
	Log.Log(level, args...)
}

// Trace Log.Trace
func Trace(args ...interface{}) {
	Log.Trace(args...)
}

// Debug Log.Debug
func Debug(args ...interface{}) {
	Log.Debug(args...)
}

// Info Log.Info
func Info(args ...interface{}) {
	Log.Info(args...)
}

// Print Log.Print
func Print(args ...interface{}) {
	Log.Print(args...)
}

// Warn Log.Warn
func Warn(args ...interface{}) {
	Log.Warn(args...)
}

// Warning Log.Warning
func Warning(args ...interface{}) {
	Log.Warning(args...)
}

// Error Log.Error
func Error(args ...interface{}) {
	Log.Error(args...)
}

// Fatal Log.Fatal
func Fatal(args ...interface{}) {
	Log.Fatal(args...)
}

// Panic Log.Panic
func Panic(args ...interface{}) {
	Log.Panic(args...)
}

// Logln Log.Logln
func Logln(level logrus.Level, args ...interface{}) {
	Log.Logln(level, args...)
}

// Traceln Log.Traceln
func Traceln(args ...interface{}) {
	Log.Traceln(args...)
}

// Debugln Log.Debugln
func Debugln(args ...interface{}) {
	Log.Debugln(args...)
}

// Infoln Log.Infoln
func Infoln(args ...interface{}) {
	Log.Infoln(args...)
}

// Println Log.Println
func Println(args ...interface{}) {
	Log.Println(args...)
}

// Warnln Log.Warnln
func Warnln(args ...interface{}) {
	Log.Warnln(args...)
}

// Warningln Log.Warningln
func Warningln(args ...interface{}) {
	Log.Warningln(args...)
}

// Errorln Log.Errorln
func Errorln(args ...interface{}) {
	Log.Errorln(args...)
}

// Fatalln Log.Fatalln
func Fatalln(args ...interface{}) {
	Log.Fatalln(args...)
}

// Panicln Log.Panicln
func Panicln(args ...interface{}) {
	Log.Panicln(args...)
}

// Exit Log.Exit
// func Exit(code int) {
// 	Log.Exit(code)
// }

//SetNoLock When file is opened with appending mode, it's safe to
//write concurrently to a file (within 4k message on Linux).
//In these cases user can choose to disable the lock.
func SetNoLock() {
	Log.SetNoLock()
}

// SetLevel sets the logger level.
// func SetLevel(level logrus.Level) {
// 	Log.SetLevel(level)
// }

// GetLevel returns the logger level.
// func GetLevel() logrus.Level {
// 	return Log.GetLevel()
// }

// AddHook adds a hook to the logger hooks.
// func AddHook(hook logrus.Hook) {
// 	Log.AddHook(hook)
// }

// IsLevelEnabled checks if the log level of the logger is greater than the level param
// func IsLevelEnabled(level logrus.Level) bool {
// 	return Log.IsLevelEnabled(level)
// }

// SetFormatter sets the logger formatter.
// func SetFormatter(formatter logrus.Formatter) {
// 	Log.SetFormatter(formatter)
// }

// SetOutput sets the logger output.
// func SetOutput(output io.Writer) {
// 	Log.SetOutput(output)
// }

// SetReportCaller Log.SetReportCaller
// func SetReportCaller(reportCaller bool) {
// 	Log.SetReportCaller(reportCaller)
// }

// ReplaceHooks replaces the logger hooks and returns the old ones
// func ReplaceHooks(hooks logrus.LevelHooks) logrus.LevelHooks {
// 	return Log.ReplaceHooks(hooks)
// }
