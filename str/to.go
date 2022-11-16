package str

import (
	"reflect"
	"unsafe"
)

// []byte 转 string
func BytetoStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b)) //转成 string
}

// string 转 []byte
func StrtoByte(str string) []byte {
	p := (*[2]uintptr)(unsafe.Pointer(&str)) //转成 uinptr
	p2 := [3]uintptr{p[0], p[1], p[1]}       //值 长度 容量
	return *(*[]byte)(unsafe.Pointer(&p2))   //转成 []byte 切片
}

// string 转 []byte
func Strtob(s string) (b []byte) {
	rb := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	rs := (*reflect.StringHeader)(unsafe.Pointer(&s))
	rb.Data = rs.Data
	rb.Len = rs.Len
	rb.Cap = rs.Len
	return
}
