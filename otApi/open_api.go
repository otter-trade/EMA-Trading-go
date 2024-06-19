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
	UpdatePosition()
	// RedPosition()
	// RedHistoryPosition()
}

// 初始化持仓
func InitPosition() {
	data := map[string]any{
		"MockName":     MockName, // 策略ID 每个策略唯一，当前用户的当前策略
		"InitialAsset": 10000,    // 初始余额 缺省值 10000
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

/*

将来会存在一个接口：罗列和整合统一  所有的数字货币交易所的 币种信息 ，他们会全部统一成 OtterTrade 标准，默认标准是OKX


*/

/*
我要回测一个策略
只要有价格K线 和产品定义  就可以完成任何 事物 的 量化策略
1小时计算一次 根据计算结果进行下单和平仓操作

# UpdatePosition 需要完成的事情

下单的时候需要的信息

InstList

交易对(InstName):  币安标准 BTC_USDT  欧意的标准 btc-usdt   OtterTrade的标准 BTC-USDT  格式:   xxxx-USDT
交易模式(tradeMode)： 永续合约 SWAP  现货 SPOT
交易类型(tradeType)： 数字货币，期货，股票 或其它 目前只有数字货币
杠杆倍数(Leverage)：  只有 永续合约 才存在 杠杆倍数  ， 现货这个值始终为1
下单方向(Side)： 开空-1   开多 1   平仓 0   (合约有三个方向)   现货 (只有两个方向   1 开多 0或者-1  平仓或者卖出)
开仓资金(xxx)： 开仓资金 (如果开仓资金>实际余额) 报错并返回说明无效参数原因，你没有那么多资金

实现方式：
如果说以更新仓位状态的方式去完成下单

*交易对(InstName):  币安标准 BTC_USDT  欧意的标准 btc-usdt   OtterTrade的标准 BTC-USDT  格式:   xxxx-USDT
*交易模式(tradeType)： 永续合约 SWAP  现货 SPOT
杠杆倍数(Leverage)： 只有 永续合约 才存在 杠杆倍数  ， 现货这个值始终为1  缺省值为1 （计算方式: 下单资金[或下单数量]*杠杆倍数）
下单方向(Side)： 买多，买空， null
开仓资金(xxx)： 开仓资金 (如果开仓资金>实际余额) 报错并返回说明无效参数原因，你没有那么多资金

如果是回测的话
只需要加上下单的时间即可
Timestamp(毫秒)：比如  2024-06-08 14:00:00 的unix 时间戳 ; 如果为 0 则为当前时间
这个时候应该采用 当前交易对 这个时间点的价格进行计算
*/

/*
AItrade 的实现方式
交易品 BTC-USDT

下单时价格 a
当前价格 b

盈利公式
盈利 =(现价-下单价格)/下单价格

涨幅计算公式：
涨幅=(现价-上一个交易日收盘价)/上一个交易日收盘价

有了比率之后
余额*(杠杆倍率*盈利比率) = 本次盈亏
余额- 本次盈亏 = 新的余额


*/

func UpdatePosition() {

	// 更新了一次持仓， BTC-USDT 合约 2x 做空
	// data := map[string]any{
	// 	"MockName":  MockName,   // 策略ID 每个策略唯一，当前用户的当前策略
	// 	"Timestamp": 1717826400, // 时间戳   2024-06-08 14:00:00
	// 	"NewPosition": PositionType{
	// 		{
	// 			"InstName":     "BTC-USDT", // 交易产品ID
	// 			"tradeType":   "SWAP",  //  交易模式
	// 			"Leverage":   5,          // 杠杆倍数
	// 			"Side":       1,          // 买卖方向 1 空 2 多  &  <应该改成  -1 卖空 1 买多 0 空仓>
	// 			"Proportion": 1,        // 持仓比例
	// 		},
	// 	},
	// }

	// data := map[string]any{
	// 	"MockName":  MockName,   // 策略ID 每个策略唯一，当前用户的当前策略
	// 	"Timestamp": 1717826400, // 时间戳   2024-06-08 14:00:00
	// 	"NewPosition": PositionType{
	// 		{
	// 			"InstName":     "BTC-USDT", // 交易产品ID
	// 			"tradeType":   "SWAP",  //  交易模式
	// 			"Leverage":   5,          // 杠杆倍数
	// 			"Side":       1,          // 买卖方向 1 空 2 多  &  <应该改成  -1 卖空 1 买多 0 空仓>
	// 			"Proportion": 0,        // 持仓比例
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
/*
我们要读取和定义一个账户的持仓
对象： MockName
时间： Timestamp 我在这个时间点的仓位状态  ，不传或者为0则是当前
这个时候应该采用 当前交易对 这个时间点的价格进行计算


读取持仓的返回数据
"StrategyID": "e395ffe2-3a99-4072-abfc-0da940f32f26",
"RunType": 1,
"InstanceID": "backtest_mo7_test_first",
"InitialAsset": 10000, // 初始余额
"FeeRate": 0.01,
"Cash": 2000,  // 当前余额

盈利(xx):  -123 为亏损  123 为盈利,
盈利比率(xx)： 距离上一单的盈利比率,
"TimestampNs": 1718108999434701800,(这个读取该持仓时的时间 均为unix毫秒时间戳)
"InitTimestampNs": 1718108206624818400, (InitPosition时间)

CurrPosition: [
  {
		交易对(InstName):  币安标准 BTC_USDT  欧意的标准 btc-usdt   OtterTrade的标准 BTC-USDT  格式:   xxxx-USDT
		交易模式(tradeType)： 永续合约 SWAP  现货 SPOT
		杠杆倍数(Leverage)： 只有 永续合约 才存在 杠杆倍数  ， 现货这个值始终为1  缺省值为1
    开仓资金(xxx)： 开仓资金 (如果开仓资金>实际余额) 报错并返回说明无效参数原因，你没有那么多资金

		下单方向(Side)： 买多，买空， null
		instPrice: 该交易品类读取时的价格 (当前价格)
		下单的时间
		成本价(下单时的价格)
		当前交易对未实现收益率: 500%
		当前未实现收益：(上一次的余额*当前收益率) 10000*500%
		余额(balance)：（上一次的余额 - 当前未实现收益) 读取当前持仓时都是未实现余额  = 50000 只有平仓之后才会有，也就是历史才会有
	},
]

*/
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
