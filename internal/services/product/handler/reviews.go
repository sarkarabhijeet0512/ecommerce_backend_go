package handler

import (
	"ecommerce_backend_project/internal/mw/jwt"
	productdetails "ecommerce_backend_project/pkg/product/productDetails"
	"ecommerce_backend_project/pkg/product/reviews"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ReviewsHandler struct {
	log                   *logrus.Logger
	jwtMiddleware         *jwt.GinJWTMiddleware
	reviewService         *reviews.Service
	productdetailsService *productdetails.Service
}

func newReviewsHandler(
	log *logrus.Logger,
	reviewService *reviews.Service,
	productdetailsService *productdetails.Service,
) *ReviewsHandler {
	c := &gin.Context{}
	return &ReviewsHandler{
		log,
		jwt.SetAuthMiddleware(productdetailsService.Repo.GetDBConnection(c)),
		reviewService,
		productdetailsService,
	}
}
