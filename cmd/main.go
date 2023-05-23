package main

import (
	"net/http"

	"github.com/danielmirceathoughtworks/golang-errors/internal/handlers"
	"github.com/danielmirceathoughtworks/golang-errors/internal/infrastructure"
	"github.com/gin-gonic/gin"
)

// curl -X POST -H "Content-Type: application/json" -d '{
//   "title": "Example Title",
//   "description": "Example Description",
//   "price": 100
// }' http://localhost:8080/ads

func main() {
	store := infrastructure.NewMemoryStore()

	router := gin.Default()
	router.POST("/ads", func(c *gin.Context) {
		ad, err := handlers.InsertAd(c, &store)

		if err != nil {
			// Unwrapping
			// var unwrappedError *OriginalError
			// if errors.As(err, &OriginalError) { ... }
			if e, ok := err.(*handlers.JSONParsingError); ok {
				// We have access custom error info
				c.JSON(http.StatusBadRequest, e.Error())
			} else if e, ok := err.(*infrastructure.InfrastructureError); ok {
				// e.OriginalError
			} else {
				c.JSON(http.StatusInternalServerError, err.Error())
			}

			return
		}

		c.JSON(http.StatusCreated, ad)
	})
	router.Run("localhost:8080")
}
