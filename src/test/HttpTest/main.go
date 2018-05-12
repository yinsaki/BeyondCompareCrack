package main

import (
	"time"
	"fmt"
	"test/HttpTest/module"
	"test/HttpTest/config"
	"test/HttpTest/core/Log"
)

func main(){
	time.LoadLocation("Asia/Chongqing")

	config.G_appConfig.ParseConfig("config.ini")
	config.G_envConfig.ParseConfig("env.ini")

	fmt.Println("config:", config.G_appConfig.APP_USER, config.G_appConfig.APP_PASS)

	loginer := new(module.Login)
	loginer.Run()
	Log.Debug("--------------------------%d",1111)
	Log.Info("dfsfafafsdfsdfs")
	Log.Error("fradfasdfsadfadfsadfs")
}













/*

func main() {
	url := "http://wiki.dw7758.com"
	fmt.Println("URL:>", url)

	userId := "dw"
	pwd := "dw888@88"

	payLoad := "{" +
		" \"UserId\":\"" + userId +
		"\",\"Password\":\"" + pwd +
		"\"}"
	fmt.Println(url, "post", payLoad)

	var jsonStr = []byte(payLoad)
	fmt.Println("jsonStr", jsonStr)

	fmt.Println("new_str", bytes.NewBuffer(jsonStr))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	fmt.Println("resp", resp)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
}
*/
