package types

import "time"

// User is the shared user record used across all services.
type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// UserRepository is the interface both API and worker use.
// Neither knows the DB implementation, they depend on this abstraction.
type UserRepository interface {
	GetUser(id string) (*User, error)
	ListUsers() ([]*User, error)
	CreateUser(u *User) error
}
