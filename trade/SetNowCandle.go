package trade

import (
	"EMA-Trading-go/global"
	"EMA-Trading-go/mJson"
	"EMA-Trading-go/mStr"
	"EMA-Trading-go/mTime"
	"EMA-Trading-go/okx"
)

func (_this *TradeObj) SetNowCandle() {
	// 检查基础数据
	resData, err := okx.GetOKXCandle(okx.GetCandleOpt{
		InstID: "BTC-USDT",
		Bar:    "1H",
	})
	if err != nil {
		global.Log.Println("出现错误: ", mStr.ToStr(err))
	}

	for _, NowItem := range resData {
		Fund := false
		FundKey := 0

		for key, Item := range _this.NowCandle {
			if NowItem.TimeUnix == Item.TimeUnix { // 相等的直接替换
				Fund = true
				FundKey = key
				break
			}
		}

		if Fund {
			_this.NowCandle[FundKey] = NowItem
		} else {
			_this.NowCandle = append(_this.NowCandle, NowItem)
		}

	}

	if len(_this.NowCandle)-_this.CandleMaxLen > 0 {
		_this.NowCandle = _this.NowCandle[len(_this.NowCandle)-_this.CandleMaxLen:]
	}

	// 检查基础数据
	for key := range _this.NowCandle {
		preIndex := key - 1
		if preIndex < 0 {
			preIndex = 0
		}
		preItem := _this.NowCandle[preIndex]
		nowItem := _this.NowCandle[key]
		if key > 0 {
			if nowItem.TimeUnix-preItem.TimeUnix != mTime.UnixTimeInt64.Hour {
				global.Log.Println("数据检查出错, ", key, mJson.Format(nowItem))
				break
			}
		}
	}
	global.Log.Println("最新数据塞入完毕,当前共", len(_this.NowCandle), "条,最后的时间为", _this.NowCandle[_this.CandleMaxLen-1].TimeStr, _this.NowCandle[_this.CandleMaxLen-1].C)
}
