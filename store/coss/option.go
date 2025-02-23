package coss

type Option func(o *options)

type options struct {
	// BulkSize 下载大文件时的块大小
	BulkSize int64
	// Concurrency 下载大文件时的并发数
	Concurrency int
}

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
