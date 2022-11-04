package main

import (
	"QuantitativeFinance/binanceApi/wallet"
	"QuantitativeFinance/setting"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type User struct {
	Id       uint   `gorm:"AUTO_INCREMENT"`
	Name     string `gorm:"size:50"`
	Age      int    `gorm:"size:3"`
	Birthday *time.Time
	Email    string `gorm:"type:varchar(50);unique_index"`
	PassWord string `gorm:"type:varchar(25)"`
}

type person struct {
	gorm.Model
	Name    string `gorm:"size:50"`
	Age     int    `gorm:"size:3"`
	address string `gorm:"size:50"`
}

func main() {
	log.Println("Hello, api 正在启动中...")
	setting.SetUp("app")
	log.Println(setting.AppSetting.Url)
	//wallet.AccountSnapshot("SPOT", "7")
	wallet.TradeFee("BTCBUSD")
	//configInfo := url2.Values{}
	//configInfo.Add("recvWindow", "5000")
	//configInfo.Add("timestamp", "16689757")
	//configInfo.Add("signature", "signature")
	////data := configInfo.Encode()
	//fmt.Println(configInfo)
}
