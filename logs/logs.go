package logs

import (
	"context"
)

const (
	TRACEID = "TRACE-ID"
)

// CtxTraceKey 定义CtxTraceKey类型 防止context重名覆盖.
type CtxTraceKey string

var (
	logger     *Logger
	dataLogger *Logger
)

func New(opts ...Option) *Logger {
	var option options
	for _, opt := range opts {
		opt(&option)
	}
	if len(option.Writers) == 0 {
		option.Writers = append(option.Writers, OUTPUT_CONSOLE)
	}

	return NewZapLogger(option)
}

// Init 全局初始化函数 调用方保证全局调用一次.
func Init(opts ...Option) {
	var option options
	for _, opt := range opts {
		opt(&option)
	}
	if len(option.Writers) == 0 {
		option.Writers = append(option.Writers, OUTPUT_CONSOLE)
	}

	logger = NewZapLogger(option)
}

// InitData 数据Log全局初始化函数 调用方保证全局调用一次.
func InitData(opts ...Option) {
	var option options
	for _, opt := range opts {
		opt(&option)
	}
	if len(option.Writers) == 0 {
		option.Writers = append(option.Writers, OUTPUT_CONSOLE)
	}

	dataLogger = NewZapLogger(option)
}

// Debugt 打印Debug日志 通过context打印trace id.
func Debugt(ctx context.Context, msg string, kv ...any) {
	args := make([]any, 0, len(kv)+2)
	args = append(args, TRACEID, GetTraceID(ctx))
	logger.Debugw(msg, append(args, kv...)...)
}

// Infot 打印Info日志 通过context打印trace id.
func Infot(ctx context.Context, msg string, kv ...any) {
	args := make([]any, 0, len(kv)+2)
	args = append(args, TRACEID, GetTraceID(ctx))
	logger.Infow(msg, append(args, kv...)...)
}

// Warnt 打印Warn日志 通过context打印trace id.
func Warnt(ctx context.Context, msg string, kv ...any) {
	args := make([]any, 0, len(kv)+2)
	args = append(args, TRACEID, GetTraceID(ctx))
	logger.Warnw(msg, append(args, kv...)...)
}

// Errort 打印Error日志 通过context打印trace id.
func Errort(ctx context.Context, msg string, kv ...any) {
	args := make([]any, 0, len(kv)+2)
	args = append(args, TRACEID, GetTraceID(ctx))
	logger.Errorw(msg, append(args, kv...)...)
}

// Fatalt 打印Fatal日志 通过context打印trace id.
func Fatalt(ctx context.Context, msg string, kv ...any) {
	args := make([]any, 0, len(kv)+2)
	args = append(args, TRACEID, GetTraceID(ctx))
	logger.Fatalw(msg, append(args, kv...)...)
}

// Debug 打印Debug日志.
func Debug(v ...any) {
	logger.Debug(v...)
}

// Info 打印Info日志.
func Info(v ...any) {
	logger.Info(v...)
}

// Warn 打印Warn日志.
func Warn(v ...any) {
	logger.Warn(v...)
}

// Error 打印Error日志.
func Error(v ...any) {
	logger.Error(v...)
}

// Fatal 打印Fatal日志.
func Fatal(v ...any) {
	logger.Fatal(v...)
}

// Debugf 打印Debug日志.
func Debugf(format string, v ...any) {
	logger.Debugf(format, v...)
}

// Infof 打印Info日志.
func Infof(format string, v ...any) {
	logger.Infof(format, v...)
}

// Warnf 打印Warn日志.
func Warnf(format string, v ...any) {
	logger.Warnf(format, v...)
}

// Errorf 打印Error日志.
func Errorf(format string, v ...any) {
	logger.Errorf(format, v...)
}

// Fatalf 打印Fatal日志.
func Fatalf(format string, v ...any) {
	logger.Fatalf(format, v...)
}

// Output 打印数据日志.
func Output(v ...any) {
	dataLogger.Info(v...)
}

// GetTraceID 根据ctx获取 追踪id.
func GetTraceID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	val, _ := ctx.Value(CtxTraceKey("trace-id")).(string)

	return val
}
