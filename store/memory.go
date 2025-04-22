package store

import (
	"errors"
	"fiber_http/model"
)

// Store defines interface for getting items.
type Store interface {
	GetItems() ([]model.Item, error)
}

// MemoryStore is an in-memory implementation of Store.
type MemoryStore struct {
	items []model.Item
}

// NewMemoryStore creates a new memory store with initial items.
func NewMemoryStore(items []model.Item) *MemoryStore {
	return &MemoryStore{items: items}
}

// GetItems returns items if present, or error if none.
func (m *MemoryStore) GetItems() ([]model.Item, error) {
	if len(m.items) == 0 {
		return nil, ErrNoItems
	}
	return m.items, nil
}

// ErrNoItems returned when there are no items in the store.
var ErrNoItems = errors.New("no items found")
