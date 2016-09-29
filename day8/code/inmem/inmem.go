package inmem

import (
	"regexp"
	"sync"
	"time"
)

func (im inMem) set(key string, value interface{}) {
	if im.closed {
		return
	}
	im.Lock()
	im.storage[key] = value
	im.Unlock()
}
func (im inMem) get(key string) (interface{}, error) {
	if im.closed {
		return nil, errNotFound
	}

	im.RLock()
	result, ok := im.storage[key]
	im.RUnlock()
	if !ok {
		return nil, errNotFound
	}
	return result, nil
}
func (im inMem) delete(key string) {
	if im.closed {
		return
	}
	im.Lock()
	delete(im.storage, key)
	im.Unlock()
}

func (im inMem) getMulti(pattern string) ([]interface{}, error) {
	matcher, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	result := make([]interface{}, 0)
	for key, value := range im.storage {
		if matcher.MatchString(key) {
			result = append(result, value)
		}
	}
	return result, nil
}

func (im inMem) expire(key string, delay time.Duration) {
	go func() {
		select {
		case <-im.closeChanel:
		case <-time.After(delay):
			im.delete(key)
		}
	}()
}
func (im inMem) close(notify chan<- struct{}) {
	im.closed = true
	close(im.closeChanel)
	close(notify)
}

type inMem struct {
	sync.RWMutex
	storage     map[string]interface{}
	closeChanel chan struct{}
	closed      bool
}

//Return new inmem storage
func New() *inMem {
	im := inMem{
		storage:     make(map[string]interface{}),
		closeChanel: make(chan struct{}),
	}
	return &im
}
