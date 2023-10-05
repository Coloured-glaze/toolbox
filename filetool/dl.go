package filetool

import (
	"crypto/tls"
	"io"
	"net/http"
	"os"
	"strconv"
)

var (
	tr = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	nochkcrtcli = &http.Client{Transport: tr}
)

// DownloadTo 下载到路径
func DownloadTo(url, file string, chkcrt bool) error {
	var resp *http.Response
	var err error
	if chkcrt {
		resp, err = http.Get(url)
	} else {
		resp, err = nochkcrtcli.Get(url) //
	}
	if err == nil {
		defer resp.Body.Close()
		total := int64(0)
		rsize := resp.Header.Get("Content-Length")
		size := resp.ContentLength // 响应内容的大小
		if rsize != "" {
			total, err = strconv.ParseInt(rsize, 10, 64) // 总计
		}
		if err == nil {
			if total > 100*1024*1024 || size > 100*1024*1024 { // 100MiB
				// fmt.Println("持续写入")
				var f *os.File
				f, err = os.Create(file)
				if err == nil {
					defer f.Close()
					_, err = io.Copy(f, resp.Body) // 使用 io.Copy 函数将响应体复制到文件中
				}
			} else {
				// fmt.Println("一次性写入", total, size)
				var data []byte
				data, err = io.ReadAll(resp.Body)
				if err == nil {
					err = os.WriteFile(file, data, 0700)
				}
			}
		}
	}
	return err
}

// IsExist 文件/路径存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
