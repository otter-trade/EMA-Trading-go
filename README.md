# EMA-Trading-go

用 go 写一个简单的 EMA 交易策略，用于测试

## 策略思路:

以 EMA 100 进行平法交易

设置 1h 的 K 线，然后用 EMA 进行平移，平移周期为 5

伪代码:

```go

var nowCandle = []array  // 当前K线

var offsetNum = 5   // 平移周期
var offsetCandle = nowCandle[:len(nowCandle)-offsetNum]  // 偏移后的K线
var nowEma = index.Eam(nowCandle,100); // 当前K线EMA
var offsetEma  = index.Eam(offsetCandle,100); // 偏移后K线Ema

// 趋势判断：
var Dir = 0
if (nowEma < offsetEma){
  // 下跌趋势
	Dir = -1
}
if (nowEma > offsetEma){
  // 上涨趋势
	Dir = 1
}

var preDir = "上一条K线的Dir"

if (Dir != preDir) {
	if (Dir <0) {
		// 发送开空信号
	}
	if (Dir >0) {
		// 发送开多信号
	}
}

```

```bash

curl --location --request POST 'http://localhost:8802/open-api/market/candles' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--header 'Content-Type: application/json' \
--header 'Accept: */*' \
--header 'Host: test-api.ottertrade.com' \
--header 'Connection: keep-alive' \
--data-raw '{
  "instId":"BTC-USDT",
  "bar":"1min",
  "before":1712404848,
  "limit":20
}'



```
