package task7

import (
	"sync"
)

// Конкурентная запись с использованием mutex
type MutexMap struct {
	mutex sync.Mutex
	data  map[int]int
}

func (muMap *MutexMap) Store(key int, value int) {
	muMap.mutex.Lock()
	muMap.data[key] = value
	muMap.mutex.Unlock()

}

// Конкурентная запись с использованием rwmutex
type RWMutexMap struct {
	rwmutex sync.RWMutex
	data    map[int]int
}

func (rwmMap *RWMutexMap) Store(key int, value int) {
	rwmMap.rwmutex.Lock()
	rwmMap.data[key] = value
	rwmMap.rwmutex.Unlock()
}

// Конкурентная запись с использованием sync.Map
// The Map type is optimized for two common use cases: (1) when the entry for a given
// key is only ever written once but read many times, as in caches that only grow,
// or (2) when multiple goroutines read, write, and overwrite entries for disjoint
// sets of keys.
func writeToSyncMap() {
	var m sync.Map
	m.Store("key", "value")
}
