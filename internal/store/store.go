package store

import (
	"fmt"
	"log"
	"sync"

	"github.com/subrotokumar/rover/internal/types"
)

type SafeMap[K comparable, V any] struct {
	mutex sync.RWMutex
	ttl   []string
	data  []map[K]V
}

var (
	instance *SafeMap[string, types.StoredValue]
	once     sync.Once
)

// GetInstance returns the singleton instance of SafeMap
func GetInstance() *SafeMap[string, types.StoredValue] {
	once.Do(func() {
		log.Printf("Rover Store is initialized")
		ttlList := []string{}
		maps := make([]map[string]types.StoredValue, 16)
		for i := range maps {
			maps[i] = make(map[string]types.StoredValue)
		}
		instance = &SafeMap[string, types.StoredValue]{
			data: maps,
			ttl:  ttlList,
		}
	})
	return instance
}

func (m *SafeMap[K, V]) Insert(db int, key K, value V) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.data[db][key] = value
}

func (m *SafeMap[K, V]) Get(db int, key K) (V, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	val, ok := m.data[db][key]
	if !ok {
		var zero V
		return zero, fmt.Errorf("key %v not found", key)
	}
	return val, nil
}

func (m *SafeMap[K, V]) Update(db int, key K, value V) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	_, ok := m.data[db][key]
	if !ok {
		return fmt.Errorf("key %v not found", key)
	}

	m.data[db][key] = value
	return nil
}

func (m *SafeMap[K, V]) Delete(db int, key K) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	_, ok := m.data[db][key]
	if !ok {
		return fmt.Errorf("key %v not found", key)
	}

	delete(m.data[db], key)
	return nil
}

func (m *SafeMap[K, V]) ContainsKey(db int, key K) bool {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	_, ok := m.data[db][key]
	return ok
}

func (m *SafeMap[K, V]) DeleteAll(db int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.data[db] = make(map[K]V)
}
