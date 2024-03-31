package main

import (
	"fmt"
	"time"

	"EMA-Trading-go/fetch"
)

func main() {
	now := time.Now()
	data := map[string]any{
		"instId": "BTC-USDT",
		"bar":    "1min",
		"before": now.Unix(),
		"limit":  20,
	}

	res, err := fetch.Post(fetch.Opt{
		Origin: "http://test-api.ottertrade.com",
		Path:   "/market/candles?abc=123",
		Data:   data,
	})
	if err != nil {
		fmt.Println("请求发生错误", err)
	}

	fmt.Println("请求结果", string(res))
}
