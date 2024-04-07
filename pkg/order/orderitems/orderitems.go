package dummy

import (
	"go.uber.org/fx"
)

// Module provides all constructor and invocation methods to facilitate credits module
var Module = fx.Options(
	fx.Provide(
		NewDBRepository,
		NewService,
	),
)
