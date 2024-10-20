package store

import (
	"fmt"
	"log"
	"sync"
)

type SafeMap[K comparable, V any] struct {
	mutex sync.RWMutex
	data  map[K]V
}

var (
	instance *SafeMap[string, interface{}]
	once     sync.Once
)

// GetInstance returns the singleton instance of SafeMap
func GetInstance() *SafeMap[string, interface{}] {
	once.Do(func() {
		log.Printf("Rover Store is initialized")
		instance = &SafeMap[string, interface{}]{
			data: make(map[string]interface{}),
		}
	})
	return instance
}

func (m *SafeMap[K, V]) Insert(key K, value V) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.data[key] = value
}

func (m *SafeMap[K, V]) Get(key K) (V, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	val, ok := m.data[key]
	if !ok {
		var zero V
		return zero, fmt.Errorf("key %v not found", key)
	}
	return val, nil
}

func (m *SafeMap[K, V]) Update(key K, value V) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	_, ok := m.data[key]
	if !ok {
		return fmt.Errorf("key %v not found", key)
	}

	m.data[key] = value
	return nil
}

func (m *SafeMap[K, V]) Delete(key K) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	_, ok := m.data[key]
	if !ok {
		return fmt.Errorf("key %v not found", key)
	}

	delete(m.data, key)
	return nil
}

func (m *SafeMap[K, V]) ContainsKey(key K) bool {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	_, ok := m.data[key]
	return ok
}

func (m *SafeMap[K, V]) DeleteAll() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.data = make(map[K]V)
}
