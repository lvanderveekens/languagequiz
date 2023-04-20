package exercise

import (
	"time"

	"github.com/google/uuid"
)

type Exercise struct {
	ID        uuid.UUID
	CreatedAt time.Time
}

func New(id uuid.UUID, createdAt time.Time) Exercise {
	return Exercise{ID: id, CreatedAt: createdAt}
}
