package handler

import (
	"context"
	"ecommerce_backend_project/er"
	"ecommerce_backend_project/internal/mw/jwt"
	productdetails "ecommerce_backend_project/pkg/product/productDetails"
	model "ecommerce_backend_project/utils/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ProductDetailsHandler struct {
	log                   *logrus.Logger
	jwtMiddleware         *jwt.GinJWTMiddleware
	productdetailsService *productdetails.Service
}

func newProductDetailsHandler(
	log *logrus.Logger,
	productdetailsService *productdetails.Service,
) *ProductDetailsHandler {
	c := &gin.Context{}
	return &ProductDetailsHandler{
		log,
		jwt.SetAuthMiddleware(productdetailsService.Repo.GetDBConnection(c)),
		productdetailsService,
	}
}

func (h *ProductDetailsHandler) UpsertProductDetails(c *gin.Context) {
	var (
		err error
		res = model.GenericRes{}
		// req  = &user.User{}
		dCtx = context.Background()
	)
	defer func() {
		if err != nil {
			c.Error(err)
			h.log.WithField("span", res).Warn(err.Error())
			return
		}
	}()
	if err = c.ShouldBind(&req); err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}
	err = h.productdetailsService.UpsertProductDetails(dCtx, req)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}
	res.Message = "Success"
	res.Data = req
	res.Success = true
	c.JSON(http.StatusOK, res)
}
func (h *ProductDetailsHandler) GetProductDetails(c *gin.Context) {
	var (
		err error
		res = model.GenericRes{}
		// req  = &user.User{}
		dCtx = context.Background()
	)
	defer func() {
		if err != nil {
			c.Error(err)
			h.log.WithField("span", res).Warn(err.Error())
			return
		}
	}()
	if err = c.ShouldBind(&req); err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}
	err = h.productdetailsService.GetProductByID(dCtx, req)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}
	res.Message = "Success"
	res.Data = req
	res.Success = true
	c.JSON(http.StatusOK, res)
}
func (h *ProductDetailsHandler) GetProductListByCategory(c *gin.Context) {
	var (
		err error
		res = model.GenericRes{}
		// req  = &user.User{}
		dCtx = context.Background()
	)
	defer func() {
		if err != nil {
			c.Error(err)
			h.log.WithField("span", res).Warn(err.Error())
			return
		}
	}()
	if err = c.ShouldBind(&req); err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}
	err = h.productdetailsService.GetProductListByCategory(dCtx, req)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}
	res.Message = "Success"
	res.Data = req
	res.Success = true
	c.JSON(http.StatusOK, res)
}
func (h *ProductDetailsHandler) GetProductCategory(c *gin.Context) {
	var (
		err error
		res = model.GenericRes{}
		// req  = &user.User{}
		dCtx = context.Background()
	)
	defer func() {
		if err != nil {
			c.Error(err)
			h.log.WithField("span", res).Warn(err.Error())
			return
		}
	}()
	if err = c.ShouldBind(&req); err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}
	err = h.productdetailsService.GetProductList(dCtx, req)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}
	res.Message = "Success"
	res.Data = req
	res.Success = true
	c.JSON(http.StatusOK, res)
}
func (h *ProductDetailsHandler) GetProductList(c *gin.Context) {
	var (
		err error
		res = model.GenericRes{}
		// req  = &user.User{}
		dCtx = context.Background()
	)
	defer func() {
		if err != nil {
			c.Error(err)
			h.log.WithField("span", res).Warn(err.Error())
			return
		}
	}()
	if err = c.ShouldBind(&req); err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}
	err = h.productdetailsService.GetProductList(dCtx, req)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}
	res.Message = "Success"
	res.Data = req
	res.Success = true
	c.JSON(http.StatusOK, res)
}
func (h *ProductDetailsHandler) DisableProduct(c *gin.Context) {
	var (
		err error
		res = model.GenericRes{}
		// req  = &user.User{}
		dCtx = context.Background()
	)
	defer func() {
		if err != nil {
			c.Error(err)
			h.log.WithField("span", res).Warn(err.Error())
			return
		}
	}()
	if err = c.ShouldBind(&req); err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}

	err = h.productdetailsService.DisableProductByID(dCtx, req)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}
}
