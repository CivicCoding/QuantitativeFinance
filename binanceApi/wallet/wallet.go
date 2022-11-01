package wallet

import (
	"QuantitativeFinance/binanceApi/common"
	"QuantitativeFinance/binanceApi/market"
	"QuantitativeFinance/setting"
	"fmt"
	"github.com/fatih/color"
	"net/http"
	"os"
	"strconv"
)

// SystemStatus "/sapi/v1/system/status"
func SystemStatus(baseUrl string) {

	resp, err := http.Get(baseUrl + "/sapi/v1/system/status")
	if err != nil {
		color.Red("Something wrong with system status", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	res := common.HandleResponse(resp)
	fmt.Println(res)
}

// AccountSnapshot 获取用户当日资产快照 kind: "SPOT","MARGIN","FUTURES"
func AccountSnapshot(baseUrl string, kind string, days string) {
	timeStamp := strconv.Itoa(market.ServerTime(baseUrl))
	param := "type=" + kind + "&limit=" + days + "&recvWindow=5000" + "&timestamp=" + timeStamp
	signature := common.HmacSha256(setting.AppSetting.SecreteKey, param)
	url := baseUrl + "/sapi/v1/accountSnapshot?type=" + kind + "&limit=" + days + "&recvWindow=5000" + "&timestamp=" + timeStamp + "&signature=" + signature
	// use original way to request an url
	res := common.HandleRequest("GET", url, setting.AppSetting.ApiKey)
	color.Green(res)
}

// GetAll /sapi/v1/capital/config/getall
func GetAll(baseUrl string) {
	timeStamp := strconv.Itoa(market.ServerTime(baseUrl))
	param := "recvWindow=5000" + "&timestamp=" + timeStamp
	signature := common.HmacSha256(setting.AppSetting.SecreteKey, param)
	url := baseUrl + "/sapi/v1/capital/config/getall?recvWindow=5000" + "&timestamp=" + timeStamp + "&signature=" + signature
	res := common.HandleRequest("GET", url, setting.AppSetting.ApiKey)
	color.Green(res)
}
