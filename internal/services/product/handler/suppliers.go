package handler

import (
	"context"
	"ecommerce_backend_project/er"
	"ecommerce_backend_project/internal/mw/jwt"
	productdetails "ecommerce_backend_project/pkg/product/productDetails"
	"ecommerce_backend_project/pkg/product/suppliers"
	model "ecommerce_backend_project/utils/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type SupplierHandler struct {
	log                   *logrus.Logger
	jwtMiddleware         *jwt.GinJWTMiddleware
	supplierService       *suppliers.Service
	productdetailsService *productdetails.Service
}

func newSupplierHandler(
	log *logrus.Logger,
	supplierService *suppliers.Service,
	productdetailsService *productdetails.Service,
) *SupplierHandler {
	c := &gin.Context{}
	return &SupplierHandler{
		log,
		jwt.SetAuthMiddleware(productdetailsService.Repo.GetDBConnection(c)),
		supplierService,
		productdetailsService,
	}
}

func (h *SupplierHandler) UpsertSuppliers(c *gin.Context) {

	var (
		err  error
		res  = &model.GenericRes{}
		req  = &suppliers.Supplier{}
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
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusBadRequest)
		res.Message = err.Error()
		return
	}

	err = h.supplierService.UpsertSuppliers(dCtx, req)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusServiceUnavailable)
		return
	}
	res.Message = "Success"
	res.Data = req
	res.Success = true
	c.JSON(http.StatusOK, res)
}

func (h *SupplierHandler) FetchSuppliers(c *gin.Context) {

	var (
		err  error
		res  = &model.GenericRes{}
		req  = model.Filter{}
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
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusBadRequest)
		res.Message = err.Error()
		return
	}

	data, err := h.supplierService.FetchSuppliersByFilter(dCtx, req)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusServiceUnavailable)
		return
	}
	res.Message = "Success"
	res.Data = data
	res.Success = true
	c.JSON(http.StatusOK, res)
}
