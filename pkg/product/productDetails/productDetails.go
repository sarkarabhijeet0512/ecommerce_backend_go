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
		ID           int              `json:"id" pg:"id,pk"`
		ProductName  string           `json:"product_name" pg:"product_name"`
		Description  string           `json:"description" pg:"description"`
		Price        float64          `json:"price" pg:"price"`
		Quantity     int              `json:"quantity" pg:"quantity"`
		CategoryID   int              `json:"category_id" pg:"category_id"`
		Category     *Category        `json:"category" pg:"rel:has-one"` // Relationship to Category with foreign key
		Brand        string           `json:"brand" pg:"brand"`
		Manufacturer string           `json:"manufacturer" pg:"manufacturer"`
		Weight       float64          `json:"weight" pg:"weight"`
		Dimensions   *Dimensions      `json:"dimensions" pg:"rel:has-one"`
		SKU          string           `json:"sku" pg:"sku,unique"`
		Tags         []string         `json:"tags" pg:"tags"`
		Rating       float64          `json:"rating" pg:"rating"`
		Reviews      []reviews.Review `json:"reviews" pg:"rel:has-many"`
		Variants     []ProductVariant `json:"product_variants" pg:"has-many"`
		SupplierID   int              `json:"supplier_id" pg:"supplier_id"`
		// Supplier          suppliers.Supplier        `json:"supplier" pg:"rel:has-one"`
		RelatedProducts   []Product                 `pg:"-"`
		Discounts         []offermangement.Discount `json:"discounts" pg:"rel:has-many"`
		ManufacturingDate time.Time                 `json:"manufacturing_date" pg:"manufacturing_date"`
		ExpiryDate        time.Time                 `json:"expiry_date" pg:"expiry_date"`
		IsActive          bool                      `json:"is_active" pg:"is_active,default:true"`
		CreatedAt         time.Time                 `json:"created_at" pg:"created_at"`
		UpdatedAt         time.Time                 `json:"updated_at" pg:"updated_at"`
	}
	// ProductImage represents the product image entity
	ProductImage struct {
		ID         int       `json:"id" pg:"id"`
		UploadedBy int       `json:"uploaded_by" pg:"uploaded_by"`
		ProductID  int       `json:"product_id" pg:"product_id"`
		ImageURL   string    `json:"image_url" pg:"image_url"`
		AltText    string    `json:"alt_text" pg:"alt_text"`
		IsPrimary  bool      `json:"is_primary" pg:"is_primary"`
		Order      int       `json:"order" pg:"order"`
		IsActive   bool      `json:"is_active" pg:"is_active"`
		CreatedAt  time.Time `json:"created_at" pg:"created_at"`
		UpdatedAt  time.Time `json:"updated_at" pg:"updated_at"`
	}
	// Dimensions represents the dimensions of a product
	Dimensions struct {
		tableName struct{}  `pg:"dimensions,discard_unknown_columns"`
		ID        int       `json:"id" pg:"id"`
		ProductID int       `json:"product_id" pg:"product_id,unique"`
		Length    float64   `json:"length" pg:"length"`
		Width     float64   `json:"width" pg:"width"`
		Height    float64   `json:"height" pg:"height"`
		IsActive  bool      `json:"is_active" pg:"is_active"`
		CreatedAt time.Time `json:"created_at" pg:"created_at"`
		UpdatedAt time.Time `json:"updated_at" pg:"updated_at"`
	}

	// ProductVariant represents a variant of a product (e.g., size, color)
	ProductVariant struct {
		tableName   struct{}  `pg:"product_variants,discard_unknown_columns"`
		ID          int       `json:"id" pg:"id"`
		ProductID   int       `json:"product_id" pg:"product_id,unique"`
		Name        string    `json:"name" pg:"name"`
		Description string    `json:"description" pg:"description"`
		SKU         string    `json:"sku" pg:"sku"`
		Price       float64   `json:"price" pg:"price"`
		Quantity    int       `json:"quantity" pg:"quantity"`
		IsActive    bool      `json:"is_active" pg:"is_active,default:true"`
		CreatedAt   time.Time `json:"created_at" pg:"created_at"`
		UpdatedAt   time.Time `json:"updated_at" pg:"updated_at"`
	}
	Category struct {
		tableName   struct{}  `pg:"categories,discard_unknown_columns"`
		ID          int       `json:"id" pg:"id,pk"`
		Name        string    `json:"name" pg:"name"`
		Description string    `json:"description" pg:"description"`
		ParentID    int       `json:"parent_id,omitempty" pg:"parent_id"`
		IsActive    bool      `json:"is_active" pg:"is_active"`
		CreatedAt   time.Time `json:"created_at" pg:"created_at"`
		UpdatedAt   time.Time `json:"updated_at" pg:"updated_at"`
	}
)
