package inventory

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
	// Inventory represents the inventory of a product
	Inventory struct {
		ID                int       `json:"id" pg:"id,pk"`
		ProductID         int       `json:"product_id" pg:"product_id"`
		StockLevel        int       `json:"stock_level" pg:"stock_level"`
		ReorderThreshold  int       `json:"reorder_threshold" pg:"reorder_threshold"`
		ReorderQuantity   int       `json:"reorder_quantity" pg:"reorder_quantity"`
		LastRestockedDate time.Time `json:"last_restocked_date" pg:"last_restocked"`
		IsActive          bool      `json:"is_active" pg:"is_active"`
		CreatedAt         time.Time `json:"created_at" pg:"created_at"`
		UpdatedAt         time.Time `json:"updated_at" pg:"updated_at"`
	}
)
