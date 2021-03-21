package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kelseyhightower/envconfig"
	"github.com/macduyhai/SmartHomeVer2/config"
	"github.com/macduyhai/SmartHomeVer2/rounters"
	"github.com/macduyhai/SmartHomeVer2/services"
)

func main() {
	r := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.Static("/", "./public")

	// Init connect mqtt
	services.MqttBegin()

	conf := config.Config{}
	if err := envconfig.Process("", &conf); err != nil {
		panic(err)
	}
	db, err := gorm.Open("mysql", conf.MySQLURL)
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	if err != nil {
		panic("open db error: " + err.Error())
	}

	if err := db.DB().Ping(); err != nil {
		panic("ping db error: " + err.Error())
	}
	router := rounters.NewRouter(&conf, db)
	app, _ := router.InitGin()
	_ = app.Run(":80")
}
