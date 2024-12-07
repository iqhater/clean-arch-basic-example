package infrastructure

import (
	"errors"
	"sync"
)

// static check interface implementation
var _ GreeterPrivateRepo = (*GreetPrivateFileDB)(nil)

// repository
type GreeterPrivateRepo interface {
	GetGreetingsWithCounter(id int) (*UserGreetPrivate, error)
}

// db implementation
type GreetPrivateFileDB struct {
	path string
	ids  map[int]GreetingsCounter
	mu   sync.RWMutex
}

// init
func NewPrivateGreetRepository(db string) *GreetPrivateFileDB {
	return &GreetPrivateFileDB{
		path: db,
		ids: map[int]GreetingsCounter{
			1: 0,
			2: 0,
			3: 0,
		}, // add a several id's already exist
	}
}

// get from file/db implementation
func (db *GreetPrivateFileDB) GetGreetingsWithCounter(id int) (*UserGreetPrivate, error) {

	if !db.isIDExist(id) {
		return &UserGreetPrivate{}, errors.New("ID does not exist!")
	}

	return &UserGreetPrivate{
		ID:      id,
		Title:   string(readDataFromFile(db.path)),
		Counter: db.incrementCounter(id),
	}, nil
}

func (db *GreetPrivateFileDB) isIDExist(id int) bool {
	db.mu.RLock()
	defer db.mu.RUnlock()

	// read from map with rlock
	if _, ok := db.ids[id]; !ok {
		return false
	}
	return true
}

func (db *GreetPrivateFileDB) incrementCounter(id int) GreetingsCounter {
	db.mu.Lock()
	defer db.mu.Unlock()

	// wrtie to map with lock
	db.ids[id]++
	counter := db.ids[id]
	return counter
}
