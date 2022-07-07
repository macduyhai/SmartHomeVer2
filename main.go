package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/macduyhai/SmartHomeVer2/config"
	"github.com/macduyhai/SmartHomeVer2/rounters"
	"github.com/macduyhai/SmartHomeVer2/services"
)

func main() {
	fmt.Println("Server starting ...")
	// Init connect mqtt
	services.MqttBegin()

	// conf := config.Config{}
	conf := config.NewConfig()

	// if err := envconfig.Process("", &conf); err != nil {
	// 	fmt.Println(err)
	// }
	db, err := gorm.Open("mysql", conf.MySQLURL)
	defer func() {
		if err := db.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	if err != nil {
		panic("open db error: " + err.Error())
	}

	if err := db.DB().Ping(); err != nil {
		panic("ping db error: " + err.Error())
	}
	router := rounters.NewRouter(conf, db)
	app, _ := router.InitGin()

	_ = app.Run(":80")
}
