package global

import (
	"fmt"
	"log"

	"EMA-Trading-go/global/config"
	"EMA-Trading-go/mLog"
	"EMA-Trading-go/mTime"
)

var (
	Log *log.Logger // 系统日志

	NowCnaLog *log.Logger // 交易模块
)

func LogInit() {
	// 创建一个log
	// 创建一个log
	Log = mLog.NewLog(mLog.NewLogParam{
		Path: config.Dir.Log,
		Name: "Sys",
	})

	NowCnaLog = mLog.NewLog(mLog.NewLogParam{
		Path: config.Dir.Log,
		Name: "nowCnaLog",
	})

	// 设定清除log
	mLog.Clear(mLog.ClearParam{
		Path:      config.Dir.Log,
		ClearTime: mTime.UnixTimeInt64.Day * 10,
	})
}

func LogErr(sum ...any) {
	str := fmt.Sprintf("系统错误: %+v", sum)
	Log.Println(str)
}
