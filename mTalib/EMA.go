package mTalib

import (
	"EMA-Trading-go/mTalib/talib"
)

func (_this *ClistObj) EMA() *ClistObj {
	if _this.CLen < _this.Period+1 {
		return _this
	}
	pArr := talib.Ema(_this.FList, _this.Period)
	_this.Result = pArr[_this.CLen-1]
	return _this
}
