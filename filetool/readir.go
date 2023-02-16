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
	NameOnly string // 文件完整路径(不带后缀)
	Path     string // 路径(不带文件名)
	FileName string // 完整文件名(不带路径)
	Suffix   string // 文件的后缀(文件类型)
}

// 获取文件夹包含指定名称的文件的绝对或相对路径
func Readir(dir, name string) ([]Files, int, error) {
	if !IsExist(dir) {
		p, _ := os.Getwd()
		return nil, 0, fmt.Errorf("当前的路径 %v\n未找到 %v 文件夹", p, dir)
	}
	files := make([]Files, 0, 32)
	num := 0 // 文件数量
	err := filepath.Walk(dir, func(filename string, fi os.FileInfo, err error) error {
		if fi.IsDir() { // 忽略目录
			return nil
		}
		d, _ := path.Split(filename) // 返回路径名, 和文件名
		f := Files{
			Name:     filename,  // 获取文件完整路径
			Path:     d,         // 路径
			FileName: fi.Name(), // 获取文件名
			//	FileName: path.Base(filename), // 获取文件名
		}
		f.Suffix = path.Ext(f.Name)                       // 获取文件的后缀(文件类型)
		f.NameOnly = strings.TrimSuffix(f.Name, f.Suffix) // 获取文件名称(不带后缀)
		if f.Path == "" {
			f.Path = trim(f.Name)
		}
		if strings.Contains(f.FileName, name) {
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

func trim(s string) string {
	//	s = strings.ReplaceAll(s, "\\", "/")
	i := strings.LastIndex(s, "\\")
	if i < 0 {
		return ""
	}
	return s[:i]
}

// f, _, err := ft.Readir("./", ".webp")
// if err != nil {
// 	panic(err)
// }
// for i := range f {
// 	fmt.Printf("%#v\n\n", f[i])
// }

// filetool.Files{Name:"four2/tmp/i_n11422.webp", NameOnly:"four2/tmp/i_n11422", Path:"four2/tmp/", FileName:"i_n11422.webp", Suffix:".webp"}
