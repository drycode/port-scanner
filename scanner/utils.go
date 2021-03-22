package utils

import (
	"sync"
)

// SafeSlice ...
type SafeSlice struct {
	mu        sync.RWMutex
	OpenPorts []string
}

func (ss *SafeSlice) length() int {
	ss.mu.RLock()
	defer ss.mu.RUnlock()
	return len(ss.OpenPorts)
}

func (ss *SafeSlice) append(val string, wg *sync.WaitGroup) {
	ss.mu.Lock()
	ss.OpenPorts = append((*ss).OpenPorts, val)
	ss.mu.Unlock()
	wg.Done()
}
