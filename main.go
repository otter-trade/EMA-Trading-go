package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"EMA-Trading-go/otApi"
)

var OtterTradeKey = "xxxxxx" // 发布时 替换为  生产令牌 或者 社区令牌

var BaseUrl = "http://localhost:9342" //  发布为生产或测试时修改为   'http://api.ottertrade.com'

func UpdateVPServer() {
	//  封装一个 http 请求    http://test-api.ottertrade.com/internal/v1/update_position
}

func Start() {
	// 读取比特币价格数据
	file, err := os.Open("btc_price_data.csv") // 从 csv 读取变成 从 交易所拿最新的数据  http://test-api.ottertrade.com/open-api/market/candles
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV data:", err)
		return
	}

	// 定义均线窗口大小
	shortWindow := 20
	// longWindow := 50

	// 初始化变量
	var shortSMA, longSMA float64
	var position float64
	initialCapital := 100000.0

	// 遍历数据计算均线和生成交易信号
	for i, row := range rows {
		if i == 0 {
			continue // 跳过标题行
		}

		closePrice, _ := strconv.ParseFloat(row[4], 64) // 假设收盘价在第五列

		if i >= shortWindow {
			// 计算短期均线
			// shortSMA = simpleMovingAverage(closePrice, shortWindow, rows[i-shortWindow:i])

			// 计算长期均线
			// longSMA = simpleMovingAverage(closePrice, longWindow, rows[i-longWindow:i])

			// 生成交易信号
			if shortSMA > longSMA {
				// 买入信号
				position = initialCapital / closePrice

				UpdateVPServer() // 更新虚拟持仓

			} else if shortSMA < longSMA {
				// 卖出信号
				initialCapital = position * closePrice

				UpdateVPServer() // 更新虚拟持仓
			}
		}
	}

	// 输出回测结果
	fmt.Println("回测结束后的资本：", initialCapital)
}

func main() {
	// 周期性执行

	/*
	 setinterval


	*/
	// Start()
	otApi.StartOpenApi()
	// otApi.StartUserApi()
}

// 计算简单移动平均值
func simpleMovingAverage(currentPrice float64, window int, prices []string) float64 {
	sum := currentPrice
	for _, priceStr := range prices {
		price, _ := strconv.ParseFloat(priceStr, 64)
		sum += price
	}
	return sum / float64(window)
}
