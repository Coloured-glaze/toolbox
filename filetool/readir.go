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

// 	s, _, err := ft.Readir("./", ".webp")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	for i, _ := range s {
// 		fmt.Printf("%#v\n\n", s[i])
// 	}

// filetool.Files{Name:"four2/tmp/i_n11422.webp", NameOnly:"four2/tmp/i_n11422", Path:"four2/tmp/", FileName:"i_n11422.webp", Suffix:".webp"}

// 获取指定文件夹指定后缀文件的 绝对/相对 路径
func Readir(dir, suffix string) ([]Files, int, error) {
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
