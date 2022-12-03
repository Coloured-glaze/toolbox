// 百度翻译
package bdfy

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Coloured-glaze/toolbox/check"
)

type Trans struct {
	From      string `json:"from"`
	To        string `json:"to"`
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	Tresult   []struct {
		Src string `json:"src"`
		Dst string `json:"dst"`
	} `json:"trans_result"`
}

type Account struct {
	Uid       int    `db:"uid"`
	AppId     string `db:"appid"`
	SecretKey string `db:"secretKey"`
}

var (
	// 原文
	//	Q string = "apple"
	// 源语言
	From string = "auto"
	// 目标语言
	To string = "zh"

	// 改成您的APP ID
	AppId string = ""
	// 改成您的密钥
	SecretKey string = ""

	Salt string = strconv.Itoa(rand.Intn(65536))

	Bdurl string = "http://api.fanyi.baidu.com/api/trans/vip/translate?"
)

func Tr(Q string) (*Trans, error) {
	t := new(Trans)
	sign := check.MD5(AppId + Q + Salt + SecretKey)
	tmp := strings.Builder{}
	tmp.Grow(len(Q) + 192)
	tmp.WriteString("q=" + url.QueryEscape(Q))
	tmp.WriteString("&from=" + From)
	tmp.WriteString("&to=" + To)
	tmp.WriteString("&appid=" + AppId)
	tmp.WriteString("&salt=" + Salt)
	tmp.WriteString("&sign=" + sign)
	r, err := http.Get(Bdurl + tmp.String())
	if err != nil {
		return t, err
	}
	defer r.Body.Close()
	resp, err := io.ReadAll(r.Body)
	if err != nil {
		return t, err
	}
	err = json.Unmarshal(resp, &t)
	if err != nil {
		return t, err
	}
	return t, nil
}
