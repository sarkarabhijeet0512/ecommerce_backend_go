package inventory

import (
	"context"
	model "ecommerce_backend_project/utils/models"
	"sync"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type Repository interface {
	upsertInventory(context.Context, *Inventory) error
	fetchInventoryByFilter(context.Context, model.Filter) ([]Inventory, error)
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
	mu  sync.Mutex
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

func (r *PGRepo) upsertInventory(ctx context.Context, inventory *Inventory) error {
	r.mu.Lock()
	_, err := r.db.ModelContext(ctx, inventory).OnConflict("(product_id) DO UPDATE").Insert()
	r.mu.Unlock()
	return err
}
func (r *PGRepo) fetchInventoryByFilter(ctx context.Context, filter model.Filter) ([]Inventory, error) {
	inventory := []Inventory{}
	query := r.db.ModelContext(ctx, &inventory)
	if filter.ProductID != 0 {
		query.Where("product_id = ?", filter.ProductID)
	}
	_, err := query.Insert()
	return inventory, err
}
