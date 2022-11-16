package strategy

import (
	"QuantitativeFinance/binanceApi/wallet"
	"log"
	"strconv"
)

// NOTE：交易和投资是一场概率游戏，
// 风险回报率 =（进场价-止损价）/（止盈价-进场价）

// 止损价格 = 账户总资产 * 0.01
// 账户风险 1%原则
// 失效点 止损范围

// position 头寸 = 账户大小 * 账户风险 / 失效点
func position(symbol string) float64 {
	asset := wallet.GetFundingAsset(symbol)
	balance, err := strconv.ParseFloat(asset.Free, 64)
	if err != nil {
		log.Println(err)
	}
	p := balance * 0.01 / 0.05
	return p
}
