package wallet

import (
	"QuantitativeFinance/binanceApi/common"
	"QuantitativeFinance/binanceApi/market"
	"QuantitativeFinance/setting"
	"fmt"
	"github.com/fatih/color"
	"log"
	url2 "net/url"
	"strconv"
	"strings"
)

// SystemStatus "/sapi/v1/system/status"
func SystemStatus() string {
	var r common.RequestFunc
	return r.GetN(setting.AppSetting.Url + "/sapi/v1/system/status")
}

type AccountStatus struct {
	Data string `json:"data"`
}

// Status 账户状态 /sapi/v1/account/status
func Status() string {
	var r common.RequestFunc
	s := r.GetN(setting.AppSetting.Url + "/sapi/v1/account/status")
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
	timeStamp := strconv.Itoa(market.ServerTime())
	param := "type=" + kind + "&limit=" + days + "&recvWindow=5000" + "&timestamp=" + timeStamp
	signature := common.HmacSha256(setting.AppSetting.SecreteKey, param)
	url := setting.AppSetting.Url + "/sapi/v1/accountSnapshot?type=" + kind + "&limit=" + days + "&recvWindow=5000" + "&timestamp=" + timeStamp + "&signature=" + signature
	// use original way to request an url
	res := common.HandleRequest("GET", url, nil)
	var ab AccountBalance
	common.JsonStringToStruct(res, &ab)
	color.Green("%+v", ab)
}

type Asset struct {
	Asset        string `json:"asset"`
	Free         string `json:"free"`
	Locked       string `json:"locked"`
	Freeze       string `json:"freeze"`
	Withdrawing  string `json:"withdrawing"`
	BtcValuation string `json:"btcValuation"`
}

// GetFundingAsset /sapi/v1/asset/get-funding-asset
func GetFundingAsset(asset string) Asset {
	timeStamp := strconv.Itoa(market.ServerTime())
	var r common.RequestFunc
	config := url2.Values{}
	config.Add("asset", asset)
	config.Add("timestamp", timeStamp)
	config.Add("needBtcValuation", "false")
	params := config.Encode()
	url := setting.AppSetting.Url + "/sapi/v1/asset/get-funding-asset"
	a := r.Post(url, strings.NewReader(params))
	var as Asset
	common.JsonStringToStruct(a, &as)
	return as
}

/*
GetAll 获取所有币信息 /sapi/v1/capital/config/getall
获取针对用户的所有(Binance支持充提操作的)币种信息
*/
func GetAll() {
	timeStamp := strconv.Itoa(market.ServerTime())
	param := "recvWindow=5000" + "&timestamp=" + timeStamp
	signature := common.HmacSha256(setting.AppSetting.SecreteKey, param)
	url := setting.AppSetting.Url + "/sapi/v1/capital/config/getall?recvWindow=5000" + "&timestamp=" + timeStamp + "&signature=" + signature
	res := common.HandleRequest("GET", url, nil)
	color.Green(res)
}

// DisableFastWithdrawSwitch 关闭站内划转 /sapi/v1/account/disableFastWithdrawSwitch
func DisableFastWithdrawSwitch() {
	timeStamp := strconv.Itoa(market.ServerTime())
	param := "recvWindow=5000" + "&timestamp=" + timeStamp
	signature := common.HmacSha256(setting.AppSetting.SecreteKey, param)
	url := setting.AppSetting.Url + "/sapi/v1/account/disableFastWithdrawSwitch"
	configInfo := url2.Values{}
	configInfo.Add("recvWindow", "5000")
	configInfo.Add("timestamp", timeStamp)
	configInfo.Add("signature", signature)
	data := configInfo.Encode()
	var r common.RequestFunc
	res := r.Post(url, strings.NewReader(data))
	if res != "" {
		color.Red("关闭站内划转失败！", res)
	} else {
		color.Green("关闭站内划转成功！")
	}

}

// EnableFastWithdrawSwitch 开启站内划转 /sapi/v1/account/enableFastWithdrawSwitch (HMAC SHA256)
func EnableFastWithdrawSwitch() {
	timeStamp := strconv.Itoa(market.ServerTime())
	param := "recvWindow=5000" + "&timestamp=" + timeStamp
	signature := common.HmacSha256(setting.AppSetting.SecreteKey, param)
	url := setting.AppSetting.Url + "/sapi/v1/account/enableFastWithdrawSwitch"
	configInfo := url2.Values{}
	configInfo.Add("recvWindow", "5000")
	configInfo.Add("timestamp", timeStamp)
	configInfo.Add("signature", signature)
	data := configInfo.Encode()
	var r common.RequestFunc
	res := r.Post(url, strings.NewReader(data))
	if res != "" {
		log.Fatalln("开启站内划转失败", res)
	}
}

// TradeFeeInfo 交易费率信息
type TradeFeeInfo struct {
	Symbol          string `json:"symbol"`
	MakerCommission string `json:"makerCommission"`
	TakerCommission string `json:"takerCommission"`
}

// TradeFee 获取交易费率 /sapi/v1/asset/tradeFee
func TradeFee(symbol string) {
	timeStamp := strconv.Itoa(market.ServerTime())
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

// GetUserAsset [POST] 获取用户持仓，仅返回>0的数据。 /sapi/v3/asset/getUserAsset
func GetUserAsset(asset string, needBtcValuation bool) string {
	timeStamp := strconv.Itoa(market.ServerTime())
	baseUrl := setting.AppSetting.Url + "/sapi/v1/asset/tradeFee"
	configInfo := url2.Values{}
	configInfo.Add("asset", asset)
	configInfo.Add("needBtcValuation", strconv.FormatBool(needBtcValuation))
	configInfo.Add("recvWindow", "5000")
	configInfo.Add("timestamp", timeStamp)
	data := configInfo.Encode()
	var r common.RequestFunc
	res := r.Post(baseUrl, strings.NewReader(data))
	return res
}
