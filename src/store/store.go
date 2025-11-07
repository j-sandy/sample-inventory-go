package store

import (
	"sample-inventory-go/models"
	"sync"
)

type InventoryStore struct {
	Items   map[string]models.Item
	Sessions map[string]models.Session
	mu      sync.RWMutex
}

func NewInventoryStore() *InventoryStore {
	return &InventoryStore{
		Items:   make(map[string]models.Item),
		Sessions: make(map[string]models.Session),
	}
}

func (s *InventoryStore) AddItem(item models.Item) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Items[item.ItemCode] = item
}

func (s *InventoryStore) GetItem(itemCode string) (models.Item, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	item, ok := s.Items[itemCode]
	return item, ok
}

func (s *InventoryStore) UpdateItem(item models.Item) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Items[item.ItemCode] = item
}

func (s *InventoryStore) DeleteItem(itemCode string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.Items, itemCode)
}

func (s *InventoryStore) AddSession(session models.Session) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Sessions[session.Token] = session
}

func (s *InventoryStore) GetSession(token string) (models.Session, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	session, ok := s.Sessions[token]
	return session, ok
}

func (s *InventoryStore) DeleteSession(token string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.Sessions, token)
}

func (s *InventoryStore) SearchItems(itemCode, name, procurementDate, expiryDate string) []models.Item {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var results []models.Item
	for _, item := range s.Items {
		match := true
		if itemCode != "" && item.ItemCode != itemCode {
			match = false
		}
		if name != "" && item.Name != name {
			match = false
		}
		if procurementDate != "" && (item.ProcurementDate == nil || *item.ProcurementDate != procurementDate) {
			match = false
		}
		if expiryDate != "" && (item.ExpiryDate == nil || *item.ExpiryDate != expiryDate) {
			match = false
		}
		if match {
			results = append(results, item)
		}
	}
	return results
}
