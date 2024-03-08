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

// 下载文件
func DownloadTo(url, savePath string, chkcrt bool) (n int64, err error) {
	var resp *http.Response
	switch chkcrt {
	case true:
		resp, err = http.Get(url)
	case false:
		resp, err = nochkcrtcli.Get(url)
	}
	if err == nil {
		defer resp.Body.Close()
		total := int64(0)
		size := resp.ContentLength
		rsize := resp.Header.Get("Content-Length") // 响应内容的大小
		if rsize != "" {
			total, err = strconv.ParseInt(rsize, 10, 64) // 总计
		}
		if err == nil {
			switch {
			case total > 100*1024*1024 || size > 100*1024*1024: // 100MiB
				var f *os.File
				f, err = os.Create(savePath) // 持续写入
				if err == nil {
					defer f.Close()
					n, err = io.Copy(f, resp.Body) // 使用 io.Copy 函数将响应体复制到文件中
				}
			default:
				var data []byte
				data, err = io.ReadAll(resp.Body)
				n = int64(len(data))
				if err == nil {
					err = os.WriteFile(savePath, data, 0700) // 一次性写入文件
				}
			}
		}
	}
	return
}

// IsExist 文件/路径存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
