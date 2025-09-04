package router

import (
	"epictectus/blog"
	"epictectus/config"
	"github.com/gin-gonic/gin"
	"net/http"

	"epictectus/handler"
	"epictectus/service"
)

type Options struct {
	Logger       blog.Logger
	Conf         *config.Config
	Dependencies *service.ServerDependencies
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Referer, User-Agent, X-Requested-With, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// To work API calls when both frontend and backend are hosted on local
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
	}
}

func InitRouter(opts Options) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(CorsMiddleware())

	router.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })
	userHandler := handler.NewUserHandler(opts.Dependencies.UserService)
	pgHandler := handler.NewPaymentGatewayHandler(opts.Dependencies.PaymentGatewayService)
	InitUserRouter(router, &userHandler)
	InitPgRouter(router, &pgHandler)
	return router
}

func InitUserRouter(router *gin.Engine, handler *handler.UserHandler) {
	v1 := router.Group("to-do/v1/user")
	v1.POST("sign-up", handler.SignUpUser)
	v1.POST("login", handler.LoginUser)
	//v1.POST("forgot-password", handler.ForgotPassword)
}

func InitPgRouter(router *gin.Engine, handler *handler.PgHandler) {
	v1 := router.Group("epictectus/v1")
	v1.POST("create-standard-payment-link", handler.CreateStandardPaymentLink)
}
