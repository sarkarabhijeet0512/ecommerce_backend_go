package suppliers

import (
	"time"

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
	// Supplier represents a supplier of products
	Supplier struct {
		ID           int       `json:"id" pg:"id,pk"`
		SupplierName string    `json:"supplier_name" pg:"supplier_name"`
		IsActive     bool      `json:"is_active" pg:"is_active"`
		CreatedAt    time.Time `json:"created_at" pg:"created_at"`
		UpdatedAt    time.Time `json:"updated_at" pg:"updated_at"`
	}
)
