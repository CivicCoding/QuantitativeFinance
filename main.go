package main

import (
	"BinanceApi/dbService"
	_ "BinanceApi/dbService"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

const baseURL = "https://api.binance.com"

const secretKey = "91IW45xj7g8GCEjtYMievrfR0n7sFUO0mQWKiSL0R3HYi7p6RwxRChmXapex2z5R"
const apiKey = "zQDyK5C0dQcymyim7y3jyDLGj6rYQqNqfpDffEL8Ojw1uhzgHkf95hPIkNl9e1UX"

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

	var pp []person
	dbService.DB.Find(&pp)
	fmt.Println(pp)

	//dbService.CreateTable(&User{})
	//dbService.DeleteTable("students")

}
