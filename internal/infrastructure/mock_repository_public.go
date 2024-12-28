package infrastructure

const GREET_TITLE = "👋 Hello Gopher"

// static check interface implementation
var _ GreeterRepo = (*GreetMockDB)(nil)

// mock repository struct
type GreetMockDB struct{}

// mock implementation without db, only mock title
func (_ *GreetMockDB) GetGreet() (*UserGreetPublic, error) {
	return &UserGreetPublic{
		Title: GREET_TITLE,
	}, nil
}
