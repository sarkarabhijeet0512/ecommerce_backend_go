package offermangement

import (
	"context"
	model "ecommerce_backend_project/utils/models"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type Repository interface {
	upsertOfferDetails(ctx context.Context, discount *Discount) error
	fetchOfferByFilter(ctx context.Context, filter model.Filter) ([]Discount, error)
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
func (r *PGRepo) upsertOfferDetails(ctx context.Context, discount *Discount) error {
	_, err := r.db.ModelContext(ctx, discount).OnConflict("(coupon_code) DO UPDATE").Insert()
	return err
}
func (r *PGRepo) fetchOfferByFilter(ctx context.Context, filter model.Filter) ([]Discount, error) {
	discount := []Discount{}
	query := r.db.ModelContext(ctx, &discount)
	if filter.ProductID != 0 {
		query.Where("product_id=?", filter.ProductID)
	}
	if filter.CouponCode != "" {
		query.Where("coupon_code=?", filter.CouponCode)
	}
	query.Select()
	return nil, nil
}
