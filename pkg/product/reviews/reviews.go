package reviews

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
	// Review represents a user review for a product
	Review struct {
		ID         int       `json:"id" pg:"id"`
		ProductID  int       `json:"product_id" pg:"product_id"`
		UserID     int       `json:"user_id" pg:"user_id"`
		Rating     float64   `json:"rating" pg:"rating"`
		Comment    string    `json:"comment" pg:"comment"`
		ReviewDate time.Time `json:"review_date" pg:"review_date"`
		ExpiryDate time.Time `json:"expiry_date" pg:"expiry_date"`
		IsActive   bool      `json:"is_active" pg:"is_active"`
		CreatedAt  time.Time `json:"created_at" pg:"created_at"`
		UpdatedAt  time.Time `json:"updated_at" pg:"updated_at"`
	}
)
