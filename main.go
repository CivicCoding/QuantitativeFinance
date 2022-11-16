package main

import (
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

func main() {
	log.Println("Hello, api 正在启动中...")
	setting.SetUp("app")
	log.Println(setting.AppSetting.Url)
	log.Println("初始化数据库")

	//l, _ := time.LoadLocation("Asia/ShangHai")
	//t := time.Date(1998, 12, 22, 1, 12, 22, 0, l)
	//u := User{
	//	Id:       1,
	//	Name:     "lcw",
	//	Age:      10,
	//	Birthday: &t,
	//	Email:    "12@163.com",
	//	PassWord: "123",
	//}

}
