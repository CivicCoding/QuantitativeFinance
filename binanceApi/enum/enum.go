package enum

type OrderSide struct {
	Buy  string
	Sell string
}

var Order = &OrderSide{
	Buy:  "BUY",
	Sell: "SELL",
}

// OrderType 订单类型
type OrderType struct {
	LIMIT           string //限价单
	MARKET          string //市价单
	StopLoss        string //止损单
	StopLossLimit   string //限价止损单
	TakeProfit      string //止盈单
	TakeProfitLimit string //限价止盈单
	LimitMaker      string //限价只挂单
}

var OrderTypes = &OrderType{
	LIMIT:           "LIMIT",
	MARKET:          "MARKET",
	StopLoss:        "STOP_LOSS",
	StopLossLimit:   "STOP_LOSS_LIMIT",
	TakeProfit:      "TAKE_PROFIT",
	TakeProfitLimit: "TAKE_PROFIT_LIMIT",
	LimitMaker:      "LIMIT_MAKER",
}

// TimeInForce 定义了订单多久能够失效
type TimeInForce struct {
	GTC string //成交为止 订单会一直有效，直到被成交或者取消
	IOC string //无法立即成交的部分就撤销 订单在失效前会尽量多的成交。
	FOK string //无法全部立即成交就撤销 如果无法全部成交，订单会失效。
}

var TimeInForces = &TimeInForce{
	GTC: "GTC",
	IOC: "IOC",
	FOK: "FOK",
}
