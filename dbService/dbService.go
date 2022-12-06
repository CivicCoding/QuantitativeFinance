// Package dbService
/*
	provide basic function to operate mysql database
*/
package dbService

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var _db *gorm.DB

// 单例模式

// init 连接数据库 使用连接池来避免高并发重复建立连接数据库连接的性能消耗
func init() {
	var err error
	// "user:password@tcp(127.0.0.1:3306)/hello"
	dsn := "root:JXlwj5831@tcp(127.0.0.1:3306)/data_analysis?charset=utf8&parseTime=True&loc=Local"
	// 不会校验密码是否正确
	_db, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
	}
	sqlDB := _db.DB()
	// 设置与数据可建立连接的最大数目
	sqlDB.SetMaxOpenConns(100)
	// 设置连接池中的最大闲置连接数
	sqlDB.SetMaxIdleConns(10)
	_db.AutoMigrate()
}

func GetDB() *gorm.DB {
	return _db
}
