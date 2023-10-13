package ydb

import (
	"time"

	"github.com/ydb-platform/ydb-go-sdk/v3"
	"gorm.io/gorm"

	"github.com/PotatoHD404/gorm-driver/internal/dialect"
)

type Option = dialect.Option

func With(opts ...ydb.Option) Option {
	return dialect.With(opts...)
}

func WithTablePathPrefix(tablePathPrefix string) Option {
	return dialect.WithTablePathPrefix(tablePathPrefix)
}

func WithMaxOpenConns(n int) Option {
	return dialect.WithMaxOpenConns(n)
}

func WithMaxIdleConns(n int) Option {
	return dialect.WithMaxIdleConns(n)
}

func WithConnMaxIdleTime(d time.Duration) Option {
	return dialect.WithConnMaxIdleTime(d)
}

func Open(dsn string, opts ...Option) gorm.Dialector {
	return dialect.New(dsn, opts...)
}
