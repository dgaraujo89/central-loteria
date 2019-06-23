package entities

// User entity
type User struct {
	ID       string `firestore:"id" json:"id"`
	Name     string `firestore:"name" json:"name"`
	Email    string `firestore:"email" json:"email"`
	Password string `firestore:"password" json:"password"`
}
