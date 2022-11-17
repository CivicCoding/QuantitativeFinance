package strategy

import (
	"QuantitativeFinance/binanceApi/enum"
	"QuantitativeFinance/binanceApi/spot"
	"QuantitativeFinance/binanceApi/wallet"
	"log"
	"strconv"
	"time"
)

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
*/
func BalancedPosition(coinSymbol, cashSymbol string, adjustPositionTime time.Time) {
	// step1 : 获取账户现金资产
	asset := wallet.GetFundingAsset(cashSymbol)
	amount, err := strconv.ParseFloat(asset.Free, 64)
	if err != nil {
		log.Println(err)
	}
	price := strconv.FormatFloat(amount/2, 'e', 5, 32)
	// step2 : 资产的一半用来购买symbol
	spot.Order(coinSymbol+cashSymbol, enum.Order.Buy, enum.OrderTypes.LimitMaker, enum.TimeInForces.GTC, price)
	// step3 ：到达指定的时间时调整仓位,循环
}
