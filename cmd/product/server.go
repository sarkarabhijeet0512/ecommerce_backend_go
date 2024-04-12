package main

import (
	config "ecommerce_backend_project/config/product"
	productServer "ecommerce_backend_project/internal/services/product"
	"ecommerce_backend_project/internal/services/product/handler"
	"ecommerce_backend_project/pkg/product/inventory"
	"ecommerce_backend_project/pkg/product/offermangement"
	productdetails "ecommerce_backend_project/pkg/product/productDetails"
	"ecommerce_backend_project/pkg/product/reviews"
	"ecommerce_backend_project/pkg/product/suppliers"
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
		inventory.Module,
		productdetails.Module,
		reviews.Module,
		inventory.Module,
		suppliers.Module,
		offermangement.Module,
	)

	// Run app forever
	app.Run()
}
