package entities

type User struct {
	ID       string `firestone:"id"`
	Name     string `firestone:"name"`
	Email    string `firestone:"email"`
	Password string `firestone:"password"`
}
