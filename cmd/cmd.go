package cmd

import (
	"os/exec"
	"strings"

	"github.com/Coloured-glaze/toolbox/str"
)

// 执行命令
func Cmd(ex string, args ...string) (string, error) {
	cmd := exec.Command(ex, args...)
	output, err := cmd.CombinedOutput() // 获取输出
	if err != nil {
		return str.BytetoStr(output), err
	}
	if len(output) == 0 {
		return "", nil
	}
	return strings.Trim(str.BytetoStr(output), "\n"), nil //去除换行符
}
