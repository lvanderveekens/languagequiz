package exercise

import (
	"time"
)

type Exercise struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(id string, createdAt, updatedAt time.Time) Exercise {
	return Exercise{ID: id, CreatedAt: createdAt, UpdatedAt: updatedAt}
}
