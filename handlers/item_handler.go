package handlers

import (
	"github.com/gofiber/fiber/v2"

	"fiber_http/store"
	"net/http"
)

func RegisterItemsHandlers(router fiber.Router, store store.Store) {
	itemsGroup := router.Group("/items")

	itemsGroup.Get("", GetItems(store))
}

// GetItems handles GET /items requests.
func GetItems(store store.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		items, err := store.GetItems()

		if err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "no items found",
			})
		}

		return c.JSON(items)
	}
}
