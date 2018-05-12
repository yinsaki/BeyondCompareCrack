package module

import (
	"test/HttpTest/config"
	"time"
)

func Sign(argv map[string]string) map[string]string {
	defaultMap := make(map[string]string)

	defaultMap["access_key"] = config.G_appConfig.ACCESS_TOKEN
	defaultMap["actionKey"] = "appkey"
	defaultMap["appkey"] = config.G_envConfig.APP_KEY
	defaultMap["appver"] = "6620"
	defaultMap["build"] = "6620"
	defaultMap["device"] = "phone"
	defaultMap["mobi_app"] = "iphone"
	defaultMap["platform"] = "ios"
	defaultMap["ts"] = time.Now().String()
	defaultMap["type"] = "json"

	for k,v := range argv {
		defaultMap[k] = v
	}


	return defaultMap
}