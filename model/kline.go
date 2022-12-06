package model

import (
	"time"
)

type Kline struct {
	Id            int       `gorm:"AUTO_INCREMENT"`
	OpeningTime   time.Time `gorm:"type:varchar(25)"` // k线开盘时间
	OpeningPrice  string    `gorm:"type:varchar(25)"` // 开盘价
	HighestPrice  string    `gorm:"type:varchar(25)"` // 最高价
	LowestPrice   string    `gorm:"type:varchar(25)"` // 最低价
	ClosingPrice  string    `gorm:"type:varchar(25)"` // 收盘价
	Volume        string    `gorm:"type:varchar(25)"` // 成交量
	ClosingTime   time.Time `gorm:"type:varchar(25)"` // k线收盘时间
	Turnover      string    `gorm:"type:varchar(25)"` // 成交额
	NumOfTrans    string    `gorm:"type:varchar(25)"` // 成交笔数
	ActiveBuyVol  string    `gorm:"type:varchar(25)"` // 主动买入成交量
	ActiveBuyTurn string    `gorm:"type:varchar(25)"` // 主动买入成交额
	Ignore        string    `gorm:"type:varchar(25)"` // 忽略该参数
}

// TODO: 未来存数据库需要分成不同币种的k线
//func GetKlineData(query interface{}, arg interface{}) {
//	res := .Select(query).Where(arg).Value
//	fmt.Println(res)
//}

type KlineCtl struct{}

// Insert 插入
func (k *KlineCtl) Insert(value []Kline) {
	//DB := dbService.InitDB()
	//for _, item := range value {
	//	DB.Create(&item)
	//}
}
