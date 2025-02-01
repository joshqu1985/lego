package coss

type Option func(o *options)

type options struct {
	// BulkSize
	BulkSize int64
}

// WithBulkSize sets upload、download bulk size.
func WithBulkSize(bulkSize int64) Option {
	return func(o *options) {
		o.BulkSize = bulkSize
	}
}
