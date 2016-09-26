package main

import (
	"sync"
)

type idGeneratorMutex struct {
	counter int32
	sync.RWMutex
}

func (g *idGeneratorMutex) Generate() int32 {
	g.Lock()
	defer g.Unlock()
	g.counter += 1
	return g.counter
}

func (g *idGeneratorMutex) Current() int32 {
	g.RLock()
	defer g.RUnlock()
	return g.counter
}

func newIdGeneratorMutex() *idGeneratorMutex {
	return new(idGeneratorMutex)
}
