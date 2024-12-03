package flood

// JobArgs
type JobArgs interface {
	Kind() string
}

// JobMeta
type JobMeta struct {
	TenantID string `json:"tenant_id,omitempty"`
	TaskID   string `json:"task_id,omitempty"`
	MaxRetry int    `json:"max_retry,omitempty"`
	Attempt  int    `json:"attempt,omitempty"`
	Queue    string `json:"queue,omitempty"`
}

// WrappedJobArgs
type wrappedJobArgs[T JobArgs] struct {
	Meta JobMeta `json:"meta,omitempty"`
	Args T       `json:"args,omitempty"`
}

// Job
type Job[T JobArgs] struct {
	Meta   JobMeta                `json:"meta,omitempty"`
	Args   T                      `json:"args,omitempty"`
	Result map[string]interface{} `json:"result,omitempty"`
}
