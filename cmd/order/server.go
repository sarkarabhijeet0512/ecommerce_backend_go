package main

import (
	orderServer "ecommerce_backend_project/internal/services/order"
	"ecommerce_backend_project/internal/services/order/handler"
	orderDB "ecommerce_backend_project/utils/db/order"
	"ecommerce_backend_project/utils/initialize"
	"ecommerce_backend_project/utils/kafka"

	"go.uber.org/fx"
)

func serverRun() {
	app := fx.New(
		fx.Provide(
			// postgres server
			orderDB.NewDB,
			kafka.NewKafkaProducer,
			kafka.NewKafkaConsumer,
		),
		// config.Module,
		initialize.Module,
		orderServer.Module,
		handler.Module,
	)

	// Run app forever
	app.Run()
}
