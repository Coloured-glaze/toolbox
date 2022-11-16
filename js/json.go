package js

import (
	"encoding/json"
	"fmt"
	"os"
)

// 解析json配置
func Analyze(a *any, config string) error {
	b, err := os.ReadFile(config)
	if err != nil {
		return fmt.Errorf("radFile: %v", err)
	}
	err = json.Unmarshal(b, &a)
	if err != nil {
		return fmt.Errorf("unmarshal: %v", err)
	}
	return nil
}

// 保存json配置. 类型, 配置名, 权限: 例如 0600
func Save(a any, config string, perm os.FileMode) error {
	b, err := json.Marshal(a)
	if err != nil {
		return fmt.Errorf("marshal: %v", err)
	}
	err = os.WriteFile(config, b, perm)
	if err != nil {
		return fmt.Errorf("writeFile: %v", err)
	}
	return nil
}
