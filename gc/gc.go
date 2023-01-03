package gc

import (
	"runtime"
	"runtime/debug"
	ft "github.com/Coloured-glaze/toolbox/filetool"
)

// 设定垃圾收集目标百分比, -1 为关闭
func SetGC(i int) {
	debug.SetGCPercent(i)
}

// 查看程序占用内存
func GetMem() string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return ft.FormatSize(int64(m.Sys))
}