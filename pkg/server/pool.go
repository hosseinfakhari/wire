package server

import (
	"log"
	"sync/atomic"
)

type ServerPool struct {
	backends []*Backend
	current  uint64
	lg       *log.Logger
}

func NewServerPool(lg *log.Logger) *ServerPool {
	lg.Println("Server Pool Created")
	return &ServerPool{
		backends: make([]*Backend, 0),
		current:  uint64(0),
		lg:       lg,
	}
}

func (s *ServerPool) AddBackend(backend *Backend) {
	s.lg.Println("New Backend Added")
	s.backends = append(s.backends, backend)
}

func (s *ServerPool) NextIndex() int {
	return int(atomic.AddUint64(&s.current, 1) % uint64(len(s.backends)))
}

func (s *ServerPool) GetNextPeer() *Backend {
	next := s.NextIndex()
	l := len(s.backends) + next
	s.lg.Printf("Trying to access %d backend", next)
	for i := next; i < l; i++ {
		idx := i % len(s.backends)
		if s.backends[idx].IsAlive() {
			s.lg.Printf("Backend %d selected", idx)
			if i != next {
				atomic.StoreUint64(&s.current, uint64(idx))
			}
			return s.backends[idx]
		}
	}
	s.lg.Println("No Suitable Backend Found!")
	return nil
}
