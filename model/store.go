package model

import (
	"sync"
	"time"
)

type Store struct {
	sync.RWMutex
	data map[string]*Pad
	logTick <-chan time.Time
}

func NewStore() Store {
	s := Store{}
	s.data = make(map[string]*Pad)
	s.logTick = time.Tick(4 * time.Second)
	return s
}

func (s Store) Add(p *Pad) {
	p.GenerateId()
	p.EnsureExpiration()
	s.Lock()
	s.data[p.Id] = p
	s.Unlock()
	go func(s Store, p *Pad) {
		select {
		case <-time.After(time.Duration(p.ExpiresInSeconds) * time.Second):
			s.Expire(p.Id)
		}
	}(s, p)
}

func (s Store) Consume(id string) (*Pad, bool) {
	s.RLock()
	p, ok := s.data[id]
	s.RUnlock()
	if ok {
		p.Use()
		if p.IsConsumed() {
			s.Expire(id)
		}
	}
	return p, ok
}

func (s Store) Expire(id string) {
	s.Lock()
	defer s.Unlock()
	delete(s.data, id)
}
