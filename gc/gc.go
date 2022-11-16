package gc

import (
	"runtime/debug"
)

// 设定垃圾收集目标百分比, -1 为关闭
func SetGC(i int) {
	debug.SetGCPercent(i)
}
