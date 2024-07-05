package otApi

import (
	"fmt"

	"EMA-Trading-go/fetch"
)

var Header = map[string]string{
	"Authorization": Authorization,
}

var strategyID = "e395ffe2-3a99-4072-abfc-0da940f32f26"

func StartUserApi() {
	fmt.Println("StartUserApi")
	// GetMockNameList()
	// GetHistoryStatus()
}

// 获取 获取历史持仓列表
func GetMockNameList() {
	data := map[string]any{
		"StrategyID": strategyID, // 策略ID 每个策略唯一，当前用户的当前策略
		"RunType":    1,          // 1：回测类型 2：预览类型 3：线上类型
	}

	res, err := fetch.Post(fetch.Opt{
		Origin: BaseUrl,
		Path:   "/api/user/position/scan",
		Header: Header,
		Data:   data,
	})
	if err != nil {
		fmt.Println("InitPosition请求发生错误", err)
	}

	fmt.Println("InitPosition请求结果", string(res))
}

func GetHistoryStatus() {
	data := map[string]any{
		"StrategyID": strategyID, // 策略ID 每个策略唯一，当前用户的当前策略
		"RunType":    1,          // 1：回测类型 2：预览类型 3：线上类型
		"MockName":   MockName,   // MockName 本次回测的名字
	}

	res, err := fetch.Post(fetch.Opt{
		Origin: BaseUrl,
		Path:   "/api/user/position/history/read",
		Header: Header,
		Data:   data,
	})
	if err != nil {
		fmt.Println("InitPosition请求发生错误", err)
	}

	fmt.Println("InitPosition请求结果", string(res))
}
