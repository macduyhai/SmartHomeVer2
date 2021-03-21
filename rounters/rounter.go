package rounters

import (
	"github.com/macduyhai/SmartHomeVer2/config"
	"github.com/macduyhai/SmartHomeVer2/controlers"
	"github.com/macduyhai/SmartHomeVer2/middlewares"
	"github.com/macduyhai/SmartHomeVer2/services"

	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Router struct {
	config *config.Config
	db     *gorm.DB
}

func NewRouter(conf *config.Config, db *gorm.DB) Router {
	return Router{config: conf, db: db}
}

func (router *Router) InitGin() (*gin.Engine, error) {

	providerService := services.NewProvider(router.config, router.db)
	controller := controlers.NewController(providerService)

	engine := gin.Default()
	engine.Use(middlewares.CORSMiddleware())
	engine.Use(middlewares.RequestLogger())
	engine.GET("/ping", controller.Ping)
	engine.GET("/static", controller.StaticPage)

	accountAuthMiddleWare := middlewares.CheckAPIKey{ApiKey: router.config.APIKey}
	{
		account := engine.Group("/api/v1/account")
		account.Use(accountAuthMiddleWare.Check)
		account.POST("", controller.CreateUser)
		account.POST("/login", controller.Login)
	}
	{
		device := engine.Group("/api/v1/device")
		device.Use(accountAuthMiddleWare.Check)
		device.POST("/add", controller.AddDevice)
		device.POST("/list", controller.ListDevice)
		device.POST("/delete", controller.DeleteDevice)
		device.POST("/edit", controller.EditDevice)
		device.POST("/upload", controller.Upload)
		device.POST("/getstatus", controller.GetstatusDevice)
	}
	{
		log := engine.Group("/api/v1/log")
		log.Use(jwt.Auth(router.config.SecretKet))
		log.Use(middlewares.SetUserID)
		log.POST("", controller.CreateLog)
		log.GET("", controller.GetLogs)
	}

	{
		analysis := engine.Group("/api/v1/analysis")
		analysis.Use(jwt.Auth(router.config.SecretKet))
		analysis.Use(middlewares.SetUserID)
		analysis.GET("/tag", controller.AnalysisByTag)
		analysis.GET("/day", controller.AnalysisByDay)
	}

	{
		avg := engine.Group("/api/v1/average")
		avg.Use(jwt.Auth(router.config.SecretKet))
		avg.Use(middlewares.SetUserID)
		avg.GET("day", controller.GetAverageByDay)
	}

	return engine, nil
}
