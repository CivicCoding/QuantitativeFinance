package main

import (
	"QuantitativeFinance/dbService"
	"QuantitativeFinance/setting"
	_ "github.com/go-sql-driver/mysql"
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

var DB *gorm.DB

func main() {
	log.Println("Hello, api 正在启动中...")
	setting.SetUp("app")
	log.Println("使用的网络是：", setting.AppSetting.Url)
	log.Println("初始化连接数据库")
	DB = dbService.InitDB()
	log.Println("初始化交易所数据...")
	InitBinance()

}

// InitBinance 初始化填充交易所数据
func InitBinance() {
	// 首次获取kline、首次获取价格、首次获取账户信息
	// 测试下单等等
}
