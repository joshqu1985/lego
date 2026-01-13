package coss

type (
	Option func(o *options)

	options struct {
		// 上传下载大文件时的块大小
		BulkSize int64
		// 上传下载大文件时的并发数
		Concurrency int
	}
)

// WithBulkSize sets upload、download bulk size.
func WithBulkSize(bulkSize int64) Option {
	return func(o *options) {
		o.BulkSize = bulkSize
	}
}

func WithConcurrency(concurrency int) Option {
	return func(o *options) {
		o.Concurrency = concurrency
	}
}
