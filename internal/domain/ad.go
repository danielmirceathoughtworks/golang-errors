package domain

import (
	"time"

	"github.com/google/uuid"
)

type Ad struct {
	Id          uuid.UUID
	Title       string
	Description string
	Price       int
	CreatedAt   time.Time
}
