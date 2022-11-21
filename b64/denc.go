package b64

import (
	"encoding/base64"

	"github.com/Coloured-glaze/toolbox/str"
)

// base64 编码
func B64es(s string) string {
	return B64e(str.StrtoByte(s))
}

// base64 编码
func B64e(src []byte) string {
	encodeString := base64.StdEncoding.EncodeToString(src)
	return encodeString
}

// base64 解码
func B64d(s string) (string, error) {

	decodeBytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return str.BytetoStr(decodeBytes), nil
}

// base64 URL 编码
func B64urles(s string) string {
	return B64urle(str.StrtoByte(s))
}

// base64 URL 编码
func B64urle(src []byte) string {
	uEnc := base64.URLEncoding.EncodeToString(src)
	return uEnc
}

// base64 URL 解码
func B64urld(s string) (string, error) {
	uDec, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return str.BytetoStr(uDec), nil
}
