package log

import (
	"fmt"
	"log/slog"
	"os"
)

const (
	LevelDebug = "debug"
	LevelInfo  = "info"
	LevelWarn  = "warn"
	LevelError = "error"
)

func init() {
	envLogLevel := os.Getenv("LOG_LEVEL")
	var slogLevel slog.Level
	switch envLogLevel {
	case LevelError:
		slogLevel = slog.LevelError
	case LevelWarn:
		slogLevel = slog.LevelWarn
	case LevelInfo:
		slogLevel = slog.LevelInfo
	case LevelDebug:
		slogLevel = slog.LevelDebug
	default:
		slogLevel = slog.LevelInfo
	}

	// 创建一个自定义的 ReplaceAttr 函数修改时间格式
	replaceAttr := func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey {
			return slog.Attr{
				Key:   slog.TimeKey,
				Value: slog.StringValue(a.Value.Time().Format("2006-01-02 15:04:05")),
			}
		}
		return a
	}
	textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:       slogLevel,
		ReplaceAttr: replaceAttr,
	})
	slog.SetDefault(slog.New(textHandler))
	fmt.Println("==> log level(slog): " + slogLevel.Level().String())
}

func standardizeArgs(level string, args ...any) []any {
	if len(args)%2 != 0 {
		return append([]any{level}, args...)
	}
	return args
}

func Debug(msg string, args ...any) {
	args = standardizeArgs(LevelDebug, args...)
	slog.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	args = standardizeArgs(LevelInfo, args...)
	slog.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	args = standardizeArgs(LevelWarn, args...)
	slog.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	args = standardizeArgs(LevelError, args...)
	slog.Error(msg, args...)
}

func Fatal(msg string, args ...any) {
	Error("fatal error: "+msg, args...)
	os.Exit(1)
}

type LogValueWrapper struct {
	GenerateValue func() string
}

func (l LogValueWrapper) LogValue() slog.Value {
	return slog.AnyValue(l.GenerateValue())
}
