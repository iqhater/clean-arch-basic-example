package usecase

import (
	"testing"

	infra "github.com/iqhater/clean-arch-basic-example/internal/infrastructure"
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
