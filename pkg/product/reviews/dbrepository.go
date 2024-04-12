package reviews

import (
	"context"
	model "ecommerce_backend_project/utils/models"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type Repository interface {
	updateReviewByProductID(context.Context, *Review) error
	fetchReviewByFilter(context.Context, *model.Filter) ([]Review, error)
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

func (r *PGRepo) updateReviewByProductID(ctx context.Context, review *Review) error {
	_, err := r.db.ModelContext(ctx, review).OnConflict("(user_id, product_id) DO UPDATE").Insert()
	return err
}
func (r *PGRepo) fetchReviewByFilter(ctx context.Context, filter *model.Filter) ([]Review, error) {
	reviews := []Review{}
	query := r.db.ModelContext(ctx, &reviews)
	if filter != nil && filter.ProductID != 0 {
		query.Where("product_id = ?", filter.ProductID)
	}
	if filter != nil && filter.UserID != 0 {
		query.Where("user_id = ?", filter.UserID)
	}
	err := query.Select()
	return reviews, err
}
