// Package dbService
/*
	provide basic function to operate mysql database
*/
package dbService

import (
	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func initDB() (err error) {
	// "user:password@tcp(127.0.0.1:3306)/hello"
	dsn := "root:JXlwj5831@tcp(127.0.0.1:3306)/data_analysis?charset=utf8&parseTime=True&loc=Local"
	// 不会校验密码是否正确
	DB, err = gorm.Open("mysql", dsn)
	defer DB.Close()
	if err != nil {
		return err
	}

	// 设置与数据可建立连接的最大数目
	DB.DB().SetMaxOpenConns(100)
	// 设置连接池中的最大闲置连接数
	DB.DB().SetMaxIdleConns(10)
	return nil
}

func init() {
	err := initDB()
	if err != nil {
		color.Red("initDB:", err)
	}
}
