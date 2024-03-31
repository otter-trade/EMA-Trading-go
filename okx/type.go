package okx

type TypeReq struct {
	Code string `bson:"Code"`
	Data any    `bson:"Data"`
	Msg  string `bson:"Msg"`
}
type (
	OkxCandleDataType [9]string // Okx 原始数据
	TypeKd            struct {
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
)
