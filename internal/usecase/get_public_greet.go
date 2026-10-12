package usecase

import (
	"github.com/iqhater/clean-arch-basic-example/internal/domain"
	infra "github.com/iqhater/clean-arch-basic-example/internal/infrastructure"
)

// static check interface implementation
var _ GreeterService = (*GreetService)(nil)

// usecase/service
type GreeterService interface {
	GetGreet(name string) (*domain.UserGreetPublic, error)
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
func (g *GreetService) GetGreet(name string) (*domain.UserGreetPublic, error) {

	// business logic here. concatinate our "greeting" and user "name"
	publicGreeting, err := g.greetRepo.GetGreet()
	publicGreeting.Title = publicGreeting.Title + " " + name + "!"

	return publicGreeting, err
}
