package flood

import "time"

// Result
type Result[T any] struct {
	// ID
	ID string `json:"id"`

	// Queue
	Queue string `json:"queue"`

	// TenantID
	TenantID string `json:"tenant_id,omitempty"`

	// Args
	Args T `json:"args,omitempty"`

	// MaxRetry
	MaxRetry int `json:"max_retry,omitempty"`

	// Retried
	Retried int `json:"retried,omitempty"`

	// LastErr
	LastErr string `json:"last_err,omitempty"`

	// LastFailedAt
	LastFailedAt time.Time `json:"last_failed_at,omitempty"`

	// Timeout
	Timeout time.Duration `json:"timeout,omitempty"`

	// Deadline
	Deadline time.Time `json:"deadline,omitempty"`

	// Group
	Group string `json:"group,omitempty"`

	// NextProcessAt
	NextProcessAt time.Time `json:"next_process_at,omitempty"`

	// IsOrphaned
	IsOrphaned bool `json:"is_orphaned"`

	// Retention
	Retention time.Duration `json:"retention,omitempty"`

	// CompletedAt
	CompletedAt time.Time `json:"completed,omitempty"`

	// Result
	Result any `json:"result,omitempty"`
}
