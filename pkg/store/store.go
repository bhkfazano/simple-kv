package store

import (
	"fmt"
	"sync"

	skvError "github.com/bhkfazano/simple-kv/pkg/error"
)

type Store[K comparable, V any] interface {
	Put(key K, value V) error
	Get(key K) (V, error)
	Update(key K, value V) error
	Delete(key K) (V, error)
}

type SimpleKV[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

func NewSimpleKV[K comparable, V any]() *SimpleKV[K, V] {
	return &SimpleKV[K, V]{
		data: make(map[K]V),
	}
}

func (s *SimpleKV[K, V]) Put(key K, value V) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value
	return nil
}

func (s *SimpleKV[K, V]) Get(key K) (V, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, ok := s.data[key]
	if !ok {
		return value, skvError.NewNotFoundError(fmt.Sprintf("key not found: %v", key))
	}
	return value, nil
}

func (s *SimpleKV[K, V]) Update(key K, value V) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.contains(key) {
		return skvError.NewNotFoundError(fmt.Sprintf("key not found: %v", key))
	}

	s.data[key] = value
	return nil
}

// Not thread-safe: do not use without a lock
func (s *SimpleKV[K, V]) contains(key K) bool {
	_, ok := s.data[key]
	return ok
}

func (s *SimpleKV[K, V]) Delete(key K) (V, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	value, ok := s.data[key]
	if !ok {
		return value, skvError.NewNotFoundError(fmt.Sprintf("key not found: %v", key))
	}

	delete(s.data, key)
	return value, nil
}
