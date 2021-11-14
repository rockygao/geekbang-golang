package data

import (
	"errors"
	"go_code/geekbang-golang/Week4/internal/model"
)

var (
	ErrNotExit    = errors.New("no such record")
	ErrAleadyExit = errors.New("recored with same id already exist")
)

type UserDaoInf interface {
	Create(user model.User) (bool, error)
	QueryOne(username string) (model.User, error)
}

type InMemoUserDao struct {
	user []model.User
}

func NewMemoUserDao() UserDaoInf {
	return &InMemoUserDao{
		users: make([]model.User, 0),
	}
}

func (dao *InMemoUserDao) Create(user model.User) (bool, error) {
	for _, u := range dao.users {
		if u.UserName == user.UserName {
			return false, ErrAleadyExit
		}
	}
	dao.users = append(dao.users, user)
	return true, nil

}

func (dao *InMemoUserDao) QueryOne(username string) (model.User, error) {
	for _, u := range dao.users {
		if u.UserName == username {
			return u, nil
		}
	}
	return model.User{}, ErrAleadyExit
}
