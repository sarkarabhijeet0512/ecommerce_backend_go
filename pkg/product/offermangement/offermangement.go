package offermangement

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
	// Discount represents a discount applicable to a product
	Discount struct {
		ID            int
		ProductID     int
		DiscountType  string
		CouponCode    string
		DiscountValue float64
		IsActive      bool
		CreatedAt     time.Time
		UpdatedAt     time.Time
	}
)
