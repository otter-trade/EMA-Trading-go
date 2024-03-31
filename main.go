package main

import (
	"fmt"

	"EMA-Trading-go/global"
	"EMA-Trading-go/okx"
)

func GetCandle() {
	resData, err := okx.GetOKXCandle(okx.GetCandleOpt{
		InstID: "BTC-USDT",
		Bar:    "1H",
	})
	if err != nil {
		fmt.Println("出现错误", err)
	}
	for k, v := range resData {
		fmt.Println(k, v)
	}
}

func main() {
	// 初始化全局参数
	global.Start()

	// GetCandle()

	global.Log.Println("打印日志")
}
