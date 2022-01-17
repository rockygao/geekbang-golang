package repo

import (
	"context"

	"go-project/pkg/repo/internal/model"

	"go-project/pkg/database"

	"github.com/pkg/errors"
)

type DbInitializer struct {
	Ds *database.Data
}

func (p *DbInitializer) Name() string {
	return "db_initializer"
}

func (p *DbInitializer) IsNeedInit(ctx context.Context) (bool, error) {
	return true, nil
}

// Initialize AutoMigrate自动建表
func (p *DbInitializer) Initialize(ctx context.Context) error {
	err := p.Ds.Gdb().AutoMigrate(
		&model.User{},
	)
	return errors.WithStack(err)
}
