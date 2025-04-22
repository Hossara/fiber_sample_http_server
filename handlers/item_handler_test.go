package handlers

import (
	"github.com/gofiber/fiber/v2"

	"fiber_http/model"
	"fiber_http/store"

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setupApp(s store.Store) *fiber.App {
	app := fiber.New()
	api := app.Group("/api/v1")
	RegisterItemsHandlers(api, s)
	return app
}

func TestGetItems(t *testing.T) {
	tests := []struct {
		name       string
		items      []model.Item
		wantStatus int
		wantBody   string
	}{
		{
			name: "items exist",
			items: []model.Item{
				{ID: 1, Name: "Laptop", Description: "Thin and light"},
			},
			wantStatus: http.StatusOK,
			wantBody:   `[{"id":1,"name":"Laptop","description":"Thin and light"}]`,
		},
		{
			name:       "no items",
			items:      []model.Item{},
			wantStatus: http.StatusNotFound,
			wantBody:   `{"error":"no items found"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := setupApp(store.NewMemoryStore(tt.items))

			req := httptest.NewRequest(http.MethodGet, "/api/v1/items", nil)
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			if err != nil {
				t.Fatalf("error making request: %v", err)
			}

			if resp.StatusCode != tt.wantStatus {
				t.Errorf("expected status %d, got %d", tt.wantStatus, resp.StatusCode)
			}

			buf := make([]byte, resp.ContentLength)
			resp.Body.Read(buf)
			resp.Body.Close()

			got := strings.TrimSpace(string(buf))
			if got != tt.wantBody {
				t.Errorf("expected body %s, got %s", tt.wantBody, got)
			}
		})
	}
}
