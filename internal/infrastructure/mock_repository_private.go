package infrastructure

import (
	"errors"
	"sync"
)

const GREET_PRIVATE_TITLE = "👋 Hello Private_Gopher"

// static check interface implementation
var _ GreeterPrivateRepo = (*GreetPrivateMockDB)(nil)

// mock repository struct
type GreetPrivateMockDB struct {
	IDs map[int]GreetingsCounter
	mu  sync.RWMutex
}

// mock implementation without db, only mock title
func (db *GreetPrivateMockDB) GetGreetingsWithCounter(id int) (*UserGreetPrivate, error) {

	if !db.isIDExist(id) {
		return &UserGreetPrivate{}, errors.New("ID does not exist!")
	}

	return &UserGreetPrivate{
		ID:      id,
		Title:   GREET_PRIVATE_TITLE,
		Counter: db.incrementCounter(id),
	}, nil
}

// mock id check
func (db *GreetPrivateMockDB) isIDExist(id int) bool {
	db.mu.RLock()
	defer db.mu.RUnlock()

	// read from map with rlock
	if _, ok := db.IDs[id]; !ok {
		return false
	}
	return true
}

// mock inner func
func (db *GreetPrivateMockDB) incrementCounter(id int) GreetingsCounter {
	db.mu.Lock()
	defer db.mu.Unlock()

	// wrtie to map with lock
	db.IDs[id]++
	counter := db.IDs[id]
	return counter
}
