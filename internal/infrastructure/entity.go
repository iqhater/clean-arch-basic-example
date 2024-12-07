package infrastructure

type GreetingsCounter uint

// entity/domain

// greeting for anonymous users
type UserGreetPublic struct {
	Title string // greeting
}

// greeting for registered users with greetings counter
type UserGreetPrivate struct {
	ID      int
	Title   string // greeting
	Counter GreetingsCounter
}
