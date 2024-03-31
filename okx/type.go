package okx

type TypeReq struct {
	Code string `bson:"Code"`
	Data any    `bson:"Data"`
	Msg  string `bson:"Msg"`
}
type OkxCandleDataType [9]string // Okx 原始数据
