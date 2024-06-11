package otApi

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"

	"EMA-Trading-go/fetch"
	"EMA-Trading-go/mJson"
	"EMA-Trading-go/mStr"
)

/*

apiKey
:
"mJsp2X90ltkBNENFh799resyud3UqhovjY5iUgpKWLBMRSNMohWjrvt9kWQanAb5"
content
:
"e9ef03e8-d611-431b-8227-b8f15fa07af0"


SignStr := mStr.Join(
    Timestamp,   // 开发者本地的时间戳，发起请求生成的
    strings.ToUpper("get"), //开发者用什么类型的请求，他就写什么，不过我们只有 POST ， 可省略
    Path, //  开发着请求哪个接口，他就得在这里写什么地址
    Body, // 他发出请求时的参数字符串
  )

*/

var MockName = "mo7_test_first"

func HmacSha256(key string, data string) string {
	mac := hmac.New(sha256.New, []byte(key))
	_, _ = mac.Write([]byte(data))

	hex := mac.Sum(nil)

	hexStr := base64.URLEncoding.EncodeToString(hex)

	return hexStr
}

func GetHeaderACCESS(path string, body string) map[string]string {
	timeUnix := time.Now().Unix()

	fmt.Println("timeUnix", timeUnix)

	SignStr := mStr.Join(
		timeUnix, // 开发者本地的时间戳，发起请求生成的
		path,     //  开发着请求哪个接口，他就得在这里写什么地址
		body,     // 他发出请求时的参数字符串
	)
	fmt.Println("SignStr", SignStr)

	Sign := HmacSha256(OtStrategyToken, SignStr)

	fmt.Println("Sign", Sign)

	OpenApiHeader := map[string]string{
		"OT-ACCESS-KEY":       OtAPIKey,
		"OT-ACCESS-SIGN":      Sign,
		"OT-ACCESS-TIMESTAMP": mStr.ToStr(timeUnix), // 2006-01-02T15:04:05Z07:00
	}
	return OpenApiHeader
}

func StartOpenApi() {
	fmt.Println("StartOpenApi")
	// InitPosition()
	// UpdatePosition()
	// RedPosition()
	RedHistoryPosition()
}

// 初始化持仓
func InitPosition() {
	data := map[string]any{
		"MockName":     MockName, // 策略ID 每个策略唯一，当前用户的当前策略
		"InitialAsset": 10000,    // 初始资产 默认 10000
	}

	path := "/openapi/position/init"
	res, err := fetch.Post(fetch.Opt{
		Origin: BaseUrl,
		Path:   path,
		Header: GetHeaderACCESS(path, mJson.ToStr(data)),
		Data:   data,
	})
	if err != nil {
		fmt.Println("InitPosition 请求发生错误", err)
	}

	fmt.Println("InitPosition 请求结果", string(res))

	/*
	 {"code":200,"msg":"成功","data":{"Position":{"StrategyID":"e395ffe2-3a99-4072-abfc-0da940f32f26","RunType":1,"InstanceID":"backtest_mo7_test_first","InitialAsset":10000,"FeeRate":0.01,"Cash":10000,"CurrPosition":null,"TimetampNs":0,"InitTimetampNs":1718108206624818400}}}
	*/
}

// 更新持仓
func UpdatePosition() {

	// 更新了一次持仓， BTC-USDT 合约 2x 做空
	// data := map[string]any{
	// 	"MockName":  MockName,   // 策略ID 每个策略唯一，当前用户的当前策略
	// 	"Timestamp": 1717826400, // 时间戳   2024-06-08 14:00:00
	// 	"NewPosition": PositionType{
	// 		{
	// 			"InstID":     "BTC-USDT", // 交易产品ID
	// 			"InstType":   "FUTURES",  // 产品类型
	// 			"Leverage":   2,          // 杠杆倍数
	// 			"Side":       1,          // 买卖方向 1 空 2 多  &  <应该改成  -1 卖空 1 买多 0 空仓>
	// 			"Type":       2,          // 1:现货, 2:合约
	// 			"Proportion": 0.8,        // 持仓比例
	// 		},
	// 	},
	// }

	// BTC 平仓
	data := map[string]any{
		"MockName":  MockName,   // 策略ID 每个策略唯一，当前用户的当前策略
		"Timestamp": 1717826400, // 时间戳   2024-06-08 14:00:00
		"NewPosition": PositionType{
			{
				"InstID":     "BTC-USDT", // 交易产品ID
				"InstType":   "FUTURES",  // 产品类型
				"Leverage":   2,          // 杠杆倍数
				"Side":       1,          // 买卖方向 1 空 2 多  &  <应该改成  -1 卖空 1 买多 0 空仓>
				"Type":       2,          // 1:现货, 2:合约
				"Proportion": 0,          // 持仓比例
			},
		},
	}

	path := "/openapi/position/update"
	res, err := fetch.Post(fetch.Opt{
		Origin: BaseUrl,
		Path:   path,
		Header: GetHeaderACCESS(path, mJson.ToStr(data)),
		Data:   data,
	})
	if err != nil {
		fmt.Println("UpdatePosition 请求发生错误", err)
	}

	fmt.Println("UpdatePosition 请求结果", string(res))
	/*
		{"code":200,"msg":"成功","data":{}}
	*/
}

// 读取当前持仓状态
func RedPosition() {
	data := map[string]any{
		"MockName":  MockName,   // 策略ID 每个策略唯一，当前用户的当前策略
		"Timestamp": 1717830000, // 时间戳   2024-06-08 15:00:00  // 这里有时间戳则应该是读取当前时间戳的仓位变化  如果时间戳为0 则读取 当下这一刻的价格计算
	}

	path := "/openapi/position/read"
	res, err := fetch.Post(fetch.Opt{
		Origin: BaseUrl,
		Path:   path,
		Header: GetHeaderACCESS(path, mJson.ToStr(data)),
		Data:   data,
	})
	if err != nil {
		fmt.Println("RedPosition 请求发生错误", err)
	}

	fmt.Println("RedPosition 请求结果", string(res))

	/*

		{"code":200,"msg":"成功","data":{"Position":{"StrategyID":"e395ffe2-3a99-4072-abfc-0da940f32f26","RunType":1,"InstanceID":"backtest_mo7_test_first","InitialAsset":10000,"FeeRate":0.01,"Cash":2000,"CurrPosition":[{"InstID":"BTC-USDT","InstType":"FUTURES","Leverage":2,"Side":1,"Type":2,"Proportion":0.8,"Last":"69336.9","Volume":"0.1142827968091729","AvgPx":"70001.787","Margin":"70001.787","ForcedClosedOut":true}],"TimetampNs":1718108999434701800,"InitTimetampNs":1718108206624818400}}}
	*/
}

// 读取历史持仓状态 openapi/position/history/read
func RedHistoryPosition() {
	data := map[string]any{
		"MockName": MockName, // 策略ID 每个策略唯一，当前用户的当前策略
	}

	path := "/openapi/position/history/read"
	res, err := fetch.Post(fetch.Opt{
		Origin: BaseUrl,
		Path:   path,
		Header: GetHeaderACCESS(path, mJson.ToStr(data)),
		Data:   data,
	})
	if err != nil {
		fmt.Println("RedHistoryPosition 请求发生错误", err)
	}

	fmt.Println("RedHistoryPosition 请求结果", string(res))

}
