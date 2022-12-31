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

// 加载 zip 文件 到 map
func Parsezip(zipFile string, filemap map[string][]*zip.File) error {
	zRead, err := zip.OpenReader(zipFile) // will not close
	if err != nil {
		return err
	}
	dir := ""
	for i := range zRead.File { // 切片
		if zRead.File[i].FileInfo().IsDir() {
			//	fmt.Println("dir=", f.Name)
			dir = zRead.File[i].Name
			//	paths[f.Name] = append(paths[f.Name], f.Name)
			continue
		}
		//	fmt.Printf("file=%v\n", f.Name)
		//	fmt.Printf("%#v\n\n", f.FileHeader)
		//	fmt.Printf("%#v\n\n", f.FileHeader.FileInfo())
		filemap[dir] = append(filemap[dir], zRead.File[i]) // f = *zip.File
	}
	return nil
}

// 打开map中的ttf文件
func LoadFontFaces(dc *gg.Context, points float64, path string, name string, filemap map[string][]*zip.File) error {
	readerData, err, ok := Findfile(path, name, filemap)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf(path + " 没有那个文件[" + name + "]")
	}
	defer readerData.Close()
	buf := &bytes.Buffer{}
	buf.ReadFrom(readerData)

	// retrieve a byte slice from bytes.Buffer
	f, err := truetype.Parse(buf.Bytes())
	if err != nil {
		return err
	}
	face := truetype.NewFace(f, &truetype.Options{
		Size: points,
		// Hinting: font.HintingFull,
	})
	dc.SetFontFace(face)
	// dc.fontFace = face
	// dc.fontHeight = points * 72 / 96
	return nil
}

// 查找map中的文件 并打开
func Findfile(path string, name string, filemap map[string][]*zip.File) (io.ReadCloser, error, bool) {
	var readerData io.ReadCloser
	var err error
	ok := false
	for i := range filemap[path] {
		if strings.Contains(filemap[path][i].Name, name) {
			readerData, err = filemap[path][i].Open()
			if err != nil {
				return nil, err, ok
			}
			ok = true
			break
		}
	}
	return readerData, nil, ok
}
