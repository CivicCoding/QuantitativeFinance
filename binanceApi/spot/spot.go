package spot

import (
	"QuantitativeFinance/binanceApi/common"
	"QuantitativeFinance/setting"
	"net/url"
	"strings"
)

// OrderTest 测试下单 /api/v3/order/test
func OrderTest() {

}

// Order 下单/api/v3/order (HMAC SHA256)
func Order(symbol, side, kind, timeInForce string) string {
	var r common.RequestFunc
	route := setting.AppSetting.Url + "/api/v3/order"
	config := url.Values{}
	config.Add("symbol", symbol)
	config.Add("side", side)
	config.Add("kind", kind)
	config.Add("timeInForce", timeInForce)
	data := config.Encode()
	res := r.Post(route, strings.NewReader(data))
	return res
}
