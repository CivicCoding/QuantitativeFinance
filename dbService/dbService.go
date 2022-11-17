// Package dbService
/*
	provide basic function to operate mysql database
*/
package dbService

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"sync"
)

var lock *sync.Mutex

var db *gorm.DB

// 单例模式

// InitDB 连接数据库
func InitDB() *gorm.DB {
	// "user:password@tcp(127.0.0.1:3306)/hello"
	dsn := "root:JXlwj5831@tcp(127.0.0.1:3306)/data_analysis?charset=utf8&parseTime=True&loc=Local"
	// 不会校验密码是否正确
	mysqlDB, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
	}
	// 设置与数据可建立连接的最大数目
	mysqlDB.DB().SetMaxOpenConns(100)
	// 设置连接池中的最大闲置连接数
	mysqlDB.DB().SetMaxIdleConns(10)
	if db == nil {
		lock.Lock()
		defer lock.Unlock()
		if db == nil {
			db = mysqlDB
		} else {
			log.Println("db is already created")
		}
	} else {
		log.Println("db is already created")
	}

	return db
}
