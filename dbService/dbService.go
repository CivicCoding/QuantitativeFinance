// Package dbService
/*
	provide basic function to operate mysql database
*/
package dbService

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	// "user:password@tcp(127.0.0.1:3306)/hello"
	dsn := "root:JXlwj5831@tcp(127.0.0.1:3306)/data_analysis?charset=utf8&parseTime=True&loc=Local"
	// 不会校验密码是否正确
	DB, _ = gorm.Open("mysql", dsn)
	// 设置与数据可建立连接的最大数目
	DB.DB().SetMaxOpenConns(100)
	// 设置连接池中的最大闲置连接数
	DB.DB().SetMaxIdleConns(10)
	return DB
}
