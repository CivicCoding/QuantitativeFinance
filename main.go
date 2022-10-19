package main

import (
	"BinanceApi/market"
	"fmt"
)

const baseURL = "https://api.binance.com"

const secretKey = "91IW45xj7g8GCEjtYMievrfR0n7sFUO0mQWKiSL0R3HYi7p6RwxRChmXapex2z5R"
const apiKey = "zQDyK5C0dQcymyim7y3jyDLGj6rYQqNqfpDffEL8Ojw1uhzgHkf95hPIkNl9e1UX"

func main() {

	//wallet.GetAll(baseURL)
	//market.ServerTime(baseURL)
	//wallet.GetAll(baseURL, apiKey, secretKey)
	klineData := market.Kline(baseURL, "ETHBUSD")
	for _, v := range *klineData {
		fmt.Println(v.OpeningTime)
	}

}
