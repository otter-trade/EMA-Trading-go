package otApi

import (
	"fmt"

	"EMA-Trading-go/fetch"
)

func GetHeaderACCESS() map[string]string {
	OpenApiHeader := map[string]string{
		"OT-ACCESS-KEY":       "1",
		"OT-ACCESS-SIGN":      "2",
		"OT-ACCESS-TIMESTAMP": "3",
	}
	return OpenApiHeader
}

func StartOpenApi() {
	fmt.Println("StartOpenApi")
	InitPosition()
}

// 初始化持仓
func InitPosition() {
	data := map[string]any{
		"MockName": "mo7_test_first", // 策略ID 每个策略唯一，当前用户的当前策略
		// "InitialAsset": "10000",          // 初始资产 默认 10000
	}

	res, err := fetch.Post(fetch.Opt{
		Origin: BaseUrl,
		Path:   "/openapi/position/init",
		Header: GetHeaderACCESS(),
		Data:   data,
	})
	if err != nil {
		fmt.Println("InitPosition请求发生错误", err)
	}

	fmt.Println("InitPosition请求结果", string(res))
}

// 更新持仓
func UpdatePosition() {
}

// 读取当前持仓状态
func RedPosition() {
}

// 读取历史持仓状态 openapi/position/history/read
func RedHistoryPosition() {
}
