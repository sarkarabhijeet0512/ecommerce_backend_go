package orderitems

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
	OrderItem struct {
		ID        int       `json:"id" pg:"id"`
		OrderID   int       `json:"order_id" pg:"order_id"`
		ProductID string    `json:"product_id" pg:"product_id"`
		Quantity  int       `json:"quantity" pg:"quantity"`
		Price     float64   `json:"price" pg:"price"`
		IsActive  bool      `json:"is_active" pg:"is_active"`
		CreatedAt time.Time `json:"created_at" pg:"created_at"`
		UpdatedAt time.Time `json:"updated_at" pg:"updated_at"`
	}
)
