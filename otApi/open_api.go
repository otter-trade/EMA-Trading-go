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

func HmacSha256(key string, data string) string {
	mac := hmac.New(sha256.New, []byte(key))
	_, _ = mac.Write([]byte(data))

	hex := mac.Sum(nil)

	hexStr := base64.URLEncoding.EncodeToString(hex)

	return hexStr
}

func GetHeaderACCESS(path string, body string) map[string]string {
	timeUnix := time.Now().Format(time.RFC3339)

	fmt.Println("timeUnix", timeUnix)

	SignStr := mStr.Join(
		timeUnix, // 开发者本地的时间戳，发起请求生成的
		path,     //  开发着请求哪个接口，他就得在这里写什么地址
		body,     // 他发出请求时的参数字符串
	)

	Sign := HmacSha256(SignStr, "e9ef03e8-d611-431b-8227-b8f15fa07af0")

	fmt.Println("Sign", Sign)

	OpenApiHeader := map[string]string{
		"OT-ACCESS-KEY":       "mJsp2X90ltkBNENFh799resyud3UqhovjY5iUgpKWLBMRSNMohWjrvt9kWQanAb5",
		"OT-ACCESS-SIGN":      Sign,
		"OT-ACCESS-TIMESTAMP": timeUnix, // 2006-01-02T15:04:05Z07:00
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

	path := "/openapi/position/init"
	res, err := fetch.Post(fetch.Opt{
		Origin: BaseUrl,
		Path:   "/openapi/position/init",
		Header: GetHeaderACCESS(path, mJson.ToStr(data)),
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
