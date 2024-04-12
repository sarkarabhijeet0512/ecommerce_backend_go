package handler

import (
	"context"
	"ecommerce_backend_project/er"
	"ecommerce_backend_project/internal/mw/jwt"
	"ecommerce_backend_project/pkg/product/offermangement"
	productdetails "ecommerce_backend_project/pkg/product/productDetails"
	model "ecommerce_backend_project/utils/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type OfferHandler struct {
	log                   *logrus.Logger
	jwtMiddleware         *jwt.GinJWTMiddleware
	offerService          *offermangement.Service
	productdetailsService *productdetails.Service
}

func newOfferHandler(
	log *logrus.Logger,
	offerService *offermangement.Service,
	productdetailsService *productdetails.Service,
) *OfferHandler {
	c := &gin.Context{}
	return &OfferHandler{
		log,
		jwt.SetAuthMiddleware(productdetailsService.Repo.GetDBConnection(c)),
		offerService,
		productdetailsService,
	}
}

func (h *OfferHandler) UpsertOfferDetails(c *gin.Context) {

	var (
		err  error
		res  = &model.GenericRes{}
		req  = &offermangement.Discount{}
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

	err = h.offerService.UpsertOfferDetails(dCtx, req)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusServiceUnavailable)
		return
	}
	res.Message = "Success"
	res.Data = req
	res.Success = true
	c.JSON(http.StatusOK, res)
}

func (h *OfferHandler) FetchOfferDetails(c *gin.Context) {
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

	data, err := h.offerService.FetchOfferByFilter(dCtx, req)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusServiceUnavailable)
		return
	}
	res.Message = "Success"
	res.Data = data
	res.Success = true
	c.JSON(http.StatusOK, res)
}
