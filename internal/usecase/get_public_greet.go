package usecase

import (
	infra "clean_arch_basic_example/internal/infrastructure"
)

// static check interface implementation
var _ GreeterService = (*GreetService)(nil)

// usecase/service
type GreeterService interface {
	GetGreet(name string) (*infra.UserGreetPublic, error)
}

// service implementation
type GreetService struct {
	greetRepo infra.GreeterRepo
}

// init
func NewGreetService(greetRepo infra.GreeterRepo) *GreetService {
	return &GreetService{
		greetRepo: greetRepo,
	}
}

// get greeting implementation
func (g *GreetService) GetGreet(name string) (*infra.UserGreetPublic, error) {

	// business logic here. concatinate our "greeting" and user "name"
	publicGreeting, err := g.greetRepo.GetGreet()
	publicGreeting.Title = publicGreeting.Title + " " + name + "!"

	return publicGreeting, err
}
