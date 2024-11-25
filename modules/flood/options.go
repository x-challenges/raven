package flood

import "time"

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

type Option func(*Options)

func WithTenantID(v string) Option {
	return func(opts *Options) {
		opts.TenantID = v
	}
}

func WithTaskID(v string) Option {
	return func(opts *Options) {
		opts.TaskID = &v
	}
}

func WithMaxRetry(v int) Option {
	return func(opts *Options) {
		opts.MaxRetry = &v
	}
}

func WithQueue(v string) Option {
	return func(opts *Options) {
		opts.Queue = &v
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(opts *Options) {
		opts.Timeout = &timeout
	}
}

func WithProcessIn(v time.Duration) Option {
	return func(opts *Options) {
		opts.ProcessIn = &v
	}
}

func WithProcessAt(v time.Time) Option {
	return func(opts *Options) {
		opts.ProcessAt = &v
	}
}

func WithDeadline(v time.Time) Option {
	return func(opts *Options) {
		opts.Deadline = &v
	}
}

func WithRetention(v time.Duration) Option {
	return func(opts *Options) {
		opts.Retention = &v
	}
}

func WithGroup(v string) Option {
	return func(opts *Options) {
		opts.Group = &v
	}
}

func WithUnique(v time.Duration) Option {
	return func(opts *Options) {
		opts.Unique = &v
	}
}
