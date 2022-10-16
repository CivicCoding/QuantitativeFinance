package market

import (
	"BinanceApi/common"
	"github.com/fatih/color"
	"net/http"
	"os"
)

type timeS struct {
	ServerTime int `json:"ServerTime"`
}

// ServerTime 获取Binance服务器时间 api/v3/time
func ServerTime(baseUrl string) int {
	resp, err := http.Get(baseUrl + "/api/v3/time")
	if err != nil {
		color.Red("Something wrong with time \n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	s := common.HandleResponse(resp)

	var t timeS

	common.JsonStringToStruct(s, &t)

	return t.ServerTime
}
