package main

type idGeneratorChan struct {
	counter     int32
	generate    chan struct{}
	current     chan struct{}
	getGenerate chan int32
	getCurrent  chan int32
}

func (g *idGeneratorChan) generator() int32 {
	for {
		select {
		case <-g.generate:
			g.counter += 1
			g.getGenerate <- g.counter
		case <-g.current:
			g.getCurrent <- g.counter
		}
	}
}

func (g *idGeneratorChan) Generate() int32 {
	g.generate <- struct{}{}
	return <-g.getGenerate
}

func (g *idGeneratorChan) Current() int32 {
	g.current <- struct{}{}
	return <-g.getCurrent
}

func newIdGeneratorChan() *idGeneratorChan {
	g := idGeneratorChan{
		generate:    make(chan struct{}),
		current:     make(chan struct{}),
		getGenerate: make(chan int32),
		getCurrent:  make(chan int32),
	}
	go g.generator()
	return &g
}
