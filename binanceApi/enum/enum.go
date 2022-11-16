package enum

type OrderSide struct {
	Buy  string
	Sell string
}

var Order = &OrderSide{
	Buy:  "BUY",
	Sell: "SELL",
}
