package bdfy

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/Coloured-glaze/toolbox/str"
)

var (
	surl = "https://transmart.qq.com/api/imt"

	Form3 = "en"
	To3   = "zh"

	Lang3 = map[string]string{
		"ar": "阿拉伯语", "de": "德语",
		"ru": "俄语", "fr": "法语",
		"tl": "菲律宾语", "km": "高棉语",
		"kr": "韩语", "lo": "老挝语",
		"pt": "葡萄牙语", "ja": "日语",
		"es": "西班牙语", "it": "意大利语",
		"id": "印度尼西亚语", "en": "英语",
		"vi": "越南语", "zh": "中文",
	}
)

type Ten struct {
	Header Header `json:"header"`
	Source Source `json:"source"`
	Target Target `json:"target"`
}

type Header struct {
	Fn        string `json:"fn"`
	ClientKey string `json:"client_key"`
}

type Source struct {
	Lang     string   `json:"lang"`
	TextList []string `json:"text_list"`
}

type Target struct {
	Lang string `json:"lang"`
}

// ====================================

type Tresp struct {
	Header          Header2  `json:"header"`
	AutoTranslation []string `json:"auto_translation"`
	Msg             string   `json:"massage"`
}

type Header2 struct {
	Type      string  `json:"type"`
	RetCode   string  `json:"ret_code"`
	TimeCost  float64 `json:"time_cost"`
	RequestID string  `json:"request_id"`
}

func Tx(texts []string) (Tresp, error) {
	h := Ten{
		Header: Header{
			Fn: "auto_translation",
			//	ClientKey: "dHJphbnNtYXJ0X2NyeF9Nb3ppbGxhLzUuMCA2oV2lpVCAxMC4wOyBXaW42NDsgeDY0KSBBcHBsZVdlYktpdC81MzcuMzYg",
			ClientKey: "browser-chrome-99.0.4951-Linux-ec27876d-" +
				strconv.FormatInt(rand.Int63n(160000000000), 10),
		},
		Source: Source{
			Lang:     Form3,
			TextList: texts,
		},
		Target: Target{
			Lang: To3,
		},
	}
	tr := Tresp{}
	b, err := json.Marshal(h)
	if err != nil {
		return tr, err
	}

	client := http.Client{}
	req, err := http.NewRequest("POST", surl, strings.NewReader(str.Btos(b)))
	if err != nil {
		return tr, err
	}
	req.Header = http.Header{
		"Host":         []string{"transmart.qq.com"},
		"accept":       []string{"application/json", "text/plain", "*/*"},
		"content-type": []string{"application/json"},
		"user-agent":   []string{"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.57 Safari/537.36"},
		"origin":       []string{"https://transmart.qq.com"},
		//	"sec-fetch-site": []string{"same-origin"},
		"sec-fetch-mode": []string{"cors"},
		"sec-fetch-dest": []string{"empty"},
		// "accept-encoding": []string{"gzip, deflate, br"},
		"accept-language": []string{"zh-CN,zh;q=0.9,en;q=0.8"},
		//		"content-length": []string{"332"},
	}
	resp, err := client.Do(req)
	if err != nil {
		return tr, err
	}
	defer resp.Body.Close()
	r, err := io.ReadAll(resp.Body)
	if err != nil {
		return tr, err
	}
	//	fmt.Printf("%v\n", string(r))
	err = json.Unmarshal(r, &tr)
	if err != nil {
		return tr, err
	}
	return tr, nil
	//	s := time.Since(now).Seconds()
}

/* sample:

func run() {
	text := "apple"
	reply, err := bdfy.Tx([]string{text})
	if err != nil {
		fmt.Println("Translate error:", err)
		return
	}
	if reply.Msg != "" {
		fmt.Println(reply.Msg)
		return
	}
	ra := reply.AutoTranslation
	for i := 0; i < len(ra); i++ {
		fmt.Println(ra[i])
	}
}
*/
