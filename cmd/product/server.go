package main

import (
	config "ecommerce_backend_project/config/product"
	productServer "ecommerce_backend_project/internal/services/product"
	"ecommerce_backend_project/internal/services/product/handler"
	productDB "ecommerce_backend_project/utils/db/product"
	"ecommerce_backend_project/utils/initialize"

	"go.uber.org/fx"
)

func serverRun() {
	app := fx.New(
		fx.Provide(
			// postgres server
			productDB.NewDB,
			initialize.NewRedisWorker,
		),
		config.Module,
		initialize.Module,
		productServer.Module,
		handler.Module,
	)

	// Run app forever
	app.Run()
}
