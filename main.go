package main

import (
	"fmt"

	"EMA-Trading-go/global"
	"EMA-Trading-go/okx"
	"EMA-Trading-go/trade"

	"EMA-Trading-go/mClock"
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

	// 定时任务走起
	Running(tradeObj)
	go mClock.New(mClock.OptType{
		Func: func() {
			Running(tradeObj)
		},
		Spec: "1 1,6,11,16,21,26,31,36,41,46,51,56 * * * ? ", // 每隔5分钟比标准时间晚一分钟 过 1 秒执行一次
	})

	fmt.Println("当前服务正在执行中.......")
	select {}
}

func Running(tradeObj *trade.TradeObj) {
	// 填充当前的最新数据
	tradeObj.SetNowCandle()
}
