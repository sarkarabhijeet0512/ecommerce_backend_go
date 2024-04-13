package server

import (
	"ecommerce_backend_project/internal/mw"
	"ecommerce_backend_project/internal/mw/jwt"

	"github.com/gin-gonic/gin"
)

func v1Routes(router *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware, o *Options) {
	r := router.Group("/v1/order/api")
	r.Use(mw.ErrorHandlerX(o.Log))
	r.PUT("/order_details", authMiddleware.MiddlewareFunc())
	r.GET("/order_details", authMiddleware.MiddlewareFunc())

}
