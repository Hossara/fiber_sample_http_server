package main

import (
	"fiber_http/app"
	"fiber_http/model"
	"fiber_http/store"

	"flag"
	"fmt"
)

var configPath = flag.String("config", "./", "Path to service config file")

func main() {
	flag.Parse()

	cfg := app.ReadConfig(configPath)

	items := []model.Item{
		{ID: 1, Name: "Laptop", Description: "Thin and light"},
		{ID: 2, Name: "Phone", Description: "Flagship phone"},
	}

	s := store.NewMemoryStore(items)

	fiberApp := app.InitServer(s)

	fiberApp.Listen(fmt.Sprintf(":%d", cfg.Server.Port))
}
