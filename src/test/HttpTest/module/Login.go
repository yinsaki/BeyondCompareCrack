package module

import (
	"fmt"
	"test/HttpTest/config"
	"time"
)

type Login struct {
	time uint64
}

func (this *Login)Run(){
	fmt.Println("正在装载终端")
	if config.G_appConfig.ACCESS_TOKEN == "" {
		fmt.Println("加载token中...")
		this.Login()
	}

	fmt.Println("正在检查token合法性...")

	if !this.info() {
		fmt.Println("令牌即将过期")
		fmt.Println("申请更换令牌中...")
		if(!this.Refresh()){
			fmt.Println("无效令牌，正在重新申请...")
			this.Login()
		}
	}

	this.time = uint64(time.Now().Unix()) + 3600
}

func (this *Login)info()bool {
	//accessToken := config.G_appConfig.ACCESS_TOKEN
	
	return true
}

func (this *Login)Refresh() bool{

	return true
}

func (this *Login)Login(){

}