package server

import (
	"ecommerce_backend_project/internal/mw"
	"ecommerce_backend_project/internal/mw/jwt"
	"ecommerce_backend_project/internal/services/auth/handler"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	ginlogrus "github.com/toorop/gin-logrus"
	"go.uber.org/fx"
)

// Module invokes mainserver
var Module = fx.Options(
	fx.Invoke(
		Run,
	),
)

const (
	addr = "0.0.0.0"
)

// Options is function arguments struct of `Run` function.
type Options struct {
	fx.In

	Config          *viper.Viper
	Log             *logrus.Logger
	PostgresDB      *pg.DB      `name:"userdb"`
	Redis           *redis.Pool `name:"redisWorker"`
	UserHandler     *handler.UserHandler
	UserRoleHandler *handler.UserRoleHandler
}

// Run starts the mainserver REST API server
func Run(o Options) {
	router := SetupRouter(&o)
	router.Run(fmt.Sprintf("%s:%s", addr, o.Config.GetString("port")))
}

// SetupRouter creates gin router and registers all user routes to it
func SetupRouter(o *Options) (router *gin.Engine) {
	gin.SetMode(gin.ReleaseMode)
	router = gin.New()

	// Logs all panic to error log
	router.Use(ginlogrus.Logger(o.Log), gin.Recovery())

	// Health routes
	router.Use(mw.RateLimiter(10, time.Minute))
	authMiddleware := jwt.SetAuthMiddleware(o.PostgresDB)
	router.GET("/_healthz", HealthHandler(o))
	router.GET("/_readyz", HealthHandler(o))

	rootRouter := router.Group("/")

	v1Routes(rootRouter, authMiddleware, o)

	return
}

// HealthHandler
func HealthHandler(o *Options) func(*gin.Context) {
	return func(c *gin.Context) {
		jwt.ExtractClaims(c)
		err := o.PostgresDB.Ping(c)
		if err != nil {
			c.AbortWithError(http.StatusFailedDependency, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"ok": "ok"})
	}
}
