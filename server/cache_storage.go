package server

import "sync"

type Store struct {

    data map[string]string
    mu   sync.RWMutex
}

func NewStorageCache() *Store {

    return &Store{
        data: make(map[string]string),
    }

}

func (s *Store) Set(key, value string) {

    s.mu.Lock()
    defer s.mu.Unlock()
    s.data[key] = value
}

func (s *Store) Get(key string) (string, bool) {

    s.mu.RLock()
    defer s.mu.RUnlock()
    value, found := s.data[key]
    return value, found
}