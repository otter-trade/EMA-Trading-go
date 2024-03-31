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
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Opt struct {
	Origin string
	Path   string
	Data   map[string]any
}

func Post(opt Opt) ([]byte, error) {
	// 地址拼接
	targetUrl := opt.Origin + opt.Path

	// 请求参数转换
	dataStr := MapToJson(opt.Data)
	payload := strings.NewReader(dataStr)

	// 创建请求
	fmt.Println("targetUrl", targetUrl)

	req, _ := http.NewRequest("POST", targetUrl, payload)

	// json 格式
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	return body, nil
}
