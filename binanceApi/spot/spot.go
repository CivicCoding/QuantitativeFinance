package spot

import (
	"QuantitativeFinance/binanceApi/common"
	"QuantitativeFinance/binanceApi/market"
	"QuantitativeFinance/model"
	"QuantitativeFinance/setting"
	"net/url"
	"strconv"
	"strings"
)

// OrderTest 测试下单 /api/v3/order/test
func OrderTest() {

}

/*
Order 下单/api/v3/order (HMAC SHA256) symbol 交易币对, side 交易方向, kind 交易类型, timeInForce 订单多久会失效
,price 交易价格 TODO: 需要区分不同交易方向 和 交易类别
*/
func Order(symbol, side, kind, timeInForce string, price string) model.OrderInfo {
	var r common.RequestFunc
	route := setting.AppSetting.Url + "/api/v3/order"

	config := url.Values{}
	config.Add("symbol", symbol)
	config.Add("side", side)
	config.Add("type", kind)
	config.Add("timeInForce", timeInForce)
	config.Add("price", price)
	params := config.Encode()
	signature := common.HmacSha256(setting.AppSetting.SecreteKey, params)
	config.Add("signature", signature)
	data := config.Encode()
	res := r.Post(route, strings.NewReader(data))
	var order model.OrderInfo
	common.JsonStringToStruct(res, order)
	return order
}

/*
TrailingDeltaOrder 追踪止盈止损(Trailing Stop)订单 /api/v3/order (HMAC SHA256)
详见 https://github.com/binance/binance-spot-api-docs/blob/master/faqs/trailing-stop-faq-cn.md
quantity : 买入数量
price : 买入价格
stopPrice : 止损价 当市场价格到达stopPrice时订单从这里开始追踪价格变动(在找到计算此价格方法前不设置)
trailingDelta : 基点 一般价格变动范围设置在手续费的4倍左右为基准开始调优较为合理
*/
func TrailingDeltaOrder(symbol, side, orderType, timeInForce, quantity, price, trailingDelta string) model.OrderInfo {
	route := setting.AppSetting.Url + "/api/v3/order"
	timeStamp := strconv.Itoa(market.ServerTime())
	config := url.Values{}
	config.Add("symbol", symbol)
	config.Add("side", side)
	config.Add("type", orderType)
	config.Add("timeInForce", timeInForce)
	//config.Add("stopPrice", stopPrice)
	config.Add("trailingDelta", trailingDelta)
	config.Add("timestamp", timeStamp)
	config.Add("quantity", quantity)
	config.Add("price", price)
	config.Add("newOrderRespType", "FULL")
	params := config.Encode()
	signature := common.HmacSha256(setting.AppSetting.SecreteKey, params)
	config.Add("signature", signature)
	data := config.Encode()
	var r common.RequestFunc
	res := r.Post(route, strings.NewReader(data))
	var order model.OrderInfo
	common.JsonStringToStruct(res, order)
	return order
}
