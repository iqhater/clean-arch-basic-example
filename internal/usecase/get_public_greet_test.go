package usecase

import (
	infra "clean_arch_super_simple_example/internal/infrastructure"
	"testing"
)

const GREET_TITLE = "👋 Hello TestGopher"

// static check interface implementation
var _ infra.GreeterRepo = (*GreetMockDB)(nil)

// mock repository struct
type GreetMockDB struct{}

// mock implementation without db, only mock title
func (_ *GreetMockDB) GetGreet() *infra.UserGreetPublic {
	return &infra.UserGreetPublic{
		Title: GREET_TITLE,
	}
}

func TestGetGreetValid(t *testing.T) {

	// arrange
	repo := &GreetMockDB{}
	service := NewGreetService(repo)

	// act
	greeting := service.GetGreet("Bug Testor")

	// assert
	if greeting != GREET_TITLE+" Bug Testor!" {
		t.Errorf("Expected: %s, got: %s", GREET_TITLE+" Bug Testor!", greeting)
	}
}
