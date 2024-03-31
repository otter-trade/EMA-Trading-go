package okx

import (
	"encoding/json"
	"fmt"
	"strconv"

	"EMA-Trading-go/fetch"
	"EMA-Trading-go/mTime"

	"EMA-Trading-go/mJson"
)

type TypeKd struct {
	InstID   string `bson:"InstID"`   // 持仓币种
	TimeUnix int64  `bson:"TimeUnix"` // 毫秒时间戳
	TimeStr  string `bson:"TimeStr"`  // 时间的字符串形式
	O        string `bson:"O"`        // 开盘
	H        string `bson:"H"`        // 最高
	L        string `bson:"L"`        // 最低
	C        string `bson:"C"`        // 收盘价格
	CBas     string `bson:"CBas"`     // 实体中心价 (收盘+最高+最低) / 3
	Vol      string `bson:"Vol"`      // 成交量( BTC 的数量 )
	Dir      int    `bson:"Dir"`      // 方向 (收盘-开盘) ，1：涨 & -1：跌 & 0：横盘
	HLPer    string `bson:"HLPer"`    // 振幅 (最高-最低)/最低 * 100%
	U_shade  string `bson:"U_shade"`  // 上影线
	D_shade  string `bson:"D_shade"`  // 下影线
	RosePer  string `bson:"RosePer"`  // 涨幅 当前收盘价 - 上一位收盘价 * 100%
	C_dir    int    `bson:"C_dir"`    // 实体中心价 (当前中心点-前中心点) 1：涨 & -1：跌 & 0：横盘
}

type GetOKXCandleOpt struct {
	InstId string
	Bar    string //
	After  string // 毫秒 1597026383085
}

// 参照文档 https://www.okx.com/docs-v5/zh/#public-data-rest-api-get-index-candlesticks-history

type GetCandleOpt struct {
	InstID string `bson:"InstID"`
	After  int64  `bson:"After"` // 此时间之前的内容
	Page   int    `bson:"Page"`  // 往前第几页
	Bar    string `bson:"Bar"`   // 1m/3m/5m/15m/30m/1h/2h/4h  默认 1 小时
}

func GetOKXCandle(opt GetCandleOpt) (resData []TypeKd, resErr error) {
	resData = []TypeKd{}
	resErr = nil

	if len(opt.InstID) < 3 {
		resErr = fmt.Errorf("instId 不能为空")
		return
	}

	BarObj := GetBarOpt(opt.Bar) // 获取时间间隔
	if BarObj.Interval < mTime.UnixTimeInt64.Minute {
		resErr = fmt.Errorf("Bar的间隔太小")
		return
	}

	Size := 100
	// 时间设置
	now := mTime.GetUnixInt64()
	after := mTime.GetUnixInt64()
	// 时间必须大于6年前，否则重置为当前
	if opt.After > now-mTime.UnixTimeInt64.Day*2190 {
		after = opt.After
	}
	// 处理分页
	if opt.Page > 0 {
		pastTime := int64(opt.Page) * BarObj.Interval * int64(Size) // 一页数据 =  100 * 时间间隔
		after = after - pastTime                                    // 减去过去的时间节点
	}

	// 判断应该采取哪个接口获取数据  after 距离 now 有多少条数据?
	path := "/api/v5/market/candles"
	fromNowItem := (now - after) / BarObj.Interval
	if fromNowItem > 800 { // 大于 800 条就从历史接口拿数据
		path = "/api/v5/market/history-index-candles"
	}

	res, err := fetch.Get(fetch.Opt{
		Origin: "https://www.okx.com",
		Path:   path,
		Data: map[string]any{
			"instId": opt.InstID,
			"bar":    BarObj.Okx,
			"after":  strconv.FormatInt(after, 10),
			"limit":  Size,
		},
	})
	if err != nil {
		resErr = err
		return
	}

	var result TypeReq
	json.Unmarshal(res, &result)
	if result.Code != "0" {
		resErr = fmt.Errorf("获取数据失败:%+v", result)
		return
	}

	jsonByte := mJson.ToJson(result.Data)

	var list []OkxCandleDataType
	err = json.Unmarshal(jsonByte, &list)
	if err != nil {
		resErr = err
		return
	}
	if len(list) != Size {
		resErr = fmt.Errorf("len(list)长度错误 %+v", len(list))
		return
	}

	KdataList := []TypeKd{} // 声明存储
	for i := len(list) - 1; i >= 0; i-- {
		item := list[i]
		kdata := TypeKd{
			InstID:   opt.InstID,
			TimeStr:  mTime.UnixFormat(item[0]),
			TimeUnix: mTime.ToUnixMsec(mTime.MsToTime(item[0], "0")),
			O:        item[1],
			H:        item[2],
			L:        item[3],
			C:        item[4],
			Vol:      item[5],
		}
		new_Kdata := NewKD(kdata, KdataList)
		KdataList = append(KdataList, new_Kdata)
	}

	if len(KdataList) != Size {
		resErr = fmt.Errorf("len(KdataList)长度错误 %+v", len(KdataList))
		return
	}

	resData = KdataList

	return
}
