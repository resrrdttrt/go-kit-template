package admin

import (
	"context"
	"time"
)

type User struct {
	ID        string    `db:"id" json:"id,omitempty"`
	Name      string    `db:"name" json:"name"`
	Username  string    `db:"username" json:"username"`
	Password  string    `db:"password" json:"password"`
	Email     string    `db:"email" json:"email"`
	Phone     string    `db:"phone" json:"phone"`
	Role      string    `db:"role" json:"role,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at,omitempty"`
	Status    string    `db:"status" json:"status,omitempty"`
}

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]User, error)
	GetUserById(ctx context.Context, id string) (User, error)
	CreateUser(ctx context.Context, user User) error
	UpdateUser(ctx context.Context, user User) error
	DeleteUser(ctx context.Context, id string) error
}
