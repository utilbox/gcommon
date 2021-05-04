// Package log wraps zap and provide an easy way to use zap to log to local file.
// For more details about zap, see https://github.com/uber-go/zap .
package log

import (
	"fmt"
	"os"
	"runtime"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var sugarLogger *zap.SugaredLogger = zap.NewNop().Sugar()

var logLevels = map[string]zapcore.Level{
	"Debug":  zapcore.DebugLevel,
	"Info":   zapcore.InfoLevel,
	"Warn":   zapcore.WarnLevel,
	"Error":  zapcore.ErrorLevel,
	"DPanic": zapcore.DPanicLevel,
	"Panic":  zapcore.PanicLevel,
	"Fatal":  zapcore.FatalLevel,
}

func InitLogger(c *LogConfig) {
	if c == nil {
		fmt.Println("invalid log config")
		os.Exit(-1)
	}
	c.SetDefault()
	writeSyncer := getLogWriter(c)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, logLevels[c.Level])

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(c.Depth))

	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	//encoderConfig := zapcore.EncoderConfig{
	//	TimeKey:        "time",
	//	LevelKey:       "level",
	//	NameKey:        "logger",
	//	CallerKey:      "linenum",
	//	MessageKey:     "msg",
	//	StacktraceKey:  "stacktrace",
	//	LineEnding:     zapcore.DefaultLineEnding,
	//	EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
	//	EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
	//	EncodeDuration: zapcore.SecondsDurationEncoder, //
	//	EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
	//	EncodeName:     zapcore.FullNameEncoder,
	//}
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = LineAndMethodCallerEndocer
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// LineAndMethodCallerEndocer serializes a caller in package/file:line package:method format.
func LineAndMethodCallerEndocer(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
	if !caller.Defined {
		encoder.AppendString(caller.TrimmedPath())
		return
	}
	method := runtime.FuncForPC(caller.PC).Name()
	encoder.AppendString(fmt.Sprintf("%s %s", caller.TrimmedPath(), method))
	return
}

func getLogWriter(c *LogConfig) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   c.Path,
		MaxSize:    c.MaxSize,
		MaxBackups: c.MaxBackups,
		MaxAge:     c.MaxAge,
		Compress:   c.Compress,
		LocalTime:  c.LocalTime,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// Debug uses fmt.Sprint to construct and log a message.
func Debug(args ...interface{}) {
	sugarLogger.Debug(args...)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	sugarLogger.Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	sugarLogger.Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	sugarLogger.Error(args...)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanic(args ...interface{}) {
	sugarLogger.DPanic(args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func Panic(args ...interface{}) {
	sugarLogger.Panic(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...interface{}) {
	sugarLogger.Fatal(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	sugarLogger.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	sugarLogger.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	sugarLogger.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	sugarLogger.Errorf(template, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanicf(template string, args ...interface{}) {
	sugarLogger.DPanicf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(template string, args ...interface{}) {
	sugarLogger.Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	sugarLogger.Fatalf(template, args...)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//  s.With(keysAndValues).Debug(msg)
func Debugw(msg string, keysAndValues ...interface{}) {
	sugarLogger.Debugw(msg, keysAndValues...)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Infow(msg string, keysAndValues ...interface{}) {
	sugarLogger.Infow(msg, keysAndValues...)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Warnw(msg string, keysAndValues ...interface{}) {
	sugarLogger.Warnw(msg, keysAndValues...)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Errorw(msg string, keysAndValues ...interface{}) {
	sugarLogger.Errorw(msg, keysAndValues...)
}

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func DPanicw(msg string, keysAndValues ...interface{}) {
	sugarLogger.DPanicw(msg, keysAndValues...)
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func Panicw(msg string, keysAndValues ...interface{}) {
	sugarLogger.Panicw(msg, keysAndValues...)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func Fatalw(msg string, keysAndValues ...interface{}) {
	sugarLogger.Fatalw(msg, keysAndValues...)
}

// Sync flushes any buffered log entries.
func Sync() error {
	return sugarLogger.Sync()
}
