package repo

import (
	"context"
	"go-project/pkg/domain"
)

//go:generate mockgen -destination=../mocks/mrepo/user.go -package=mrepo . UserRepo

type UserRepo interface {
	MustGet(ctx context.Context, id int) (*domain.User, error)
	MultiGet(ctx context.Context, id ...int) (domain.UserList, error)
	Create(ctx context.Context, in *domain.User) error
	Update(ctx context.Context, in *domain.User) error
}
