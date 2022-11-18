package spot

import (
	"QuantitativeFinance/binanceApi/common"
	"QuantitativeFinance/model"
	"QuantitativeFinance/setting"
	"net/url"
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
