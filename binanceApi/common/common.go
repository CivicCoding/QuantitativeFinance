// Package common
/*
	The common package implements a simple library for universal func
*/
package common

import (
	"QuantitativeFinance/setting"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io"
	"log"
	"net/http"
	"os"
)

// HandleResponse 处理响应
func HandleResponse(resp *http.Response) string {
	// TODO: 此处会引发性能问题需要未来优化
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Something wrong with HandleResponse", err)
		os.Exit(1)
	}

	s := string(b[:])
	return s
}

// HandleRequest 处理请求
func HandleRequest(method, url string, body io.Reader) string {

	req, err := http.NewRequest(method, url, body)
	req.Header.Add("X-MBX-APIKEY", setting.AppSetting.ApiKey)

	fmt.Println(req.URL)
	if err != nil {
		color.Red("Something wrong", err)
		os.Exit(1)
	}
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	switch {
	case resp.StatusCode == 403:
		log.Fatalf("违反WAF限制")
	case resp.StatusCode == 404:
		log.Fatalf("404 NOT FOUND")
	case resp.StatusCode == 419:
		log.Fatalf("访问频次超限，即将被封IP")
	}

	res := HandleResponse(resp)
	return res
}

// JsonStringToStruct 转换Json string 为 struct
func JsonStringToStruct(s string, stru interface{}) {
	err := json.Unmarshal([]byte(s), &stru)
	if err != nil {
		fmt.Println("Json string to struct went wrong", err)
		return
	}
}

func StringToJson(s string) []byte {
	b, err := json.Marshal(s)
	if err != nil {
		color.Red("StringToJson:", err)
	}
	return b
}

// HmacSha256 为 secretKey 加密 其他参数(除了 SecretKey 以外的参数)作为 HmacSha256 的操作对象
func HmacSha256(secretKey string, params string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(params))
	return hex.EncodeToString(h.Sum(nil))
}

// TODO : 错误处理

// HandleError
func HandleError(statusCode int) {

}

type RequestFunc struct {
}

// GetN 不需要 APIKey
func (r *RequestFunc) GetN(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	HandleError(resp.StatusCode)

	res := HandleResponse(resp)
	return res
}

func (r *RequestFunc) GetA(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("X-MBX-APIKEY", setting.AppSetting.ApiKey)
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	HandleError(resp.StatusCode)

	res := HandleResponse(resp)
	return res
}

func (r *RequestFunc) Post(url string, body io.Reader) string {
	resp, err := http.Post(url, "application/x-www-form-urlencoded", body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	HandleError(resp.StatusCode)

	res := HandleResponse(resp)
	return res
}
