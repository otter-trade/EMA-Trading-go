package main

import (
	"fmt"

	"EMA-Trading-go/global"
	"EMA-Trading-go/okx"
	"EMA-Trading-go/trade"
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
	global.Log.Println("系统启动.....")

	// 新建一个策略
	tradeObj := trade.New()

	// 填充基础数据
	tradeObj.FillBaseCandle()
	// 填充最新的数据
	tradeObj.SetNowCandle()
}
