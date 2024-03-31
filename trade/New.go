package trade

import "EMA-Trading-go/okx"

type TradeObj struct {
	NowCandle    []okx.TypeKd // 当前K线
	OffsetNum    int          // 平移周期
	OffsetCandle []okx.TypeKd // 平移后的K线
	CandleMaxLen int          // K线最大长度
}

// 新建一个策略
func New() *TradeObj {
	obj := TradeObj{}
	// 初始化数据
	obj.OffsetNum = 5
	obj.NowCandle = []okx.TypeKd{}
	obj.OffsetCandle = []okx.TypeKd{}
	obj.CandleMaxLen = 400

	return &obj
}
