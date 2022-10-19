// Package common
/*
	The common package implements a simple library for universal func
*/
package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io"
	"net/http"
	"os"
)

type base struct {
	url       string
	apiKey    string
	secretKey string
}

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
func HandleRequest(method, url, apiKey string) string {
	req, err := http.NewRequest(method, url, nil)
	req.Header.Add("X-MBX-APIKEY", apiKey)

	if err != nil {
		color.Red("Something wrong", err)
		os.Exit(1)
	}
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	res := HandleResponse(resp)
	return res
}

// JsonStringToStruct 转换Json string 为 struct
func JsonStringToStruct(s string, stru interface{}) {
	err := json.Unmarshal([]byte(s), &stru)
	if err != nil {
		fmt.Println("Json string to struct went wrong", err)
		os.Exit(1)
	}
}

// HmacSha256 为 secretKey 加密 其他参数(除了 SecretKey 以外的参数)作为 HmacSha256 的操作对象
func HmacSha256(secretKey string, params string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(params))
	return hex.EncodeToString(h.Sum(nil))
}
