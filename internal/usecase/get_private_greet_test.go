package usecase

import (
	infra "clean_arch_basic_example/internal/infrastructure"
	"testing"
)

func TestGetGreetingsWithCounterValid(t *testing.T) {

	// arrange
	repo := &infra.GreetPrivateMockDB{
		IDs: map[int]infra.GreetingsCounter{
			1: 0,
		},
	}
	service := NewPrivateGreetService(repo)

	// act
	greeting, _ := service.GetGreetingsWithCounter(1, "Bug Testor")

	// assert

	// check title
	if greeting.Title != infra.GREET_PRIVATE_TITLE+" Bug Testor!" {
		t.Errorf("Expected: %s, got: %s", infra.GREET_PRIVATE_TITLE+" Bug Testor!", greeting.Title)
	}

	// check counter
	if greeting.Counter != 1 {
		t.Errorf("Expected: %d, got: %d", 1, greeting.Counter)
	}
}

func TestGetGreetingsWithCounterIDValid(t *testing.T) {

	// arrange
	repo := &infra.GreetPrivateMockDB{
		IDs: map[int]infra.GreetingsCounter{
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
	repo := &infra.GreetPrivateMockDB{
		IDs: map[int]infra.GreetingsCounter{
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
