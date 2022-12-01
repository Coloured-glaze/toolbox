package cmd

import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/Coloured-glaze/toolbox/str"
)

// 调用 命令
func Cmd(ex string, args ...string) string {
	cmd := exec.Command(ex, args...)
	obyte, err := cmd.CombinedOutput() // 获取输出
	if err != nil {
		return fmt.Sprintf("%v %v", str.BytetoStr(obyte), err)
	}
	if len(obyte) == 0 {
		return ""
	}
	return strings.Trim(str.BytetoStr(obyte), "\n") //去除换行符
}

/*
// 执行命令
func ChCmd(ex string, args ...string, time int) {
	dc := gg.NewContext(0, 0)
	tmp := strings.Builder{}
	cmd := exec.Command(ex, args...)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(time))


}
*/
// 内存占用
func GetRAM() string {
	if runtime.GOOS == "windows" {
		return "Error! Don't Windows!"
	}
	R := Cmd("free -m | awk '/Mem/ {print $2\" \"$3}'")
	if len(R) == 0 {
		R = Cmd("free -m | awk '/内存/ {print $2\" \"$3}'")
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
