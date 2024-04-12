package suppliers

import (
	"context"
	model "ecommerce_backend_project/utils/models"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type Repository interface {
	upsertSuppliers(ctx context.Context, suppliers *Supplier) error
	fetchSuppliers(ctx context.Context, Filter model.Filter) (suppliers []Supplier, err error)
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

// IsActive checks if DB is connected
func (r *PGRepo) upsertSuppliers(ctx context.Context, suppliers *Supplier) error {
	_, err := r.db.ModelContext(ctx, suppliers).OnConflict("(supplier_name) DO UPDATE").Insert()
	return err
}

func (r *PGRepo) fetchSuppliers(ctx context.Context, Filter model.Filter) ([]Supplier, error) {
	suppliers := []Supplier{}
	err := r.db.ModelContext(ctx, &suppliers).Where("id=?", Filter.SupplierID).Select()
	return suppliers, err
}
