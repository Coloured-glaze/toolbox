package gc

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/Coloured-glaze/toolbox/cmd"
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

// 内存占用
func GetRAM() string {
	if runtime.GOOS == "windows" {
		return "Error! Don't Windows!"
	}
	R := cmd.Cmd("free -m | awk '/Mem/ {print $2\" \"$3}'")
	if len(R) == 0 {
		R = cmd.Cmd("free -m | awk '/内存/ {print $2\" \"$3}'")
	}
	if len(R) > 2 {
		sp := strings.Split(R, " ")
		Per, _ := strconv.ParseFloat(sp[0], 64)
		Per2, _ := strconv.ParseFloat(sp[1], 64)
		Per3 := Per2 / Per * 100
		return fmt.Sprintf(sp[1]+"M/"+sp[0]+"M") + fmt.Sprintf("(%.2f%%)", Per3)
	}
	return "null"
}
