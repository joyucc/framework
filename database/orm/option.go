package orm

import "time"

type Option func(*Options)

type Options struct {
	Dialect     string
	DSN         string
	Active      int           // pool
	Idle        int           // pool
	IdleTimeout time.Duration // connect max life time.
}

func NewOptions(opts ...Option) *Options {
	options := &Options{}

	for _, opt := range opts {
		opt(options)
	}

	return options
}

func WithDialect(dialect string) Option {
	return func(o *Options) {
		o.Dialect = dialect
	}
}

func WithActive(act int) Option {
	return func(o *Options) {
		o.Active = act
	}
}

func WithIdle(idle int) Option {
	return func(o *Options) {
		o.Idle = idle
	}
}

func WithIdleTimeout(dur time.Duration) Option {
	return func(o *Options) {
		o.IdleTimeout = dur
	}
}

func WithDSN(uri string) Option {
	return func(o *Options) {
		o.DSN = uri
	}
}
