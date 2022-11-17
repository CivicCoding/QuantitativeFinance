package strategy

import (
	"QuantitativeFinance/binanceApi/wallet"
	"log"
	"strconv"
)

// NOTE：交易和投资是一场概率游戏，

// RiskRewardRatio 风险回报率 =（进场价-止损价）/（止盈价-进场价）
// TODO:  Risk reward ratio usually consider trade fee.
func RiskRewardRatio(entryPrice, takeProfitPrice float64, symbol string) float64 {
	stopPrice := entryPrice * 0.01
	tradeFee := wallet.TradeFee(symbol)
	makeFee, _ := strconv.ParseFloat(tradeFee.MakerCommission, 64)
	//takeFee, _ := strconv.ParseFloat(tradeFee.TakerCommission, 64)
	// Consume all trade direction is maker
	rRatio := (entryPrice*makeFee - stopPrice) / (takeProfitPrice*makeFee - entryPrice)
	return rRatio
}

// 账户风险 1%原则
// 止损价格 = 账户总资产 * 0.01
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
