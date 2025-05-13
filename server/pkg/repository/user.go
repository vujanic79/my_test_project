package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"github.com/vujanic79/golang-react-todo-app/pkg/internal/database"
	"github.com/vujanic79/golang-react-todo-app/pkg/logger"
	"time"
)

type UserRepository struct {
	Db *database.Queries
}

var _ domain.UserRepository = (*UserRepository)(nil)

func NewUserRepository(db *database.Queries) (ur *UserRepository) {
	return &UserRepository{Db: db}
}

func (ur *UserRepository) CreateUser(ctx context.Context, params domain.CreateUserParams) (u domain.User, err error) {
	l := logger.FromContext(ctx)
	dbU, err := ur.Db.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Email:     params.Email,
	})

	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("repository.params", zerolog.Dict().
				Str("func", "CreateUser").
				Object("params", params)).
			Msg("Creating user error")
		return domain.User{}, err
	}

	return MapDbUserToUser(dbU), err
}

func (ur *UserRepository) GetUserIdByEmail(ctx context.Context, email string) (id uuid.UUID, err error) {
	l := logger.FromContext(ctx)
	id, err = ur.Db.GetUserIdByEmail(ctx, email)

	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("repository.params", zerolog.Dict().
				Str("func", "GetUserUserIdByEmail").
				Dict("params", zerolog.Dict().
					Str("email", email))).
			Msg("Getting user id by email error")
		return uuid.Nil, err
	}

	return id, err
}

func MapDbUserToUser(dbU database.AppUser) (u domain.User) {
	return domain.User{
		ID:        dbU.ID,
		CreatedAt: dbU.CreatedAt,
		UpdatedAt: dbU.UpdatedAt,
		FirstName: dbU.FirstName,
		LastName:  dbU.LastName,
		Email:     dbU.Email,
	}
}
