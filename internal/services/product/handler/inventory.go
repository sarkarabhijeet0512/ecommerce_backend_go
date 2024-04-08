package handler

import (
	"ecommerce_backend_project/internal/mw/jwt"
	"ecommerce_backend_project/pkg/product/inventory"
	productdetails "ecommerce_backend_project/pkg/product/productDetails"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type InventoryHandler struct {
	log                   *logrus.Logger
	jwtMiddleware         *jwt.GinJWTMiddleware
	invertoryService      *inventory.Service
	productdetailsService *productdetails.Service
}

func newInventoryHandler(
	log *logrus.Logger,
	invertoryService *inventory.Service,
	productdetailsService *productdetails.Service,
) *InventoryHandler {
	c := &gin.Context{}
	return &InventoryHandler{
		log,
		jwt.SetAuthMiddleware(productdetailsService.Repo.GetDBConnection(c)),
		invertoryService,
		productdetailsService,
	}
}
