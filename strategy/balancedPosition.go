package strategy

import (
	"QuantitativeFinance/binanceApi/enum"
	"QuantitativeFinance/binanceApi/market"
	"QuantitativeFinance/binanceApi/spot"
	"QuantitativeFinance/binanceApi/wallet"
	"log"
	"math"
	"strconv"
)

// 账户余额 单位BUSD
var accountBalance float64

// 账户持仓 单位BUSD
var latestPosition float64

// 币对 example："BNBBUSD"
var coinPair string

/*
BalancedPosition 均仓策略，50-50策略，也就是持仓和现金保持50:50比例的策略，
先买入一个股票过一段时间调整仓位比例为50:50。
均仓策略本质为网格策略，其收益来源于价格一定范围内来回波动，所以在震荡行情中表现 会更好
优势：均仓策略本质为网格策略，其收益来源于价格一定范围内来回波动，所以在震荡行情中表现 会更好

劣势：风险在于执行调仓操作后，价格继续单边上涨或下跌。

补充：一般价格变动范围设置在手续费的4倍左右为基准开始调优较为合理，市场活跃时，可以是手续费的1.5~2倍左右。
市场不太活跃时，手续费的8倍，十倍，甚至二十，五十倍都是可以的。

控制最小交易数量，也可以降低交易频率，增加抓取到更优点位的概率，从而提高收益。

请注意，该策略是在现货市场对现价范围波动进行调仓。
TODO: 将下单信息保存至数据库
*/
func BalancedPosition(coin1, coin2 string) {
	FirstBalancedPosition(coin1, coin2)
	// step3 ：价格到达变动范围时调整仓位,循环到结束信号传入
	// TODO: 设定到达一段时间后平衡仓位
	for {
		// 查询最新价格
		latestPrice, err := strconv.ParseFloat(market.Price(coinPair).Price, 32)
		if err != nil {
			log.Println(err)
		}
		// 获取用户持币数量
		u := wallet.GetUserAsset(coin1, false)
		free, err := strconv.ParseFloat(u.Free, 32)
		if err != nil {
			log.Println(err)
		}
		//计算用户持仓
		latestPosition = latestPrice * free
		// step 4: 将持仓coin换算成cash多的取出来少的添进去
		balancePosition(latestPrice, latestPosition, accountBalance, coinPair)
	}
}

// FirstBalancedPosition 第一次对资产配置
func FirstBalancedPosition(coin1, coin2 string) {

	//币对
	coinPair = coin1 + coin2
	// step1 : 获取账户现金资产
	asset := wallet.GetFundingAsset(coin2)
	amount, err := strconv.ParseFloat(asset.Free, 64)
	if err != nil {
		log.Println(err)
	}
	// 获取币种最新价格
	latestPrice, err := strconv.ParseFloat(market.Price(coinPair).Price, 32)
	if err != nil {
		log.Println(err)
	}
	// 计算购买数量
	quantity := strconv.FormatFloat(amount/2/latestPrice, 'e', 5, 32)
	// 设置最新价格为限价
	price := market.Price(coinPair).Price
	//一般价格变动范围设置在手续费的4倍左右为基准开始调优较为合理
	fee, _ := strconv.ParseFloat(wallet.TradeFee(coinPair).MakerCommission, 64)
	trailingDelta := strconv.FormatFloat(fee*10000*4, 'e', 5, 32)
	spot.TrailingDeltaOrder(coinPair, enum.Order.Buy, enum.OrderTypes.StopLossLimit, enum.TimeInForces.GTC, quantity, price, trailingDelta)

}

// balancePrice 平衡仓位 将持仓coin换算成cash多的取出来少的添进去
func balancePosition(latestPrice, latestPosition, accountBalance float64, coinPair string) {
	switch {
	case latestPosition > accountBalance:
		//spot.OrderTest()
		price := market.Price(coinPair).Price
		quantity := strconv.FormatFloat(calculateQuantity(accountBalance, latestPosition, latestPrice), 'e', 5, 32)
		fee, _ := strconv.ParseFloat(wallet.TradeFee(coinPair).MakerCommission, 64)
		trailingDelta := strconv.FormatFloat(fee*10000*4, 'e', 5, 32)
		spot.TrailingDeltaOrder(coinPair, enum.Order.Sell, enum.OrderTypes.TakeProfitLimit, enum.TimeInForces.GTC, quantity, price, trailingDelta)
	case latestPosition < accountBalance:
		price := market.Price(coinPair).Price
		quantity := strconv.FormatFloat(calculateQuantity(accountBalance, latestPosition, latestPrice), 'e', 5, 32)
		fee, _ := strconv.ParseFloat(wallet.TradeFee(coinPair).MakerCommission, 64)
		trailingDelta := strconv.FormatFloat(fee*10000*4, 'e', 5, 32)
		spot.TrailingDeltaOrder(coinPair, enum.Order.Buy, enum.OrderTypes.StopLossLimit, enum.TimeInForces.GTC, quantity, price, trailingDelta)
	}
}

// 换算数量
func calculateQuantity(accountBalance, latestPosition, latestPrice float64) float64 {
	delta := math.Abs(latestPosition - accountBalance)
	return delta / 2 / latestPrice
}
