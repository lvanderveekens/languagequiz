package exercise

import (
	"time"

	"github.com/google/uuid"
)

type Exercise struct {
	ID        uuid.UUID
	CreatedAt time.Time
}
