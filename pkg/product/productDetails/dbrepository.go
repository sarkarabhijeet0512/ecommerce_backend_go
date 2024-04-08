package productdetails

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type Repository interface {
	GetDBConnection(context.Context) *pg.DB
	upsertProductDetails(context.Context, Product) error
	getProductByID(context.Context, Product) error
	getProductListByCategory(context.Context, Product) error
	getProductList(context.Context, Product) error
	disableProductByID(context.Context, Product) error
}

// NewRepositoryIn is function param struct of func `NewRepository`
type NewRepositoryIn struct {
	fx.In

	Log *logrus.Logger
	DB  *pg.DB `name:"productdb"`
}

// PGRepo is postgres implementation
type PGRepo struct {
	log *logrus.Logger
	db  *pg.DB
}

// NewDBRepository returns a new persistence layer object which can be used for
// CRUD on db
func NewDBRepository(i NewRepositoryIn) (Repo Repository, err error) {

	Repo = &PGRepo{
		log: i.Log,
		db:  i.DB,
	}

	return
}

func (r *PGRepo) GetDBConnection(dCtx context.Context) *pg.DB {
	return r.db
}

func (r *PGRepo) upsertProductDetails(dCtx context.Context, req Product) error {
	return nil
}

func (r *PGRepo) getProductByID(dCtx context.Context, req Product) error {
	return nil
}

func (r *PGRepo) getProductListByCategory(dCtx context.Context, req Product) error {
	return nil
}

func (r *PGRepo) getProductList(dCtx context.Context, req Product) error {
	return nil
}

func (r *PGRepo) disableProductByID(dCtx context.Context, req Product) error {
	return nil
}
