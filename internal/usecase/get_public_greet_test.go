package usecase

import (
	infra "clean_arch_basic_example/internal/infrastructure"
	"testing"
)

func TestGetGreetValid(t *testing.T) {

	// arrange
	repo := &infra.GreetMockDB{}
	service := NewGreetService(repo)

	// act
	greeting, _ := service.GetGreet("Bug Testor")

	// assert
	if greeting.Title != infra.GREET_TITLE+" Bug Testor!" {
		t.Errorf("Expected: %s, got: %s", infra.GREET_TITLE+" Bug Testor!", greeting)
	}
}
