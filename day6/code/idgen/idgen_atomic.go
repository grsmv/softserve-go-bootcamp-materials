package main

import "sync/atomic"

type idGeneratorAtomic struct {
	counter int32
}

func (g *idGeneratorAtomic) Generate() int32 {
	return atomic.AddInt32(&g.counter, 1)
}

func (g *idGeneratorAtomic) Current() int32 {
	return atomic.LoadInt32(&g.counter)
}

func newIdGeneratorAtomic() *idGeneratorAtomic {
	return new(idGeneratorAtomic)
}
