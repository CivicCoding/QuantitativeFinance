package market

import (
	"QuantitativeFinance/binanceApi/common"
	"github.com/fatih/color"
	"net/http"
	"os"
	"strings"
)

type timeS struct {
	ServerTime int `json:"ServerTime"`
}

// ServerTime 获取Binance服务器时间 api/v3/time
func ServerTime(baseUrl string) int {
	resp, err := http.Get(baseUrl + "/api/v3/time")
	if err != nil {
		color.Red("Something wrong with time \n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	s := common.HandleResponse(resp)

	var t timeS

	common.JsonStringToStruct(s, &t)

	return t.ServerTime
}

type kline struct {
	OpeningTime   string // k线开盘时间
	OpeningPrice  string // 开盘价
	HighestPrice  string // 最高价
	LowestPrice   string // 最低价
	ClosingPrice  string // 收盘价
	Volume        string // 成交量
	ClosingTime   string // k线收盘时间
	Turnover      string // 成交额
	NumOfTrans    string // 成交笔数
	ActiveBuyVol  string // 主动买入成交量
	ActiveBuyTurn string // 主动买入成交额
	ignore        string // 忽略该参数
}

// Kline
/*
获取指定币对的k线数据

response:
[
	  [
	    1499040000000,      // k线开盘时间
	    "0.01634790",       // 开盘价
	    "0.80000000",       // 最高价
	    "0.01575800",       // 最低价
	    "0.01577100",       // 收盘价(当前K线未结束的即为最新价)
	    "148976.11427815",  // 成交量
	    1499644799999,      // k线收盘时间
	    "2434.19055334",    // 成交额
	    308,                // 成交笔数
	    "1756.87402397",    // 主动买入成交量
	    "28.46694368",      // 主动买入成交额
	    "17928899.62484339" // 请忽略该参数
	  ]
]
*/
func Kline(baseUrl string, symbol string) *[]kline {
	url := baseUrl + "/api/v3/klines" + "?symbol=" + symbol + "&interval=" + "1h" + "&limit=500"
	resp, err := http.Get(url)
	if err != nil {
		color.Red("Something wrong with kline:\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	res := common.HandleResponse(resp)

	// step1 : create a string slice to store response
	var data []string
	// step2 : remove "," and "[]"
	s := strings.Split(res, ",")
	for _, item := range s {
		temp := strings.Trim(item, "[]")
		data = append(data, temp)
	}

	// step3 : append elem to slice of kline struct

	var kl kline
	var klines []kline

	for i := 0; i < len(data); i++ {
		switch {
		case i%12 == 0:
			kl.OpeningTime = data[i]
		case i%12 == 1:
			kl.OpeningPrice = data[i]
		case i%12 == 2:
			kl.HighestPrice = data[i]
		case i%12 == 3:
			kl.LowestPrice = data[i]
		case i%12 == 4:
			kl.ClosingPrice = data[i]
		case i%12 == 5:
			kl.Volume = data[i]
		case i%12 == 6:
			kl.ClosingTime = data[i]
		case i%12 == 7:
			kl.Turnover = data[i]
		case i%12 == 8:
			kl.NumOfTrans = data[i]
		case i%12 == 9:
			kl.ActiveBuyVol = data[i]
		case i%12 == 10:
			kl.ActiveBuyTurn = data[i]
		case i%12 == 11:
			kl.ignore = data[i]
			klines = append(klines, kl)
		}

	}

	return &klines
}

// Depth 获取深度数据 /api/v3/depth
func Depth(baseUrl string, pair string) string {
	resp, err := http.Get(baseUrl + "/api/v3/depth" + "?symbol=" + pair)
	if err != nil {
		color.Red("Depth:", err)
	}
	return common.HandleResponse(resp)
}

type avgPrice struct {
	Mins  int    `json:"mins"`
	Price string `json:"price"`
}

// AvgPrice 获取币对当前均价 /api/v3/avgPrice
func AvgPrice(baseUrl string, pair string) string {
	resp, err := http.Get(baseUrl + "/api/v3/avgPrice" + "?symbol=" + pair)
	if err != nil {
		color.Red("average", err)
	}
	return common.HandleResponse(resp)
}
