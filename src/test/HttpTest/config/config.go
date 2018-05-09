package main

import (
	"test/HttpTest/config"
)
type ConfigInterface interface {
	ParseConfig(fileName string)
}

type AppConfig struct {
	APP_USER string
	APP_PASS string

	ACCESS_TOKEN string
	REFRESH_TOKEN string

	APP_DEBUG bool

	ROOM_ID int
}

type EnvConfig struct {
	APP_KEY string
	APP_SECRET string
}


func (this *AppConfig)ParseConfig(fileName string)  {
	iniconf, err := config.NewConfig("ini", fileName)
	if err != nil {
		panic(err)
		return
	}

	this.APP_USER = iniconf.DefaultString("APP_USER", "")
	this.APP_PASS = iniconf.DefaultString("APP_PASS", "")

	this.ACCESS_TOKEN = iniconf.DefaultString("ACCESS_TOKEN", "")
	this.REFRESH_TOKEN = iniconf.DefaultString("REFRESH_TOKEN", "")

	this.APP_DEBUG = iniconf.DefaultBool("APP_DEBUG",true)
	this.ROOM_ID = iniconf.DefaultInt("ROOM_ID", 10001)

}

func (this *EnvConfig)ParseConfig(fileName string) {
	iniconf, err := config.NewConfig("ini", fileName)
	if err != nil {
		panic(err)
		return
	}

	this.APP_KEY = iniconf.DefaultString("APP_KEY", "")
	this.APP_SECRET = iniconf.DefaultString("APP_SECRET", "")
}

var g_appConfig  = new(AppConfig)
var g_envConfig  = new(EnvConfig)