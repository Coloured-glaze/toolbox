package filetool

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// 文件路径
type Files struct {
	Name         string // 文件完整路径
	NameOnly     string // 文件完整路径(不带后缀)
	FileName     string // 完整文件名(不带路径)
	NameNosuffix string // 文件名(不带后缀和路径)
	Suffix       string // 文件的后缀(文件类型)
	Path         string // 路径(不带文件名)
}

// 获取 dir 包含的指定 name 的文件的绝对或相对路径, exclude 为排除列表,可为 nil
func Readir(dir string, name, exclude []string) ([]Files, int, error) {
	if !IsExist(dir) {
		p, _ := os.Getwd()
		return nil, 0, fmt.Errorf("当前的路径 %v 未找到 >%v< 文件夹", p, dir)
	}
	files := make([]Files, 0, 128)
	num := 0 // 文件数量
	err := filepath.Walk(dir,
		func(filename string, fi os.FileInfo, _ error) error {
			if fi.IsDir() { // 忽略目录
				return nil
			}
			if exclude != nil {
				if Excluder(exclude, filename) {
					return nil
				}
			}
			p, _ := path.Split(filename) // 返回路径名, 和文件名
			f := Files{
				Name:     filename,  // 获取文件完整路径
				Path:     p,         // 路径
				FileName: fi.Name(), // 获取文件名
				//	FileName: path.Base(filename), // 获取文件名
			}
			f.Suffix = path.Ext(f.Name)                               // 获取文件的后缀(文件类型)
			f.NameOnly = strings.TrimSuffix(f.Name, f.Suffix)         // 获取文件名称(不带后缀)
			f.NameNosuffix = strings.TrimSuffix(f.FileName, f.Suffix) // 获取文件名(不带后缀和路径)
			if f.Path == "" {
				f.Path = Trim(f.Name)
			}
			f.addFile(&files, &f, &name, &num)
			return nil
		})
	if err != nil {
		return nil, num, fmt.Errorf("路径读取错误%v", err)
	}
	return files, num, nil
}

func Trim(s string) string {
	//	s = strings.ReplaceAll(s, "\\", "/")
	i := strings.LastIndex(s, "\\")
	if i < 0 {
		return ""
	}
	return s[:i]
}

// 排除文件
func Excluder(list []string, s string) bool {
	for i := 0; i < len(list); i++ {
		if list[i] != "" {
			if exist := strings.Contains(s, list[i]); exist {
				return true
			}
		}
	}
	return false
}

// 文件名添加进切片
func (*Files) addFile(files *[]Files, f *Files, suffix *[]string, num *int) {
	for i := range *suffix {
		if f.Suffix == (*suffix)[i] || strings.Contains(f.FileName, (*suffix)[i]) {
			(*files) = append((*files), *f)
			*num++
		}
	}
}

/*
package main
import (
	"fmt"
	"github.com/Coloured-glaze/toolbox/filetool"
)
func main() {
	f, _, err := filetool.Readir(`E:\1\demo2\test_ft`, []string{"."}, nil)
	if err != nil {
		panic(err)
	}
	for i := range f {
		fmt.Printf("%#v\n\n", f[i])
	}
}
*/
// filetool.Files{Name:"E:\\1\\demo2\\test_ft\\read.go", NameOnly:"E:\\1\\demo2\\test_ft\\read", FileName:"read.go", NameNosuffix:"read", Suffix:".go", Path:"E:\\1\\demo2\\test_ft"}
