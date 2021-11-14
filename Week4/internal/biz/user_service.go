package biz

import (
	"context"
	"go_code/geekbang-golang/Week4/internal/data"
	"go_code/geekbang-golang/Week4/internal/model"

	"github.com/pkg/errors"
)

type UserService struct {
	userDao data.UserDaoInf
}

func NewUserService(dao data.UserDaoInf) *UserService {
	return &UserService{
		userDao: dao,
	}
}

func (service *UserService) Register(ctx context.Context, user model.User) (bool, error) {
	return service.userDao.Create(user)
}

func (service *UserService) Login(ctx context.Context, username string, password string) (bool, error) {
	user, err := service.userDao.QueryOne(username)
	if err != nil {
		if errors.Is(err, data.ErrNotExit) {
			return false, errors.Wrap(err, "user not exit")
		} else {
			return false, err
		}
	}
	return user.Password == password, nil
}
