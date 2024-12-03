package flood

import "time"

// Options
type Options struct {
	TenantID  string         `json:"tenant_id"`
	TaskID    *string        `json:"task_id,omitempty"`
	Queue     *string        `json:"queue,omitempty"`
	Timeout   *time.Duration `json:"timeout,omitempty"`
	MaxRetry  *int           `json:"max_retry,omitempty"`
	ProcessIn *time.Duration `json:"process_in,omitempty"`
	ProcessAt *time.Time     `json:"process_at,omitempty"`
	Deadline  *time.Time     `json:"deadline,omitempty"`
	Retention *time.Duration `json:"retention,omitempty"`
	Group     *string        `json:"group,omitempty"`
	Unique    *time.Duration `json:"unique,omitempty"`
}

// Option
type Option func(*Options)

// WithTenantID
func WithTenantID(v string) Option {
	return func(opts *Options) {
		opts.TenantID = v
	}
}

// WithTaskID
func WithTaskID(v string) Option {
	return func(opts *Options) {
		opts.TaskID = &v
	}
}

// WithMaxRetry
func WithMaxRetry(v int) Option {
	return func(opts *Options) {
		opts.MaxRetry = &v
	}
}

// WithQueue
func WithQueue(v string) Option {
	return func(opts *Options) {
		opts.Queue = &v
	}
}

// WithTimeout
func WithTimeout(timeout time.Duration) Option {
	return func(opts *Options) {
		opts.Timeout = &timeout
	}
}

// WithProcessIn
func WithProcessIn(v time.Duration) Option {
	return func(opts *Options) {
		opts.ProcessIn = &v
	}
}

// WithProcessAt
func WithProcessAt(v time.Time) Option {
	return func(opts *Options) {
		opts.ProcessAt = &v
	}
}

// WithDeadline
func WithDeadline(v time.Time) Option {
	return func(opts *Options) {
		opts.Deadline = &v
	}
}

// WithRetention
func WithRetention(v time.Duration) Option {
	return func(opts *Options) {
		opts.Retention = &v
	}
}

// WithGroup
func WithGroup(v string) Option {
	return func(opts *Options) {
		opts.Group = &v
	}
}

// WithUnique
func WithUnique(v time.Duration) Option {
	return func(opts *Options) {
		opts.Unique = &v
	}
}
