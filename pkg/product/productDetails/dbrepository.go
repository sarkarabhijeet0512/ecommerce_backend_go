package productdetails

import (
	"context"
	"ecommerce_backend_project/utils"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type Repository interface {
	GetDBConnection(context.Context) *pg.DB
	upsertProductDetails(context.Context, *Product) error
	upsertDimentions(context.Context, *Dimensions) error
	upsertProductVariants(context.Context, []ProductVariant) error
	getProductByID(context.Context, int) (*Product, error)
	getProductListByCategory(context.Context, int) ([]Product, error)
	getProductCategoryList(context.Context) ([]Category, error)
	disableProductByID(context.Context, Product) error
	productImageDetails(context.Context, *ProductImage) error
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

func (r *PGRepo) GetDBConnection(ctx context.Context) *pg.DB {
	return r.db
}

func (r *PGRepo) upsertProductDetails(ctx context.Context, product *Product) error {
	utils.SetGenericFieldValue(product)
	_, err := r.db.ModelContext(ctx, product).OnConflict("(sku) DO UPDATE").Insert()
	return err
}

func (r *PGRepo) getProductByID(ctx context.Context, productID int) (*Product, error) {
	product := &Product{}
	err := r.db.Model(product).
		// Relation("Images").
		Relation("Variants").
		Relation("Category").
		// Relation("Reviews").
		// Relation("Supplier").
		// Relation("Discounts").
		Where("product.id = ?", productID).
		Select()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *PGRepo) getProductListByCategory(ctx context.Context, categoryID int) ([]Product, error) {
	product := []Product{}
	err := r.db.Model(&product).
		// Relation("Images").
		Relation("Variants").
		Relation("Category").
		// Relation("Reviews").
		// Relation("Supplier").
		// Relation("Discounts").
		Where("product.category_id = ?", categoryID).
		Select()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *PGRepo) getProductCategoryList(ctx context.Context) ([]Category, error) {
	category := []Category{}
	err := r.db.Model(&category).Select()
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (r *PGRepo) disableProductByID(ctx context.Context, product Product) error {
	return nil
}

func (r *PGRepo) productImageDetails(ctx context.Context, productImage *ProductImage) error {
	utils.SetGenericFieldValue(productImage)
	_, err := r.db.ModelContext(ctx, productImage).Insert()
	return err
}

func (r *PGRepo) upsertDimentions(ctx context.Context, dimensions *Dimensions) error {
	utils.SetGenericFieldValue(dimensions)
	_, err := r.db.ModelContext(ctx, dimensions).OnConflict("(product_id) DO UPDATE").Insert()
	return err
}
func (r *PGRepo) upsertProductVariants(ctx context.Context, productVariant []ProductVariant) error {
	_, err := r.db.ModelContext(ctx, &productVariant).OnConflict("(product_id) DO UPDATE").Insert()
	return err
}
