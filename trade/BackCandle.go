package trade

import (
	"EMA-Trading-go/global"
	"EMA-Trading-go/mStr"
	"EMA-Trading-go/mTime"
	"EMA-Trading-go/okx"
)

func (_this *TradeObj) FillBaseCandle() {
	limit := 100
	page := _this.CandleMaxLen / limit // 每次 100 条计算最大页数
	nowTime := mTime.GetUnixInt64()
	baseCandle := []okx.TypeKd{}
	for i := page - 1; i >= 0; i-- {
		// 每次递减的时间戳 毫秒
		diffTime := mTime.UnixTimeInt64.Hour * int64(limit) * int64(i)
		after := nowTime - diffTime
		resData, err := okx.GetOKXCandle(okx.GetCandleOpt{
			InstID: "BTC-USDT",
			Bar:    "1H",
			After:  after,
		})
		if err != nil {
			global.Log.Println("出现错误: ", mStr.ToStr(err))
		}
		baseCandle = append(baseCandle, resData...)
	}

	// 检查基础数据
	// for key := range baseCandle {
	// 	preIndex := key - 1
	// 	if preIndex < 0 {
	// 		preIndex = 0
	// 	}
	// 	preItem := baseCandle[preIndex]
	// 	nowItem := baseCandle[key]
	// 	if key > 0 {
	// 		fmt.Println(key, nowItem.TimeUnix-preItem.TimeUnix, nowItem.TimeStr)
	// 		if nowItem.TimeUnix-preItem.TimeUnix != mTime.UnixTimeInt64.Hour {
	// 			global.Log.Println("数据检查出错, ", key, mJson.Format(nowItem))
	// 			break
	// 		}
	// 	}
	// }

	_this.NowCandle = baseCandle
}
