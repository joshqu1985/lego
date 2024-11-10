package logs

import (
	"context"
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"lego/broker"
)

// Logger 避免业务代码对zap直接依赖
type Logger = zap.SugaredLogger

// NewZapLogger 创建zap
func NewZapLogger(opts options) *Logger {
	zap.NewProduction()
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
		case "console":
			writers = append(writers, newConsoleWriter(opts))
		case "file":
			writers = append(writers, newFileWriter(opts))
		case "kafka":
			writers = append(writers, newKafkaWriter(opts))
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

func newKafkaWriter(opts options) zapcore.WriteSyncer {
	return &kafkaWriter{
		producer: opts.WriterKafka.Producer,
		topic:    opts.WriterKafka.Topic,
	}
}

type kafkaWriter struct {
	producer broker.Producer
	topic    string
}

func (this *kafkaWriter) Write(p []byte) (int, error) {
	if this.producer == nil {
		return 0, fmt.Errorf("producer is nil")
	}
	err := this.producer.Send(context.Background(), this.topic, &broker.Message{Payload: p})
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

func (k *kafkaWriter) Sync() error {
	return nil
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
