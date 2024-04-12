package server

import (
	"ecommerce_backend_project/internal/mw"
	"ecommerce_backend_project/internal/mw/jwt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
)

func v1Routes(router *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware, awsSession *session.Session, o *Options) {
	r := router.Group("/v1/product/api/")
	// add new routes here
	r.PUT("/product_details", authMiddleware.MiddlewareFunc(), o.ProductDetailsHandler.UpsertProductDetails)
	r.GET("/product_details/:product_id", authMiddleware.MiddlewareFunc(), o.ProductDetailsHandler.GetProductDetails)

	r.GET("/catergory/:category_id", authMiddleware.MiddlewareFunc(), o.ProductDetailsHandler.GetProductListByCategory)
	r.GET("/catergory", authMiddleware.MiddlewareFunc(), o.ProductDetailsHandler.GetProductCategory)

	r.POST("/upload/product_images", mw.AWSSessionAttach(awsSession), authMiddleware.MiddlewareFunc(), o.ProductDetailsHandler.UploadProductImages)

	// Reviews routes
	r.POST("/review/review_comments", authMiddleware.MiddlewareFunc(), o.ReviewsHandler.UpsertReviewByProductID)
	r.GET("/review", authMiddleware.MiddlewareFunc(), o.ReviewsHandler.FetchReview)

	r.PUT("/offer_details", authMiddleware.MiddlewareFunc(), o.OfferHandler.UpsertOfferDetails)
	r.GET("/offer_details", authMiddleware.MiddlewareFunc(), o.OfferHandler.FetchOfferDetails)

	r.PUT("/supplier_details", authMiddleware.MiddlewareFunc(), o.SupplierHandler.UpsertSuppliers)
	r.GET("/supplier_details", authMiddleware.MiddlewareFunc(), o.SupplierHandler.FetchSuppliers)

	r.PUT("/inventory_details", authMiddleware.MiddlewareFunc(), o.InventoryHandler.UpsertInventory)
	r.POST("/inventory_details", authMiddleware.MiddlewareFunc(), o.InventoryHandler.FetchInventory)
}
