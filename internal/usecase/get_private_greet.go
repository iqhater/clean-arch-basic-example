package usecase

import (
	infra "clean_arch_basic_example/internal/infrastructure"
)

// static check interface implementation
var _ GreeterPrivateService = (*GreetPrivateService)(nil)

// usecase/service
type GreeterPrivateService interface {
	GetGreetingsWithCounter(id int, name string) (*infra.UserGreetPrivate, error)
}

// service implementation
type GreetPrivateService struct {
	greetPrivateRepo infra.GreeterPrivateRepo
}

// init
func NewPrivateGreetService(greetPrivateRepo infra.GreeterPrivateRepo) *GreetPrivateService {
	return &GreetPrivateService{
		greetPrivateRepo: greetPrivateRepo,
	}
}

// get private greeting implementation
func (g *GreetPrivateService) GetGreetingsWithCounter(id int, name string) (*infra.UserGreetPrivate, error) {

	// return greetings counter for registered user
	privateGreeting, err := g.greetPrivateRepo.GetGreetingsWithCounter(id)
	privateGreeting.Title = privateGreeting.Title + " " + name + "!"

	return privateGreeting, err
}
