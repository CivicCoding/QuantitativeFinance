package model

import (
	"QuantitativeFinance/binanceApi/market"
	"QuantitativeFinance/dbService"
	"QuantitativeFinance/setting"
	"fmt"
	"time"
)

type kline struct {
	Id            int
	OpeningTime   time.Time // k线开盘时间
	OpeningPrice  string    // 开盘价
	HighestPrice  string    // 最高价
	LowestPrice   string    // 最低价
	ClosingPrice  string    // 收盘价
	Volume        string    // 成交量
	ClosingTime   time.Time // k线收盘时间
	Turnover      string    // 成交额
	NumOfTrans    string    // 成交笔数
	ActiveBuyVol  string    // 主动买入成交量
	ActiveBuyTurn string    // 主动买入成交额
	ignore        string    // 忽略该参数
}

func GetKlineData(query interface{}, arg interface{}) {
	res := dbService.DB.Select(query).Where(arg).Value
	fmt.Println(res)
}

type klineCtl struct{}

func (k *klineCtl) Insert() {
	res := market.Kline(setting.AppSetting.Url, "ETHBUSD")
	for _, item := range *res {
		dbService.DB.Create(item)
	}
}
