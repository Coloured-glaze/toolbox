package cmd

import (
	"fmt"
	"os/exec"
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
