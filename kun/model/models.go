package model

import (
	"time"
)

// Model interface
type Model interface {
	// GetID
	GetID() ID

	// TableName
	TableName() string
}

type Base struct {
	ID        string    `gorm:"column:id;type:uuid" json:"id" validate:"required,uuid4"`
	CreatedAt time.Time `gorm:"column:created_at;default:now()" json:"created_at" validate:"required"`
}

// GetID
func (b Base) GetID() string { return b.ID }
