package infrastructure

import "github.com/iqhater/clean-arch-basic-example/internal/domain"

const GREET_TITLE = "👋 Hello Gopher"

// static check interface implementation
var _ GreeterRepo = (*GreetMockDB)(nil)

// mock repository struct
type GreetMockDB struct{}

// mock implementation without db, only mock title
func (_ *GreetMockDB) GetGreet() (*domain.UserGreetPublic, error) {
	return &domain.UserGreetPublic{
		Title: GREET_TITLE,
	}, nil
}
