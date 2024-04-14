package otApi

import (
	"fmt"
	"time"

	"EMA-Trading-go/fetch"
)

var BaseUrl = "http://test-api.ottertrade.com"

var (
	StrategyID = "mo7_EMA_Test"
	BackTestID = "mo7_local_Test1"
	RunType    = 3
)

func Test() {
	// InitPosition()
	ReadPosition()
	// UpoDatePosition()
	// MarketCandles()
}

// 第一步：初始化虚拟持仓
func InitPosition() {
	data := map[string]any{
		"StrategyID":   StrategyID, // 策略ID 每个策略唯一，当前用户的当前策略
		"RunType":      RunType,    // 运行类型 1：线上类型 2：预览类型 3：回测类型
		"BackTestID":   BackTestID, // 自定义ID 最好是用户本地生成的 UUID 或者时间戳 ,如果该ID变更，则本次虚拟账户状态重置
		"InitialAsset": 1000,       // 初始资产  10000
		"FeeRate":      0.5,        // 手续费率 0.5
	}

	res, err := fetch.Post(fetch.Opt{
		Origin: BaseUrl,
		Path:   "/internal/v1/init_position",
		Data:   data,
	})
	if err != nil {
		fmt.Println("InitPosition请求发生错误", err)
	}

	fmt.Println("InitPosition请求结果", string(res))
}

// 读取当前的虚拟持仓

func ReadPosition() {
	data := map[string]any{
		"StrategyID": "mo7_EMA_Test",    // 策略ID 每个策略唯一，当前用户的当前策略
		"RunType":    RunType,           // 运行类型 1：线上类型 2：预览类型 3：回测类型
		"BackTestID": "mo7_local_Test1", // 自定义ID 最好是用户本地生成的 UUID 或者时间戳 ,如果该ID变更，则本次虚拟账户状态重置

		"Timestamp": 0, // 查询某个时间点的持仓情况，取出当时的K线价格进行计算
	}

	res, err := fetch.Post(fetch.Opt{
		Origin: BaseUrl,
		Path:   "/internal/v1/read_position",
		Data:   data,
	})
	if err != nil {
		fmt.Println("ReadPosition 请求发生错误", err)
	}

	fmt.Println("ReadPosition 请求结果", string(res))
}

// 更新当前的虚拟持仓

type PositionType []map[string]any

func UpoDatePosition() {
	data := map[string]any{
		"StrategyID": StrategyID,
		"RunType":    RunType,
		"BackTestID": BackTestID,
		"Timestamp":  0,
		"NewPosition": PositionType{
			{
				"InstID": "BTC-USDT",
				// -1 开空  0 空仓  1 开多  & 如果 type 为1 则 -1 和 0 都默认为平仓 1 则为买入
				"Side":       -1,
				"Type":       1,         //  1：合约 0：现货
				"InstType":   "FUTURES", // 只有在 Type 为 1 的时候才生效 默认为永续合约
				"Leverage":   10,        // 只有在 Type 为 1 的时候才生效 且为必填
				"Proportion": 0.8,       // 必填，默认传 1 为当前账户的全部资金 ， 0.5 则为半仓开仓
				// "Last":       "string",
				// "Volume":     "string",
				// "AvgPx":      "string",
			},
		},
	}

	res, err := fetch.Post(fetch.Opt{
		Origin: BaseUrl,
		Path:   "/internal/v1/update_position",
		Data:   data,
	})
	if err != nil {
		fmt.Println("UpoDatePosition 请求发生错误", err)
	}

	fmt.Println("UpoDatePosition 请求结果", string(res))
}

func MarketCandles() {
	now := time.Now()
	data := map[string]any{
		"instId": "BTC-USDT",
		"bar":    "1min",
		"before": now.Unix(), // 秒
		"limit":  20,
	}

	fmt.Println(now.Unix())

	res, err := fetch.Post(fetch.Opt{
		Origin: BaseUrl,
		Path:   "/open-api/market/candles",
		Data:   data,
	})
	if err != nil {
		fmt.Println("MarketCandles 请求发生错误", err)
	}

	fmt.Println("MarketCandles 请求结果", string(res))
}
