package str

import (
	"strconv"
)

// unicode 转 string
// 输入 > `"\u7a0b\u5e8f"` <
func Utostr(s string) (string, error) {
	su , err := strconv.Unquote(s)
	if err !=nil {
		return "", err
	}
	return su, nil
}
// string 转 unicode
func StrtoU(s string) string {
	return strconv.QuoteToASCII(s)
}