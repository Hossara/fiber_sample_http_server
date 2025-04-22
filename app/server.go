package app

import (
	"encoding/json"
	"fiber_http/config"
	"fiber_http/handlers"
	"fiber_http/store"
	"github.com/gofiber/fiber/v2"
	"os"
)

func InitServer(store store.Store) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	api := app.Group("/api/v1")

	handlers.RegisterItemsHandlers(api, store)

	return app
}

func ReadConfig(path *string) config.Config {
	if v := os.Getenv("CONFIG_PATH"); len(v) > 0 {
		*path = v
	}

	return config.MustReadConfig(*path)
}
