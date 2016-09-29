package inmem

import "sync"

func (im inMem) set(key string, value interface{}) {
	im.Lock()
	im.storage[key] = value
	im.Unlock()
}
func (im inMem) get(key string) (interface{}, error) {
	im.RLock()
	result, ok := im.storage[key]
	im.RUnlock()
	if !ok {
		return nil, errNotFound
	}
	return result, nil
}
func (im inMem) delete(key string) {
	im.Lock()
	delete(im.storage, key)
	im.Unlock()
}

type inMem struct {
	sync.RWMutex
	storage map[string]interface{}
}

//Return new inmem storage
func New() *inMem {
	im := inMem{
		storage: make(map[string]interface{}),
	}
	return &im
}
