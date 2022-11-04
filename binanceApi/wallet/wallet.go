package wallet

import (
	"QuantitativeFinance/binanceApi/common"
	"QuantitativeFinance/binanceApi/market"
	"QuantitativeFinance/setting"
	"fmt"
	"github.com/fatih/color"
	"net/http"
	url2 "net/url"
	"os"
	"strconv"
	"strings"
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

type AccountStatus struct {
	Data string `json:"data"`
}

// Status 账户状态 /sapi/v1/account/status
func Status(baseUrl string) string {
	resp, err := http.Get(baseUrl + "/sapi/v1/account/status")
	if err != nil {
		color.Red("get account status error", err)
	}
	s := common.HandleResponse(resp)
	var as AccountStatus
	common.JsonStringToStruct(s, &as)
	return as.Data
}

type AccountBalance struct {
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	SnapshotVos []struct {
		Data struct {
			Balances []struct {
				Asset  string `json:"asset"`
				Free   string `json:"free"`
				Locked string `json:"locked"`
			} `json:"balances"`
			TotalAssetOfBtc string `json:"totalAssetOfBtc"`
		} `json:"data"`
		Type       string `json:"type"`
		UpdateTime int64  `json:"updateTime"`
	} `json:"snapshotVos"`
}

/*
AccountSnapshot 获取用户当日资产快照 /sapi/v1/accountSnapshot
kind: "SPOT","MARGIN","FUTURES"
*/
func AccountSnapshot(kind string, days string) {
	timeStamp := strconv.Itoa(market.ServerTime(setting.AppSetting.Url))
	param := "type=" + kind + "&limit=" + days + "&recvWindow=5000" + "&timestamp=" + timeStamp
	signature := common.HmacSha256(setting.AppSetting.SecreteKey, param)
	url := setting.AppSetting.Url + "/sapi/v1/accountSnapshot?type=" + kind + "&limit=" + days + "&recvWindow=5000" + "&timestamp=" + timeStamp + "&signature=" + signature
	// use original way to request an url
	res := common.HandleRequest("GET", url, nil)
	var ab AccountBalance
	common.JsonStringToStruct(res, &ab)
	color.Green("%+v", ab)
}

/*
GetAll 获取所有币信息 /sapi/v1/capital/config/getall
获取针对用户的所有(Binance支持充提操作的)币种信息
*/
func GetAll() {
	timeStamp := strconv.Itoa(market.ServerTime(setting.AppSetting.Url))
	param := "recvWindow=5000" + "&timestamp=" + timeStamp
	signature := common.HmacSha256(setting.AppSetting.SecreteKey, param)
	url := setting.AppSetting.Url + "/sapi/v1/capital/config/getall?recvWindow=5000" + "&timestamp=" + timeStamp + "&signature=" + signature
	res := common.HandleRequest("GET", url, nil)
	color.Green(res)
}

// DisableFastWithdrawSwitch 关闭站内划转 /sapi/v1/account/disableFastWithdrawSwitch
func DisableFastWithdrawSwitch() {
	timeStamp := strconv.Itoa(market.ServerTime(setting.AppSetting.Url))
	param := "recvWindow=5000" + "&timestamp=" + timeStamp
	signature := common.HmacSha256(setting.AppSetting.SecreteKey, param)
	url := setting.AppSetting.Url + "/sapi/v1/account/disableFastWithdrawSwitch"
	configInfo := url2.Values{}
	configInfo.Add("recvWindow", "5000")
	configInfo.Add("timestamp", timeStamp)
	configInfo.Add("signature", signature)
	data := configInfo.Encode()
	res := common.HandleRequest("POST", url, strings.NewReader(data))
	if res != "" {
		color.Red("关闭站内划转失败！", res)
	} else {
		color.Green("关闭站内划转成功！")
	}

}

// EnableFastWithdrawSwitch 开启站内划转 /sapi/v1/account/enableFastWithdrawSwitch (HMAC SHA256)
func EnableFastWithdrawSwitch() {
	timeStamp := strconv.Itoa(market.ServerTime(setting.AppSetting.Url))
	param := "recvWindow=5000" + "&timestamp=" + timeStamp
	signature := common.HmacSha256(setting.AppSetting.SecreteKey, param)
	url := setting.AppSetting.Url + "/sapi/v1/account/enableFastWithdrawSwitch"
	configInfo := url2.Values{}
	configInfo.Add("recvWindow", "5000")
	configInfo.Add("timestamp", timeStamp)
	configInfo.Add("signature", signature)
	data := configInfo.Encode()
	var r common.RequestFunc
	r.Post(url, strings.NewReader(data))
}

// TradeFeeInfo 交易费率信息
type TradeFeeInfo struct {
	Symbol          string `json:"symbol"`
	MakerCommission string `json:"makerCommission"`
	TakerCommission string `json:"takerCommission"`
}

// TradeFee 获取交易费率 /sapi/v1/asset/tradeFee
func TradeFee(symbol string) {
	timeStamp := strconv.Itoa(market.ServerTime(setting.AppSetting.Url))
	param := "symbol=" + symbol + "&recvWindow=5000" + "&timestamp=" + timeStamp
	signature := common.HmacSha256(setting.AppSetting.SecreteKey, param)
	url := setting.AppSetting.Url + "/sapi/v1/asset/tradeFee?" + "symbol=" + symbol + "&recvWindow=5000" + "&timestamp=" + timeStamp + "&signature=" + signature
	//res := common.HandleRequest("GET", url, nil)
	var r common.RequestFunc
	res := r.GetA(url)
	var td TradeFeeInfo
	common.JsonStringToStruct(res, &td)
	fmt.Printf("%+v\n%s", td, res)
}
