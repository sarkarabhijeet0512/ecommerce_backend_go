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
		ReviewID   int
		ProductID  int // Foreign key referencing Product entity
		UserID     int // Foreign key referencing User entity
		Rating     float64
		Comment    string
		ReviewDate time.Time
		IsActive   bool
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}
)
