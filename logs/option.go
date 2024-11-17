package logs

import (
	"github.com/joshqu1985/lego/broker"
)

type Level int8

const (
	DEBUG Level = iota - 1
	INFO
	WARN
	ERROR
	PANIC
	FATAL
)

const (
	Development = 0
	Production  = 1
)

type Option func(o *options)

type options struct {
	Environment int
	Level       Level
	Writers     []string
	WriterFile  *FileOption
	WriterKafka *KafkaOption
}

type KafkaOption struct {
	Producer broker.Producer
	Topic    string
}

type FileOption struct {
	Filename   string // 日志文件路径
	MaxSize    int    // 单文件最大尺寸 MB
	MaxBackups int    // 最大备份数
	MaxAge     int    // 最大保留天数
	Compress   bool   // 是否压缩
}

func WithDevelopment() Option {
	return func(o *options) {
		o.Environment = Development
	}
}

func WithProduction() Option {
	return func(o *options) {
		o.Environment = Production
	}
}

// WithLevel sets level.
func WithLevel(level Level) Option {
	return func(o *options) {
		o.Level = level
	}
}

// WithWriterFile output to file.
func WithWriterFile(fileWriter *FileOption) Option {
	return func(o *options) {
		o.Writers = append(o.Writers, "file")
		o.WriterFile = fileWriter
	}
}

// WithWriterConsole output to console.
func WithWriterConsole() Option {
	return func(o *options) {
		o.Writers = append(o.Writers, "console")
	}
}

// WithWriterKafka output to kafka.
func WithWriterKafka(producer broker.Producer, topic string) Option {
	return func(o *options) {
		o.WriterKafka = &KafkaOption{
			Producer: producer,
			Topic:    topic,
		}
		o.Writers = append(o.Writers, "kafka")
	}
}
