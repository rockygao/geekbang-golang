package biz

import "context"

type User struct {
	username string
	password string
}

type UserRepo interface {
	CreateUser(ctx context.Context,user *User) error
}

type UserUsecase struct {
	repo UserRepo
}

func (uc *UserUsecase) Create(ctx context.Context, user *User) error {
	return uc.repo.CreateUser(ctx, user)
}
