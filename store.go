package main

import (
	"strconv"
	"sync"
)

type store struct {
	mu    sync.Mutex
	links map[string]string
	seq   int
}

func newStore() *store {
	return &store{links: make(map[string]string)}
}

func (s *store) save(url string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.seq++
	code := strconv.FormatInt(int64(s.seq), 36)
	s.links[code] = url
	return code
}

func (s *store) get(code string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	url, ok := s.links[code]
	return url, ok
}
