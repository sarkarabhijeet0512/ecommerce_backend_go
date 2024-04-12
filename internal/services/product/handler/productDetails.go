package handler

import (
	"context"
	"ecommerce_backend_project/er"
	"ecommerce_backend_project/internal/mw/jwt"
	productdetails "ecommerce_backend_project/pkg/product/productDetails"
	"ecommerce_backend_project/utils"
	model "ecommerce_backend_project/utils/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aws/aws-sdk-go/aws/session"
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
		err  error
		res  = &model.GenericRes{}
		req  = &productdetails.Product{}
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
		err  error
		res  = model.GenericRes{}
		dCtx = context.Background()
	)
	defer func() {
		if err != nil {
			c.Error(err)
			h.log.WithField("span", res).Warn(err.Error())
			return
		}
	}()
	productID, err := strconv.Atoi(fmt.Sprint(c.Param("product_id")))
	if err != nil {
		h.log.WithField("span", productID).Info("error while converting string to int: " + err.Error())
		return
	}
	data, err := h.productdetailsService.GetProductByID(dCtx, productID)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}
	res.Message = "Success"
	res.Data = data
	res.Success = true
	c.JSON(http.StatusOK, res)
}

func (h *ProductDetailsHandler) GetProductListByCategory(c *gin.Context) {
	var (
		err  error
		res  = model.GenericRes{}
		dCtx = context.Background()
	)
	defer func() {
		if err != nil {
			c.Error(err)
			h.log.WithField("span", res).Warn(err.Error())
			return
		}
	}()
	categoryID, err := strconv.Atoi(fmt.Sprint(c.Param("category_id")))
	if err != nil {
		h.log.WithField("span", categoryID).Info("error while converting string to int: " + err.Error())
		return
	}
	data, err := h.productdetailsService.GetProductListByCategory(dCtx, categoryID)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}
	res.Message = "Success"
	res.Data = data
	res.Success = true
	c.JSON(http.StatusOK, res)
}

func (h *ProductDetailsHandler) GetProductCategory(c *gin.Context) {
	var (
		err  error
		res  = model.GenericRes{}
		dCtx = context.Background()
	)
	defer func() {
		if err != nil {
			c.Error(err)
			h.log.WithField("span", res).Warn(err.Error())
			return
		}
	}()
	data, err := h.productdetailsService.GetProductCategoryList(dCtx)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}
	res.Message = "Success"
	res.Data = data
	res.Success = true
	c.JSON(http.StatusOK, res)
}

// func (h *ProductDetailsHandler) GetProductList(c *gin.Context) {
// 	var (
// 		err error
// 		res = model.GenericRes{}
// 		// req  = &user.User{}
// 		dCtx = context.Background()
// 	)
// 	defer func() {
// 		if err != nil {
// 			c.Error(err)
// 			h.log.WithField("span", res).Warn(err.Error())
// 			return
// 		}
// 	}()
// 	if err = c.ShouldBind(&req); err != nil {
// 		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
// 		return
// 	}
// 	err = h.productdetailsService.GetProductList(dCtx, req)
// 	if err != nil {
// 		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
// 		return
// 	}
// 	res.Message = "Success"
// 	res.Data = req
// 	res.Success = true
// 	c.JSON(http.StatusOK, res)
// }
// func (h *ProductDetailsHandler) DisableProduct(c *gin.Context) {
// 	var (
// 		err error
// 		res = model.GenericRes{}
// 		// req  = &user.User{}
// 		dCtx = context.Background()
// 	)
// 	defer func() {
// 		if err != nil {
// 			c.Error(err)
// 			h.log.WithField("span", res).Warn(err.Error())
// 			return
// 		}
// 	}()
// 	if err = c.ShouldBind(&req); err != nil {
// 		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
// 		return
// 	}

//		err = h.productdetailsService.DisableProductByID(dCtx, req)
//		if err != nil {
//			err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
//			return
//		}
//	}
func (h *ProductDetailsHandler) UploadProductImages(c *gin.Context) {
	var (
		err  error
		res  = model.GenericRes{}
		req  = &productdetails.ProductImage{}
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

	userID, ok := c.Get("id") //take this value from user session
	if !ok {
		err = er.New(err, er.Unauthorized).SetStatus(http.StatusUnauthorized)
		return
	}
	if v, ok := userID.(int); ok {
		req.UploadedBy = v
	}

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		h.log.WithField("span", file).Info(err.Error())
		return
	}
	defer file.Close()
	filename, contentType := utils.FileProcessing(file, header, req.ProductID)
	if ProductID, ok := c.GetQuery("product_id"); ok {
		if ProductID != "" {
			ID, err := strconv.Atoi(ProductID)
			if err != nil {
				return
			}
			req.ProductID = ID
		}
	}
	err = h.productdetailsService.UploadProductImage(req, file, filename, contentType, c.MustGet("sess").(*session.Session), dCtx)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusServiceUnavailable)
		return
	}
	res.Message = "Success"
	res.Data = req
	res.Success = true
	c.JSON(http.StatusOK, res)
}
