package trigger

import (
	"go.uber.org/zap"
)

// Factory
type Factory struct {
	logger       *zap.Logger
	errorHandler ErrorHandler
}

// NewFactory
func NewFactory(logger *zap.Logger) *Factory {
	return &Factory{
		logger: logger,
	}
}

// Timer
func (f *Factory) Timer(callback TimerCallback, options ...OptionFunc) *Timer {
	var opts = &Options{
		callback:     callback,
		errorHandler: DefaultErrorHandler,
	}

	applyOptions(opts, options...)

	return &Timer{
		logger:       f.logger,
		callback:     callback,
		errorHandler: opts.errorHandler,
	}
}

// Queue
func (f *Factory) Queue(callback QueueCallback, options ...OptionFunc) *Queue {
	var opts = &Options{
		callback:     callback,
		errorHandler: DefaultErrorHandler,
	}

	applyOptions(opts, options...)

	return &Queue{
		logger:       f.logger,
		callback:     callback,
		errorHandler: opts.errorHandler,
	}
}

// Stream
func (f *Factory) Stream(callback StreamCallback, options ...OptionFunc) *Stream {
	var opts = &Options{
		callback:     callback,
		errorHandler: DefaultErrorHandler,
	}

	applyOptions(opts, options...)

	return &Stream{
		logger:       f.logger,
		callback:     callback,
		errorHandler: opts.errorHandler,
	}
}

// Options
type Options struct {
	callback     any
	errorHandler ErrorHandler
}

// OptionFunc
type OptionFunc func(*Options)

// ApplyOptions
func applyOptions(options *Options, opts ...OptionFunc) {
	for _, opt := range opts {
		opt(options)
	}
}

// WithErrorHandler
func WithErrorHandler(errorHandler ErrorHandler) OptionFunc {
	return func(o *Options) {
		o.errorHandler = errorHandler
	}
}
