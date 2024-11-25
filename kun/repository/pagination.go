package repository

import "github.com/x-challenges/raven/kun/model"

type CursorDirection = int

const (
	CursorNext CursorDirection = iota
	CursorPrev
)

type Cursor = string

func StartCursor[T model.Model](entries []T) Cursor {
	if len(entries) > 0 {
		return Cursor(entries[0].GetID())
	}
	return ""
}

func EndCursor[T model.Model](entries []T) Cursor {
	if len(entries) > 0 {
		return Cursor(entries[len(entries)-1].GetID())
	}
	return ""
}

type Pager interface {
	GetFirst() string
	GetAfter() int
}

type PageRequest struct {
	First string
	After int
}

func (pr *PageRequest) GetFirst() string { return pr.First }
func (pr *PageRequest) GetAfter() int    { return pr.After }

// PageInfo type
type PageInfo struct {
	StartCursor Cursor
	LastCursor  Cursor
	HasNext     bool
	HasPrevious bool
}

func (p *PageInfo) GetStart() Cursor     { return p.StartCursor }
func (p *PageInfo) GetEnd() Cursor       { return p.LastCursor }
func (p *PageInfo) GetHasNext() bool     { return p.HasNext }
func (p *PageInfo) GetHasPrevious() bool { return p.HasPrevious }
