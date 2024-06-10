package otApi

import (
	"fmt"

	"EMA-Trading-go/fetch"
)

var Header = map[string]string{
	"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IjY3MDE4ODMwN0BxcS5jb20iLCJleHAiOjE3MTgwMDgxMDYsImlhdCI6MTcxNzkyMTcwNiwidXNlcklkIjoxfQ.IloKF7HgIXSOWO9QV8KKJcwh9uMqEGUYnCSExLJ-xH8",
}

var strategyID = "5cb44aa4-1041-4601-8c32-3d98b131e4ba"

func StartOpenApi() {
	fmt.Println("StartOpenApi")
	GetMockNameList()
}

// 获取 获取历史持仓列表
func GetMockNameList() {
	data := map[string]any{
		"StrategyID": strategyID, // 策略ID 每个策略唯一，当前用户的当前策略
		"RunType":    1,
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
