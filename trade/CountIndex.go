package trade

import (
	"EMA-Trading-go/global"
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
	TradeCandleList := []TradeCandle{}

	for _, v := range _this.NowCandle {
		EMA100 := mTalib.ClistNew(mTalib.ClistOpt{
			KDList: _this.NowCandle,
			Period: 100,
		}).EMA().ToStr()

		newDataObj := TradeCandle{
			TypeKd: v,
			EMA100: EMA100,
		}

		TradeCandleList = append(TradeCandleList, newDataObj)
	}
}
