package orderdetails

import (
	"ecommerce_backend_project/pkg/order/orderitems"
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
	Order struct {
		ID                int                    `json:"id" pg:"id"`
		UserID            int                    `json:"user_id" pg:"user_id"`
		OrderDate         time.Time              `json:"order_date" pg:"order_date"`
		TotalAmount       float64                `json:"total_amount" pg:"total_amount"`
		OrderStatus       string                 `json:"order_status" pg:"order_status"`
		OrderItemsID      int                    `json:"order_items_id" pg:"order_items_id"`
		OrderItems        []orderitems.OrderItem `json:"order_items" pg:"order_items"`
		ShippingAddressID int                    `json:"shipping_address_id" pg:"shipping_address"`
		BillingAddressID  int                    `json:"billing_address_id" pg:"billing_address"`
		PaymentID         int                    `json:"payment_id" pg:"payment_id"`
		DeliveryDate      time.Time              `json:"delivery_date" pg:"delivery_date"`
		IsPaid            bool                   `json:"is_paid" pg:"is_paid"`
		IsDelivered       bool                   `json:"is_delivered" pg:"is_delivered"`
		IsActive          bool                   `json:"is_active" pg:"is_active"`
		CreatedAt         time.Time              `json:"created_at" pg:"created_at"`
		UpdatedAt         time.Time              `json:"updated_at" pg:"updated_at"`
	}
)
