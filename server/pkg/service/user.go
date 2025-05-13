package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"github.com/vujanic79/golang-react-todo-app/pkg/logger"
)

type UserService struct {
	Ur domain.UserRepository
}

var _ domain.UserService = (*UserService)(nil)

func NewUserService(ur domain.UserRepository) (us *UserService) {
	return &UserService{Ur: ur}
}

func (us *UserService) CreateUser(ctx context.Context, params domain.CreateUserParams) (u domain.User, err error) {
	l := logger.FromContext(ctx)

	l = l.With().
		Dict("service.params", zerolog.Dict().
			Str("func", "CreateUser").
			Object("params", params)).
		Logger()
	ctx = logger.WithLogger(ctx, l)

	u, err = us.Ur.CreateUser(ctx, params)
	return u, err
}

func (us *UserService) GetUserIdByEmail(ctx context.Context, email string) (id uuid.UUID, err error) {
	l := logger.FromContext(ctx)

	l = l.With().
		Dict("service.params", zerolog.Dict().
			Str("func", "GetUserUserIdByEmail").
			Dict("params", zerolog.Dict().
				Str("email", email))).
		Logger()
	ctx = logger.WithLogger(ctx, l)

	id, err = us.Ur.GetUserIdByEmail(ctx, email)
	return id, err
}
