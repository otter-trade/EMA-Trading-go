package mTalib

import (
	"strconv"

	"EMA-Trading-go/mCount"
	"EMA-Trading-go/okx"
)

type ClistOpt struct {
	CList  []string     // 数据
	KDList []okx.TypeKd // 数据
	Period int          // 周期
}

type ClistObj struct {
	CList  []string // 数据
	FList  []float64
	Period int   // 周期
	CLen   int   // 数组长度
	DotNum int32 // 小数点位数
	Result float64
}

func ClistNew(opt ClistOpt) *ClistObj {
	obj := ClistObj{}
	obj.Period = opt.Period

	var CList []string
	var floatList []float64
	if len(opt.CList) > 0 {
		obj.DotNum = mCount.GetDecimal(opt.CList[0])
		obj.CLen = len(opt.CList)
		for _, val := range opt.CList {
			valDot := mCount.GetDecimal(val)
			if valDot > obj.DotNum { // 如果当前小数点位数大于现存小数点位数，则替换
				obj.DotNum = valDot
			}
			floatVal := mCount.ToFloat(val, obj.DotNum)
			floatList = append(floatList, floatVal)
			CList = append(CList, val)
		}
	} else if len(opt.KDList) > 0 {
		obj.DotNum = mCount.GetDecimal(opt.KDList[0].C)
		obj.CLen = len(opt.KDList)
		for _, val := range opt.KDList {
			valDot := mCount.GetDecimal(val.C)
			if valDot > obj.DotNum { // 如果当前小数点位数大于现存小数点位数，则替换
				obj.DotNum = valDot
			}
			floatVal := mCount.ToFloat(val.C, obj.DotNum)
			floatList = append(floatList, floatVal)
			CList = append(CList, val.C)
		}
	}

	obj.FList = floatList
	obj.CList = CList

	return &obj
}

func (_this *ClistObj) ToStr() string {
	rStr := strconv.FormatFloat(_this.Result, 'f', 10, 32)
	rStr = mCount.CentRound(rStr, _this.DotNum)

	return rStr
}
