package main

import (
	config "ecommerce_backend_project/config/auth"
	authServer "ecommerce_backend_project/internal/services/auth"
	"ecommerce_backend_project/internal/services/auth/handler"
	"ecommerce_backend_project/pkg/auth/payment"
	"ecommerce_backend_project/pkg/auth/rbac"
	"ecommerce_backend_project/pkg/auth/user"
	authDB "ecommerce_backend_project/utils/db/auth"
	"ecommerce_backend_project/utils/initialize"

	"go.uber.org/fx"
)

func serverRun() {
	app := fx.New(
		fx.Provide(
			// postgres server
			authDB.NewDB,
			initialize.NewRedisWorker,
		),
		config.Module,
		initialize.Module,
		authServer.Module,
		handler.Module,
		user.Module,
		rbac.Module,
		payment.Module,
	)

	// Run app forever
	app.Run()
}
