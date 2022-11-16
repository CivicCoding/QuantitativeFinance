package strategy

// TODO:设置失效点退出策略

/*
Grid
网格策略，根据设定的 lowPrice 和 highPrice 分布等差或者等比分布价格网格，
当价格向下突破网格区间执行买单，同时在价格区间上部署卖单。
价格向上突破价格区间，执行卖单，同时在区间下部署买单。
当价格突破设定的价格区间时停止网格策略并平仓。
*/
func Grid(lowPrice, highPrice int, symbol, strategy string) {

}

// Isochromatic low - high 范围的分成 phr 份的等差数列的价差
func Isochromatic(low, high, phr float64) float64 {
	// 买卖价差
	baSpread := (high - low) / phr
	return baSpread
}
