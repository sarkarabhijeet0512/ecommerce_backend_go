// Package mw is user Middleware package
package mw

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"ecommerce_backend_project/er"
	"ecommerce_backend_project/pkg/cache"
	model "ecommerce_backend_project/utils/models"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"github.com/sirupsen/logrus"
)

func ErrorHandlerX(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := c.Errors.Last()
			if err == nil {
				// no errors, abort with success
				return
			}

			e := er.From(err.Err)

			if !e.NOP {
				sentry.CaptureException(e)
			}

			httpStatus := http.StatusInternalServerError
			if e.Status > 0 {
				httpStatus = e.Status
			}

			c.JSON(httpStatus, e)
		}()

		c.Next()
	}
}

func RateLimiter(limit int64, duration time.Duration) gin.HandlerFunc {
	// Create a new token bucket with given limit and duration
	bucket := ratelimit.NewBucketWithQuantum(duration, limit, limit)

	return func(c *gin.Context) {
		// If bucket has tokens available, allow the request
		if bucket.TakeAvailable(1) >= 1 {
			c.Next()
			return
		}

		// If no tokens available, return 429 Too Many Requests
		c.AbortWithStatus(http.StatusTooManyRequests)
	}
}

func AWSSessionAttach(sess *session.Session) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("sess", sess)
		c.Next()
	}
}
func RoleCheckMiddleware(rdb *cache.Service, role int) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, ok := c.Get("id")
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "Please login uisng your secured credentials!"})
			c.Abort()
		}
		userRoles := model.UserRoles{}
		err := rdb.Get(fmt.Sprint(userID), &userRoles)
		if err != nil {
			err = er.New(err, er.Unauthorized).SetStatus(http.StatusUnauthorized)
			c.Abort()
		}
		roleFound := false
		for _, userRole := range userRoles.Resource {
			if role == userRole {
				roleFound = true
				break
			}
		}
		if !roleFound {
			er.New(errors.New("invalid access"), er.Unauthorized).SetStatus(http.StatusUnauthorized)
			c.Abort()
			return
		}
		c.Next()
	}
}
