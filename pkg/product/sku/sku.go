package sku

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
		ProductID        int
		ProductName      string
		Description      string
		Price            float64
		Quantity         int
		CategoryID       int
		Brand            string
		Manufacturer     string
		Weight           float64
		Dimensions       Dimensions
		SKU              string
		Tags             []string
		Rating           float64
		Reviews          []reviews.Review
		Variants         []ProductVariant
		SupplierID       int                       // Foreign key referencing Supplier entity
		RelatedProducts  []Product                 // Related products
		Discounts        []offermangement.Discount // Discounts applicable to the product
		CreatedDate      time.Time
		LastModifiedDate time.Time
		IsActive         bool
		CreatedAt        time.Time
		UpdatedAt        time.Time
	}
	// ProductImage represents the product image entity
	ProductImage struct {
		ImageID      int
		ProductID    int
		ImageURL     string
		AltText      string
		IsPrimary    bool
		Order        int
		ThumbnailURL string
		IsActive     bool
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}
	// Dimensions represents the dimensions of a product
	Dimensions struct {
		Length    float64
		Width     float64
		Height    float64
		IsActive  bool
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	// ProductVariant represents a variant of a product (e.g., size, color)
	ProductVariant struct {
		VariantID   int
		ProductID   int // Foreign key referencing Product entity
		Name        string
		Description string
		SKU         string
		Price       float64
		Quantity    int
		IsActive    bool
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
)
