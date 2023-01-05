package gid

import (
	"bytes"
	"runtime"
	"strconv"

	s "github.com/Coloured-glaze/toolbox/str"
)

// 获取 goroutine id
func Get() int64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseInt(s.BytetoStr(b), 10, 64)
	return n
}

func Gets() string {
	b := make([]byte, 32)
	b = bytes.TrimPrefix(b[:runtime.Stack(b, false)], s.StrtoByte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	return s.BytetoStr(b)
}
