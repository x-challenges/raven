package flood

type JobArgs interface {
	Kind() string
}

type JobMeta struct {
	TenantID string `json:"tenant_id,omitempty"`
	TaskID   string `json:"task_id,omitempty"`
	MaxRetry int    `json:"max_retry,omitempty"`
	Attempt  int    `json:"attempt,omitempty"`
	Queue    string `json:"queue,omitempty"`
}

type wrappedJobArgs[T JobArgs] struct {
	Args T       `json:"args,omitempty"`
	Meta JobMeta `json:"meta,omitempty"`
}

type Job[T JobArgs] struct {
	Args   T                      `json:"args,omitempty"`
	Meta   JobMeta                `json:"meta,omitempty"`
	Result map[string]interface{} `json:"result,omitempty"`
}
