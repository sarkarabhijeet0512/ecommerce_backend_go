package server

import (
	"ecommerce_backend_project/internal/mw"
	"ecommerce_backend_project/internal/mw/jwt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
)

func v1Routes(router *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware, awsSession *session.Session, o *Options) {
	r := router.Group("/v1/product/api/")
	r.Use(mw.ErrorHandlerX(o.Log))
	// add new routes here
	r.PUT("/product_details", authMiddleware.MiddlewareFunc(), o.ProductDetailsHandler.UpsertProductDetails)
	r.GET("/product_details/:product_id", authMiddleware.MiddlewareFunc(), o.ProductDetailsHandler.GetProductDetails)
	r.GET("/catergory/:category_id", authMiddleware.MiddlewareFunc(), o.ProductDetailsHandler.GetProductListByCategory)
	r.GET("/catergory", authMiddleware.MiddlewareFunc(), o.ProductDetailsHandler.GetProductCategory)
	// r.PUT("/disable/products", authMiddleware.MiddlewareFunc(), o.ProductDetailsHandler.DisableProduct)
	// r.GET("/products", authMiddleware.MiddlewareFunc(), o.ProductDetailsHandler.GetProductList)
	r.POST("/upload/product_images", mw.AWSSessionAttach(awsSession), authMiddleware.MiddlewareFunc(), o.ProductDetailsHandler.UploadProductImages)
}
