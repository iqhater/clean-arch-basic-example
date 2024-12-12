package usecase

import (
	infra "clean_arch_basic_example/internal/infrastructure"
	"errors"
	"sync"
	"testing"
)

const GREET_PRIVATE_TITLE = "👋 Hello Private_Gopher"

// static check interface implementation
var _ infra.GreeterPrivateRepo = (*GreetPrivateMockDB)(nil)

// mock repository struct
type GreetPrivateMockDB struct {
	ids map[int]infra.GreetingsCounter
	mu  sync.RWMutex
}

// mock implementation without db, only mock title
func (db *GreetPrivateMockDB) GetGreetingsWithCounter(id int) (*infra.UserGreetPrivate, error) {

	if !db.isIDExist(id) {
		return &infra.UserGreetPrivate{}, errors.New("ID does not exist!")
	}

	return &infra.UserGreetPrivate{
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
	if _, ok := db.ids[id]; !ok {
		return false
	}
	return true
}

// mock inner func
func (db *GreetPrivateMockDB) incrementCounter(id int) infra.GreetingsCounter {
	db.mu.Lock()
	defer db.mu.Unlock()

	// wrtie to map with lock
	db.ids[id]++
	counter := db.ids[id]
	return counter
}

func TestGetGreetingsWithCounterValid(t *testing.T) {

	// arrange
	repo := &GreetPrivateMockDB{
		ids: map[int]infra.GreetingsCounter{
			1: 0,
		},
	}
	service := NewPrivateGreetService(repo)

	// act
	greeting, _ := service.GetGreetingsWithCounter(1, "Bug Testor")

	// assert

	// check title
	if greeting.Title != GREET_PRIVATE_TITLE+" Bug Testor!" {
		t.Errorf("Expected: %s, got: %s", GREET_PRIVATE_TITLE+" Bug Testor!", greeting.Title)
	}

	// check counter
	if greeting.Counter != 1 {
		t.Errorf("Expected: %d, got: %d", 1, greeting.Counter)
	}
}

func TestGetGreetingsWithCounterIDValid(t *testing.T) {

	// arrange
	repo := &GreetPrivateMockDB{
		ids: map[int]infra.GreetingsCounter{
			42: 0,
		},
	}
	service := NewPrivateGreetService(repo)

	// act
	validID := 42
	_, err := service.GetGreetingsWithCounter(validID, "")

	// assert

	// check id error
	if err != nil {
		t.Errorf("Expected: %v, got: ID: %d '%v'", nil, validID, err)
	}
}

func TestGetGreetingsWithCounterIDInvalid(t *testing.T) {

	// arrange
	repo := &GreetPrivateMockDB{
		ids: map[int]infra.GreetingsCounter{
			42: 0,
		},
	}
	service := NewPrivateGreetService(repo)

	// act
	invalidID := 39
	_, err := service.GetGreetingsWithCounter(invalidID, "")

	// assert

	// check id error
	if err == nil {
		t.Errorf("Expected: %v, got: ID: %d '%v'", nil, invalidID, err)
	}
}
