package filetool

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/Coloured-glaze/gg"
	"github.com/golang/freetype/truetype"
)

type Ftool struct {
	Ok   bool          // 是否存在
	Open bool          // 是否打开
	Err  error         // 错误
	Data io.ReadCloser //
	File *zip.File     //
}

// 加载 zip 文件到 map
func Parsezip(zipFile string, filemap map[string][]*zip.File) error {
	data, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	dir := ""
	for i := range data.File {
		if data.File[i].FileInfo().IsDir() {
			dir = data.File[i].Name
			filemap[dir] = make([]*zip.File, 0, 12)
			continue
		}
		filemap[dir] = append(filemap[dir], data.File[i])
	}
	return nil
}

// 打开map中的ttf文件
func LoadFontFaces(dc *gg.Context, points float64, path string, name string, filemap map[string][]*zip.File) error {
	ft := Findfile(path, name, true, filemap) // open
	if ft.Err != nil {
		return ft.Err
	}
	if !ft.Ok {
		return fmt.Errorf(path + " 没有那个文件[" + name + "]")
	}
	defer ft.Data.Close()
	buf := &bytes.Buffer{}
	buf.ReadFrom(ft.Data)

	// retrieve a byte slice from bytes.Buffer
	f, err := truetype.Parse(buf.Bytes())
	if err != nil {
		return err
	}
	face := truetype.NewFace(f, &truetype.Options{
		Size: points,
	})
	dc.SetFontFace(face)
	return nil
}

// 查找map中的文件
func Findfile(path string, name string, open bool, filemap map[string][]*zip.File) Ftool {
	var ft Ftool
	ft.Open = open
	for i := range filemap[path] {
		if strings.Contains(filemap[path][i].Name, name) {
			if open {
				ft.Data, ft.Err = filemap[path][i].Open()
				if ft.Err != nil {
					return ft
				}
			} else {
				ft.File = filemap[path][i]
			}
			ft.Ok = true
			break
		}
	}
	return ft
}
