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
		InventoryID       int
		ProductID         int // Foreign key referencing Product entity
		StockLevel        int
		ReorderThreshold  int // Threshold at which inventory needs to be reordered
		ReorderQuantity   int // Quantity to reorder when inventory falls below threshold
		LastRestockedDate time.Time
	}
)
