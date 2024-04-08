package server

import (
	"ecommerce_backend_project/internal/mw"
	"ecommerce_backend_project/internal/mw/jwt"

	"github.com/gin-gonic/gin"
)

func v1Routes(router *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware, o *Options) {
	r := router.Group("/v1/auth/api/")
	// middlewares
	r.Use(mw.ErrorHandlerX(o.Log))
	r.PUT("/user_registration", o.UserHandler.UserRegistration)
	r.POST("/user_login", o.UserHandler.UserLogin)
	// RBAC
	r.PUT("/user_role", authMiddleware.MiddlewareFunc(), o.UserRoleHandler.CreateUserRole)
	r.POST("/assign_role", authMiddleware.MiddlewareFunc(), o.UserRoleHandler.RoleAssignment)
	r.GET("/user/assigned_role", authMiddleware.MiddlewareFunc(), o.UserRoleHandler.UserRoleAssignedDetails)
	r.GET("/role_details/:role_id", authMiddleware.MiddlewareFunc(), o.UserRoleHandler.RoleDetails)
	r.GET("/role_list", authMiddleware.MiddlewareFunc(), o.UserRoleHandler.RoleList)
}
