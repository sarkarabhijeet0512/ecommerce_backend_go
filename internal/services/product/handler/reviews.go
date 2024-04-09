package handler

import (
	"context"
	"ecommerce_backend_project/er"
	"ecommerce_backend_project/internal/mw/jwt"
	productdetails "ecommerce_backend_project/pkg/product/productDetails"
	"ecommerce_backend_project/pkg/product/reviews"
	model "ecommerce_backend_project/utils/models"
	"net/http"

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

func (h *ReviewsHandler) UpsertReviewByProductID(c *gin.Context) {

	var (
		err  error
		res  = model.GenericRes{}
		req  = &reviews.Review{}
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
		res.Message = err.Error()
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}

	err = h.reviewService.UpdateReviewByProductID(dCtx, req)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}
	res.Message = "Success"
	res.Data = req
	res.Success = true
	c.JSON(http.StatusOK, res)
}

func (h *ReviewsHandler) FetchReview(c *gin.Context) {
	var (
		err  error
		res  = model.GenericRes{}
		req  = &model.Filter{}
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
		res.Message = err.Error()
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}
	data, err := h.reviewService.FetchReviewByFilter(dCtx, req)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}
	res.Message = "Success"
	res.Data = data
	res.Success = true
	c.JSON(http.StatusOK, res)
}
