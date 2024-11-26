package messages

import (
	"time"
)

// EventMetadata
type EventMetadata struct {
	CloudID        string      `json:"cloud_id"`
	FolderID       string      `json:"folder_id"`
	EventID        string      `json:"event_id"`
	EventType      string      `json:"event_type"`
	TracingContext interface{} `json:"tracing_context,omitempty"`
	CreatedAt      time.Time   `json:"created_at"`
}

// Message
type Message[T any] struct {
	EventMetadata EventMetadata `json:"event_metadata"`
	Details       T             `json:"details"`
}

// Messages
type Messages[T any] struct {
	Messages []Message[T] `json:"messages"`
}
