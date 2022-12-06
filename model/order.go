package model

import "QuantitativeFinance/dbService"

type OrderInfo struct {
	Symbol              string `json:"symbol"`              // 交易对
	OrderId             int    `json:"orderId"`             // 系统的订单ID
	OrderListId         int    `json:"orderListId"`         // OCO订单ID，否则为 -1
	ClientOrderId       string `json:"clientOrderId"`       // 客户自己设置的ID
	TransactTime        int64  `json:"transactTime"`        // 交易的时间戳
	Price               string `json:"price"`               // 订单价格
	OrigQty             string `json:"origQty"`             // 用户设置的原始订单数量
	ExecutedQty         string `json:"executedQty"`         // 交易的订单数量
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"` // 累计交易的金额
	Status              string `json:"status"`              // 订单状态
	TimeInForce         string `json:"timeInForce"`         // 订单的时效方式
	Type                string `json:"type"`                // 订单类型， 比如市价单，现价单等
	Side                string `json:"side"`                // 订单方向，买还是卖
	StrategyId          int    `json:"strategyId"`          // 下单填了参数才会返回
	StrategyType        int    `json:"strategyType"`        // 下单填了参数才会返回
	Fills               []Fill `json:"fills"`
}

type Fill struct {
	Price           string `json:"price"`           // 交易的价格
	Qty             string `json:"qty"`             // 交易的数量
	Commission      string `json:"commission"`      // 手续费金额
	CommissionAsset string `json:"commissionAsset"` // 手续费的币种
	TradeId         int    `json:"tradeId"`         // 交易ID
}

func InitOrder() {
	dbService.GetDB().CreateTable(&OrderInfo{})
	dbService.GetDB().CreateTable(&Fill{})
}
