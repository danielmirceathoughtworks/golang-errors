package infrastructure

import (
	"fmt"

	"github.com/danielmirceathoughtworks/golang-errors/internal/domain"
	"github.com/google/uuid"
)

type MemoryStore struct {
	entries map[uuid.UUID]domain.Ad
}

func NewMemoryStore() MemoryStore {
	return MemoryStore{
		entries: make(map[uuid.UUID]domain.Ad),
	}
}

type InfrastructureError struct {
	OriginalError error
	Message       string
}

func (e *InfrastructureError) Error() string {
	return fmt.Sprintf("%s: %s", e.Message, e.OriginalError.Error())
}

func (a *MemoryStore) Add(ad *domain.Ad) (*domain.Ad, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		// return nil, fmt.Errorf("error generating id: %v", err)
		return nil, &InfrastructureError{
			OriginalError: err,
			Message:       "Failed to generate UUID",
		}
	}

	ad.Id = uuid
	a.entries[ad.Id] = *ad
	return ad, nil
}
