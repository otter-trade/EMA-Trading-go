package mTalib

import (
	"EMA-Trading-go/mTalib/talib"
)

func (_this *ClistObj) MA() *ClistObj {
	if _this.CLen < _this.Period+1 {
		return _this
	}
	pArr := talib.Sma(_this.FList, _this.Period)
	_this.Result = pArr[_this.CLen-1]
	return _this
}
