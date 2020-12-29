package main

import "sync"

var (
	store = NewStore()
)

type Store struct {
	store map[string]string
	mu    sync.RWMutex
}

func NewStore() *Store {
	return &Store{store: make(map[string]string)}
}

func (s *Store) Set(key, val string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[key] = val
}

func (s *Store) Get(key string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	val, ok := s.store[key]
	return val, ok
}
