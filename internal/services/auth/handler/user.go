package handler

import (
	"context"
	"ecommerce_backend_project/er"
	"ecommerce_backend_project/internal/mw/jwt"
	"ecommerce_backend_project/pkg/auth/user"
	"fmt"
	"net/http"

	model "ecommerce_backend_project/utils/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	log           *logrus.Logger
	jwtMiddleware *jwt.GinJWTMiddleware
	userService   *user.Service
}

func newUserHandler(
	log *logrus.Logger,
	userService *user.Service,
) *UserHandler {
	c := &gin.Context{}
	return &UserHandler{
		log,
		jwt.SetAuthMiddleware(userService.Repo.GetDBConnection(c)),
		userService,
	}
}

func (h *UserHandler) UserRegistration(c *gin.Context) {
	var (
		err  error
		res  = model.GenericRes{}
		req  = &user.User{}
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
	err = h.userService.UpsertUserRegistration(dCtx, req)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}
	res.Message = "Registration Sucessfully Done"
	res.Success = true
	res.Data = req
	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) UserLogin(c *gin.Context) {
	var (
		err  error
		dCtx = context.Background()
		req  = user.User{}
		res  = model.GenericRes{}
	)
	defer func() {
		if err != nil {
			c.Error(err)
			h.log.WithField("span", req).Warn(err.Error())
		}
	}()

	//check if location pings are present for today
	//if not, then verify that rider should be at store before proceeding
	if err = c.ShouldBind(&req); err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}

	user, ok, err := h.userService.FetchUserByMobileNumberOrEmail(dCtx, req)
	if err != nil {
		err = er.New(err, er.UserNotFound).SetStatus(http.StatusNotFound)
		return
	}
	if ok {
		code, token, expire := h.jwtMiddleware.SetToken(c, user)

		if code == 0 {
			err = fmt.Errorf("jwt set token failed")
			return
		}
		fmt.Println(token, expire)
		// res.Code = code
		// res.Token = token
		// res.Expire = &expire
		// res.Success = success
		c.JSON(http.StatusOK, res)
		return
	}

}
