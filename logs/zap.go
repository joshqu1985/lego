package logs

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	OUTPUT_CONSOLE = "console"
	OUTPUT_FILE    = "file"
)

type (
	// Logger 避免业务代码对zap直接依赖.
	Logger = zap.SugaredLogger
)

// NewZapLogger 创建zap.
func NewZapLogger(opts options) *Logger {
	return zap.New(newZapCore(opts), zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
}

func newZapCore(opts options) zapcore.Core {
	var encoder zapcore.Encoder
	if opts.Environment == Development {
		encoder = zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	} else {
		encoder = zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	}

	writers := make([]zapcore.WriteSyncer, 0)
	for _, w := range opts.Writers {
		switch w {
		case OUTPUT_CONSOLE:
			writers = append(writers, newConsoleWriter(opts))
		case OUTPUT_FILE:
			writers = append(writers, newFileWriter(opts))
		}
	}

	level := zap.NewAtomicLevel()
	level.SetLevel(toZapLevel(opts.Level))

	return zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writers...), level)
}

func newConsoleWriter(_ options) zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

func newFileWriter(opts options) zapcore.WriteSyncer {
	config := &lumberjack.Logger{}
	if opts.WriterFile == nil {
		config.Filename = "./logs/server.log"
		config.MaxSize = 100
		config.MaxBackups = 1
		config.MaxAge = 28
		config.Compress = true
	} else {
		config.Filename = opts.WriterFile.Filename
		config.MaxSize = opts.WriterFile.MaxSize
		config.MaxBackups = opts.WriterFile.MaxBackups
		config.MaxAge = opts.WriterFile.MaxAge
		config.Compress = opts.WriterFile.Compress
	}

	return zapcore.AddSync(config)
}

func toZapLevel(level Level) zapcore.Level {
	switch level {
	case DEBUG:
		return zap.DebugLevel
	case INFO:
		return zap.InfoLevel
	case WARN:
		return zap.WarnLevel
	case ERROR:
		return zap.ErrorLevel
	case PANIC:
		return zap.PanicLevel
	case FATAL:
		return zap.FatalLevel
	}

	return zap.InfoLevel
}
