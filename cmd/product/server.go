package main

import (
	product_server "ecommerce_backend_project/internal/services/product"
	"ecommerce_backend_project/internal/services/product/handler"
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
		product_server.Module,
		handler.Module,
	)

	// Run app forever
	app.Run()
}
