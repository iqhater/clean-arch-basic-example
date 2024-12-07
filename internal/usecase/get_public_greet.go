package usecase

import (
	infra "clean_arch_super_simple_example/internal/infrastructure"
)

// static check interface implementation
var _ GreeterService = (*GreetService)(nil)

// usecase/service
type GreeterService interface {
	GetGreet(name string) string
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
func (g *GreetService) GetGreet(name string) string {

	// business logic here. concatinate our "greeting" and user "name"
	return g.greetRepo.GetGreet().Title + " " + name + "!"
}
