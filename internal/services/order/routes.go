package server

import (
	"ecommerce_backend_project/internal/mw"

	"github.com/gin-gonic/gin"
)

func v1Routes(router *gin.RouterGroup, o *Options) {
	r := router.Group("/v1/")

	// middlewares
	r.Use(mw.ErrorHandlerX(o.Log))
	// add new routes here

}
