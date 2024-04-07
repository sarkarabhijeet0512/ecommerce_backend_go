package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	ginlogrus "github.com/toorop/gin-logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
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

	Config *viper.Viper
	Log    *logrus.Logger

	MysqlDB *gorm.DB `name:"pointoDB"`
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

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.

	// Logs all panic to error log
	router.Use(ginlogrus.Logger(o.Log), gin.Recovery())

	// Health routes
	router.GET("/_healthz", HealthHandler(o))
	router.GET("/_readyz", HealthHandler(o))

	rootRouter := router.Group("/")

	v1Routes(rootRouter, o)

	return
}

// HealthHandler
func HealthHandler(o *Options) func(*gin.Context) {
	return func(c *gin.Context) {
		var err error
		db, err := o.MysqlDB.DB()
		if err != nil {
			c.AbortWithError(http.StatusFailedDependency, err)
			return
		}
		err = db.Ping()
		if err != nil {
			c.AbortWithError(http.StatusFailedDependency, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"ok": "ok"})
	}
}
