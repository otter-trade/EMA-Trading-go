package main

import (
	"fmt"

	"EMA-Trading-go/okx"
)

func main() {
	resData, err := okx.GetOKXCandle(okx.GetCandleOpt{
		InstID: "BTC-USDT",
		Bar:    "1H",
	})
	if err != nil {
		fmt.Println("出现错误", err)
	}

	fmt.Println("结果", resData)
}
