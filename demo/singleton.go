// +build all demo sync singleton

package demo

import (
	"sync"
	"sync/atomic"
)

// Singleton struct
type Singleton struct {
	initialized bool
}

// variables for GetInstance() func
var instance *Singleton
var singletonAtom uint32
var singletonMutex sync.Mutex

// variables for GetSingleton() func
var singleton *Singleton
var once sync.Once

// GetInstance demonstrates an idiomatic approach to get a singleton instance
func GetInstance() *Singleton {
	if atomic.LoadUint32(&singletonAtom) == 1 {
		return instance
	}

	singletonMutex.Lock()
	defer singletonMutex.Unlock()

	if singletonAtom == 0 {
		instance = &Singleton{initialized: true}
		atomic.StoreUint32(&singletonAtom, 1)
	}

	return instance
}

// GetSingleton demonstrates using sync.Once to get a singleton
func GetSingleton() *Singleton {
	once.Do(func() {
		singleton = &Singleton{initialized: true}
	})
	return singleton
}
