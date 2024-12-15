package recipes

import (
	"fmt"
)

type MemStore struct {
	data map[string]Recipe
}

func NewMemStore() *MemStore {
	return &MemStore{data: make(map[string]Recipe)}
}

func (m *MemStore) Add(name string, recipe Recipe) error {
	if _, exists := m.data[name]; exists {
		return fmt.Errorf("recipe already exists")
	}
	m.data[name] = recipe
	return nil
}

func (m *MemStore) Get(name string) (Recipe, error) {
	recipe, exists := m.data[name]
	if !exists {
		return Recipe{}, fmt.Errorf("recipe not found")
	}

	return recipe, nil
}

func (m *MemStore) Update(name string, recipe Recipe) error {
	if _, exists := m.data[name]; !exists {
		return fmt.Errorf("recipe not found")
	}

	m.data[name] = recipe
	return nil
}

func (m *MemStore) List() (map[string]Recipe, error) {
	return m.data, nil
}

func (m *MemStore) Remove(name string) error {
	if _, exists := m.data[name]; !exists {
		return fmt.Errorf("recipe not found")
	}

	delete(m.data, name)
	return nil
}
