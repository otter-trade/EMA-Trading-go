package trade

import (
	"EMA-Trading-go/global"
	"EMA-Trading-go/mJson"
	"EMA-Trading-go/okx"

	"EMA-Trading-go/mTalib"
)

// 指标计算
func (_this *TradeObj) CountIndex() {
	OffsetCandle := _this.NowCandle[:len(_this.NowCandle)-_this.OffsetNum]
	_this.OffsetCandle = OffsetCandle
	global.Log.Println("OffsetCandle 塞入完毕", _this.OffsetCandle[len(_this.OffsetCandle)-1].TimeStr)
}

type TradeCandle struct {
	okx.TypeKd
	EMA100 string
}

func (_this *TradeObj) FormatCandle() {
	nowTradeCandleList := []TradeCandle{}

	for _, v := range _this.NowCandle {
		EMA100 := mTalib.ClistNew(mTalib.ClistOpt{
			KDList: _this.NowCandle,
			Period: 100,
		}).EMA().ToStr()

		newDataObj := TradeCandle{
			TypeKd: v,
			EMA100: EMA100,
		}

		nowTradeCandleList = append(nowTradeCandleList, newDataObj)

		global.NowCnaLog.Println("NowKdata:", mJson.Format(newDataObj))
	}

	offsetTradeCandleList := []TradeCandle{}
	for _, v := range _this.OffsetCandle {
		EMA100 := mTalib.ClistNew(mTalib.ClistOpt{
			KDList: _this.OffsetCandle,
			Period: 100,
		}).EMA().ToStr()

		newDataObj := TradeCandle{
			TypeKd: v,
			EMA100: EMA100,
		}
		offsetTradeCandleList = append(offsetTradeCandleList, newDataObj)
		global.Log.Println("offsetKdata:", mJson.Format(newDataObj))
	}

	nowEma := nowTradeCandleList[len(nowTradeCandleList)-1].EMA100
	offsetEma := offsetTradeCandleList[len(offsetTradeCandleList)-1].EMA100

	Dir := 0
	if nowEma > offsetEma {
		Dir = 1
	}
	if nowEma < offsetEma {
		// 上涨趋势
		Dir = -1
	}

	global.Log.Println("当前信号指标为:", Dir)
}
