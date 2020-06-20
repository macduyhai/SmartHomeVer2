package main

import (
	"SmartHomeVer2/config"
	"SmartHomeVer2/rounters"
	"SmartHomeVer2/services"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kelseyhightower/envconfig"
)

func main() {
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
	_ = app.Run(":9191")
}
