package productdetails

import (
	"time"

	"ecommerce_backend_project/pkg/product/offermangement"
	"ecommerce_backend_project/pkg/product/reviews"

	"go.uber.org/fx"
)

// Module provides all constructor and invocation methods to facilitate credits module
var Module = fx.Options(
	fx.Provide(
		NewDBRepository,
		NewService,
	),
)

type (
	// Product represents the product entity
	Product struct {
		ID                int                       `json:"id" pg:"id,pk"`
		ProductName       string                    `json:"product_name" pg:"product_name"`
		Description       string                    `json:"description" pg:"description"`
		Price             float64                   `json:"price" pg:"price"`
		Quantity          int                       `json:"quantity" pg:"quantity"`
		CategoryID        int                       `json:"category_id" pg:"category_id"`
		Brand             string                    `json:"brand" pg:"brand"`
		Manufacturer      string                    `json:"manufacturer" pg:"manufacturer"`
		Weight            float64                   `json:"weight" pg:"weight"`
		Dimensions        Dimensions                `pg:"-"`
		SKU               string                    `json:"sku" pg:"sku"`
		Tags              []string                  `json:"tags" pg:"tags"`
		Rating            float64                   `json:"rating" pg:"rating"`
		Reviews           []reviews.Review          `pg:"-"`
		Variants          []ProductVariant          `pg:"-"`
		SupplierID        int                       `json:"supplier_id" pg:"supplier_id"`
		RelatedProducts   []Product                 `pg:"-"`
		Discounts         []offermangement.Discount `pg:"-"`
		ManufacturingDate time.Time                 `json:"manufacturing_date" pg:"manufacturing_date"`
		LastModifiedDate  time.Time                 `json:"last_modified_date" pg:"last_modified_date"`
		ExpiryDate        time.Time                 `json:"expiry_date" pg:"expiry_date"`
		IsActive          bool                      `json:"is_active" pg:"is_active"`
		CreatedAt         time.Time                 `json:"created_at" pg:"created_at"`
		UpdatedAt         time.Time                 `json:"updated_at" pg:"updated_at"`
	}
	// ProductImage represents the product image entity
	ProductImage struct {
		ID        int       `json:"id" pg:"id"`
		ProductID int       `json:"product_id" pg:"product_id"`
		ImageURL  string    `json:"image_url" pg:"image_url"`
		AltText   string    `json:"alt_text" pg:"alt_text"`
		IsPrimary bool      `json:"is_primary" pg:"is_primary"`
		Order     int       `json:"order" pg:"order"`
		IsActive  bool      `json:"is_active" pg:"is_active"`
		CreatedAt time.Time `json:"created_at" pg:"created_at"`
		UpdatedAt time.Time `json:"updated_at" pg:"updated_at"`
	}
	// Dimensions represents the dimensions of a product
	Dimensions struct {
		ID        int       `json:"id" pg:"id"`
		ProductID int       `json:"product_id" pg:"product_id"`
		Length    float64   `json:"length" pg:"length"`
		Width     float64   `json:"width" pg:"width"`
		Height    float64   `json:"height" pg:"height"`
		IsActive  bool      `json:"is_active" pg:"is_active"`
		CreatedAt time.Time `json:"created_at" pg:"created_at"`
		UpdatedAt time.Time `json:"updated_at" pg:"updated_at"`
	}

	// ProductVariant represents a variant of a product (e.g., size, color)
	ProductVariant struct {
		ID          int       `json:"id" pg:"id"`
		ProductID   int       `json:"product_id" pg:"product_id"`
		Name        string    `json:"name" pg:"name"`
		Description string    `json:"description" pg:"description"`
		SKU         string    `json:"sku" pg:"sku"`
		Price       float64   `json:"price" pg:"price"`
		Quantity    int       `json:"quantity" pg:"quantity"`
		IsActive    bool      `json:"is_active" pg:"is_active"`
		CreatedAt   time.Time `json:"created_at" pg:"created_at"`
		UpdatedAt   time.Time `json:"updated_at" pg:"updated_at"`
	}
)
