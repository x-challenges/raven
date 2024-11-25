package repository

import "github.com/x-challenges/raven/kun/model"

type ptr[T any] interface {
	*T
	model.Model
}
