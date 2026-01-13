package logs

const (
	DEBUG Level = iota - 1
	INFO
	WARN
	ERROR
	PANIC
	FATAL

	Development = 0
	Production  = 1
)

type (
	Level int8

	Option func(o *options)

	options struct {
		WriterFile  *FileOption
		Writers     []string
		Environment int
		Level       Level
	}

	FileOption struct {
		Filename   string // 日志文件路径
		MaxSize    int    // 单文件最大尺寸 MB
		MaxBackups int    // 最大备份数
		MaxAge     int    // 最大保留天数
		Compress   bool   // 是否压缩
	}
)

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
		o.Writers = append(o.Writers, OUTPUT_FILE)
		o.WriterFile = fileWriter
	}
}

// WithWriterConsole output to console.
func WithWriterConsole() Option {
	return func(o *options) {
		o.Writers = append(o.Writers, OUTPUT_CONSOLE)
	}
}
