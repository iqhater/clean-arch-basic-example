package infrastructure

import "testing"

func TestIsIDExist(t *testing.T) {

	// arrange
	repo := NewPrivateGreetRepository("")
	id := 3

	// act
	result := repo.isIDExist(id)

	// assert
	if !result {
		t.Errorf("Expected: true, got: false")
	}
}

func TestIsIDNotExist(t *testing.T) {

	// arrange
	repo := NewPrivateGreetRepository("")
	id := 42

	// act
	result := repo.isIDExist(id)

	// assert
	if result {
		t.Errorf("Expected: true, got: false")
	}
}

func TestIncrementCounter(t *testing.T) {

	// arrange
	repo := NewPrivateGreetRepository("")
	id := 3
	prevValue := repo.ids[id]

	// act
	result := repo.incrementCounter(id)

	// assert
	if result <= prevValue {
		t.Errorf("Expected: %d, got: %d", prevValue+1, result)
	}
}

func TestGetGreetingsWithCounter(t *testing.T) {

	// arrange
	repo := NewPrivateGreetRepository(DB_PATH)
	id := 2
	privateGreeting := "👋 Hello Gopher"

	// act
	result, _ := repo.GetGreetingsWithCounter(id)

	// assert
	if result.ID != id {
		t.Errorf("Expected: %d, got: %d", id, result.ID)
	}

	if result.Title != privateGreeting {
		t.Errorf("Expected: %s, got: %s", privateGreeting, result.Title)
	}

	if result.Counter != 1 {
		t.Errorf("Expected: %d, got: %d", 1, result.Counter)
	}

	if result == nil {
		t.Errorf("Expected: %v, got: %v", nil, result)
	}
}

func TestGetGreetingsWithCounterIDNotExist(t *testing.T) {

	// arrange
	repo := NewPrivateGreetRepository(DB_PATH)
	id := 42

	// act
	result, err := repo.GetGreetingsWithCounter(id)

	// assert
	if err == nil {
		t.Errorf("Expected: %v, got: ID: %d '%v'", nil, id, err)
	}

	// empty zero value struct
	if result.ID != 0 {
		t.Errorf("Expected: %d, got: %d", 0, result.ID)
	}

	if result.Title != "" {
		t.Errorf("Expected: %s, got: %s", "", result.Title)
	}

	if result.Counter != 0 {
		t.Errorf("Expected: %d, got: %d", 0, result.Counter)
	}
}
