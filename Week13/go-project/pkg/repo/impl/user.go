package impl

import (
	"context"
	"go-project/pkg/database"
	"go-project/pkg/domain"
	"go-project/pkg/repo/internal/model"

	"github.com/ahmetb/go-linq/v3"
	"github.com/pkg/errors"
)

type UserRepoImpl struct {
	Ds *database.Data
}

func (p *UserRepoImpl) Update(ctx context.Context, in *domain.User) error {
	if err := in.Valid(); err != nil {
		return err
	}

	err := p.Ds.Gdb().Save(model.User{}.New(in)).Error

	return errors.WithStack(err)
}

func (p *UserRepoImpl) Create(ctx context.Context, in *domain.User) error {
	if err := in.Valid(); err != nil {
		return err
	}
	obj := model.User{}.New(in)
	if err := p.Ds.Gdb().Create(obj).Error; err != nil {
		return errors.WithStack(err)
	}
	in.Id = obj.Id
	return nil
}

func (p *UserRepoImpl) MustGet(ctx context.Context, id int) (*domain.User, error) {
	var o model.User
	if err := p.Ds.Gdb().Take(&o, id).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return o.ToDomain(), nil
}

func (p *UserRepoImpl) MultiGet(ctx context.Context, id ...int) (domain.UserList, error) {
	if len(id) == 0 {
		return domain.UserList{}, nil
	}

	var l []*model.User
	if err := p.Ds.Gdb().
		Where("id in (?)", id).
		Find(&l).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	var out domain.UserList
	linq.From(l).SelectT(model.User{}.ModelToDomain).ToSlice(&out)
	return out, nil
}
