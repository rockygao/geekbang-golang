package data

import (
	"Week4/internal/biz"
	//"log"
	"context"
)

type userRepo struct {
	data *Data
}

// NewArticleRepo .
func NewArticleRepo(data *Data) biz.UserRepo {
	return &userRepo{
		data: data,
	}
}

func (ar *userRepo) CreateUser(ctx context.Context, user *biz.User) error {
	//_, err := ar.data.db.Use
	//	Create()
	//	Create().
	//	SetTitle(user.username).
	//	SetContent(user.pwssword).
	//	Save(ctx)
	return ctx
}
