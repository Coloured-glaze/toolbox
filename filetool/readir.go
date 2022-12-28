package filetool

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Files struct {
	Name     string // 文件完整路径
	NameOnly string // 文件名称(不带后缀)
	Suffix   string // 文件的后缀(文件类型)
}

// s, _, err := ft.Readir("E:\\1\\z2\\ys0", ".psd")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	for i, _ := range s {
// 		fmt.Printf("%#v\n", s[i])
// 	}

// 获取指定文件夹指定后缀文件的 绝对/相对 路径
func Readir(dir, suffix string) ([]Files, int, error) {
	files := make([]Files, 0, 32)
	num := 0 // 文件数量
	err := filepath.Walk(dir, func(filename string, fi os.FileInfo, err error) error {
		if fi.IsDir() { // 忽略目录
			return nil
		}
		f := Files{
			Name: path.Base(filename), // 获取文件完整路径
		}
		f.Suffix = path.Ext(f.Name)                       // 获取文件的后缀(文件类型)
		f.NameOnly = strings.TrimSuffix(f.Name, f.Suffix) // 获取文件名称(不带后缀)
		if f.Suffix == suffix {
			num++                    // 判断类型 数量+1
			files = append(files, f) // 文件名添加进切片
		}
		return nil
	})
	if err != nil {
		return nil, num, fmt.Errorf("路径读取错误%v", err)
	}
	return files, num, nil
}
