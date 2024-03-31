package fetch

/*
	使用方法:

	now := time.Now()
	data := map[string]any{
		"instId": "BTC-USDT",
		"bar":    "1min",
		"before": now.Unix(),
		"limit":  20,
	}

	res, err := fetch.Post(fetch.Opt{
		Origin: "http://test-api.ottertrade.com",
		Path:   "/market/candles",
		Data:   data,
	})
	if err != nil {
		fmt.Println("请求发生错误", err)
	}

	fmt.Println("请求结果", string(res))

*/

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"EMA-Trading-go/global"
)

type Opt struct {
	Origin string
	Path   string
	Data   map[string]any
}

func Post(opt Opt) ([]byte, error) {
	// 请求参数转换
	targetUrl := opt.Origin + opt.Path
	dataStr := MapToJson(opt.Data)
	payload := strings.NewReader(dataStr)

	// 创建请求
	req, err := http.NewRequest("POST", targetUrl, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// 解析返回结果
	body, _ := io.ReadAll(response.Body)
	return body, nil
}

func Get(opt Opt) ([]byte, error) {
	parseURL, err := url.Parse(opt.Origin + opt.Path)
	if err != nil {
		return nil, err
	}

	params := url.Values{}

	for k, v := range opt.Data {
		params.Set(k, ToStr(v))
	}
	// 如果参数中有中文参数,这个方法会进行URLEncode
	parseURL.RawQuery = params.Encode()
	targetUrl := parseURL.String()

	global.Log.Println("发出请求: ", targetUrl)

	// 创建请求
	req, err := http.NewRequest("GET", targetUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// 解析返回结果
	body, _ := io.ReadAll(response.Body)
	return body, nil
}
