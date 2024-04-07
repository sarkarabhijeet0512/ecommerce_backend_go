package main

import (
	order_server "ecommerce_backend_project/internal/services/order"
	"ecommerce_backend_project/internal/services/order/handler"
	"ecommerce_backend_project/utils/initialize"

	"go.uber.org/fx"
)

func serverRun() {
	app := fx.New(
		fx.Provide(
			// postgres server
			initialize.NewDB,
		),
		// config.Module,
		initialize.Module,
		order_server.Module,
		handler.Module,
	)

	// Run app forever
	app.Run()
}
