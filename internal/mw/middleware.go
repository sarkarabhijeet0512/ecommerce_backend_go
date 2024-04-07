// Package mw is user Middleware package
package mw

import (
	"net/http"
	"time"

	"ecommerce_backend_project/er"

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
