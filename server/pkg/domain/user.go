package domain

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"net/http"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
}

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func (cup CreateUserParams) MarshalZerologObject(e *zerolog.Event) {
	e.
		Str("first_name", cup.FirstName).
		Str("last_name", cup.LastName).
		Str("email", cup.Email)
}

type UserService interface {
	CreateUser(ctx context.Context, params CreateUserParams) (u User, err error)
	GetUserIdByEmail(ctx context.Context, email string) (id uuid.UUID, err error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, params CreateUserParams) (u User, err error)
	GetUserIdByEmail(ctx context.Context, email string) (id uuid.UUID, err error)
}

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}
