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
		ID           int `json:"id" pg:"id,pk"`
		SupplierName string
		IsActive     bool
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}
)
