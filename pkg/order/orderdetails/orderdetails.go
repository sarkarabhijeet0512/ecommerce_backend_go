package orderdetails

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
	// Order represents the order entity
	Order struct {
		OrderID           int
		UserID            int
		OrderDate         time.Time
		TotalAmount       float64
		OrderStatus       string
		ShippingAddressID int
		BillingAddressID  int
		PaymentID         int
		DeliveryDate      time.Time
		IsPaid            bool
		IsDelivered       bool
		IsActive          bool
		CreatedAt         time.Time
		UpdatedAt         time.Time
	}
)
