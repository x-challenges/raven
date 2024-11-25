package model

// Model interface
type Model interface {
	// GetID
	GetID() ID

	// TableName
	TableName() string
}
