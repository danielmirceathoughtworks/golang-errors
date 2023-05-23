package handlers

import (
	"github.com/danielmirceathoughtworks/golang-errors/internal/domain"
	"github.com/danielmirceathoughtworks/golang-errors/internal/infrastructure"
	"github.com/gin-gonic/gin"
)

type JSONParsingError struct{}

func (e *JSONParsingError) Error() string {
	return "Received input is not valid json"
}

func InsertAd(c *gin.Context, store *infrastructure.MemoryStore) (*domain.Ad, error) {
	ad := domain.Ad{}

	err := c.BindJSON(&ad)
	if err != nil {
		return nil, &JSONParsingError{}
	}

	// add validations here and return if they fail

	return store.Add(&ad)
}
