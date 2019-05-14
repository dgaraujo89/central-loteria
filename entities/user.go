package entities

// User entity
type User struct {
	ID       string `firestone:"id" json:"id"`
	Name     string `firestone:"name" json:"name"`
	Email    string `firestone:"email" json:"email"`
	Password string `firestone:"password" json:"password"`
}
