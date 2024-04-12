package handler

import (
	"context"
	"ecommerce_backend_project/er"
	"ecommerce_backend_project/internal/mw/jwt"
	"ecommerce_backend_project/pkg/product/inventory"
	productdetails "ecommerce_backend_project/pkg/product/productDetails"
	model "ecommerce_backend_project/utils/models"
	"net/http"

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

func (h *InventoryHandler) UpsertInventory(c *gin.Context) {
	var (
		err  error
		res  = &model.GenericRes{}
		req  = &inventory.Inventory{}
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

	err = h.invertoryService.UpsertInventory(dCtx, req)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusServiceUnavailable)
		return
	}
	res.Message = "Success"
	res.Data = req
	res.Success = true
	c.JSON(http.StatusOK, res)
}

func (h *InventoryHandler) FetchInventory(c *gin.Context) {
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
	data, err := h.invertoryService.FetchInventoryByFilter(dCtx, req)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusServiceUnavailable)
		return
	}
	res.Message = "Success"
	res.Data = data
	res.Success = true
	c.JSON(http.StatusOK, res)
}
