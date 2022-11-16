package str

import (
	"math/rand"
	"strings"
	"time"
)

/*
goos: windows
goarch: amd64
cpu: Intel(R) Core(TM) i3-10100F CPU @ 3.60GHz

BenchmarkR-8   	  925561	      1392 ns/op	     791 B/op	       5 allocs/op

BenchmarkR2-8    	   11523	    103789 ns/op	     120 B/op	       3 allocs/op
*/

// 乱序输出字符串
func RandText(str string) string {
	m := make(map[int]rune)
	for i, v := range str {
		m[i] = v
	}
	tmp := strings.Builder{}
	tmp.Grow(len(str))
	for _, v := range m {
		tmp.WriteRune(v)
	}
	return tmp.String()
}

// 乱序输出字符串
func RandText2(str string, number int) string {
	if number > 25 {
		number = 25
	}
	rand.Seed(time.Now().UnixNano())
	textmp := make([]string, 0, number)

	for i := 0; i < number; i++ {
		tmp := strings.Builder{}
		tmp.Grow(len(str))
		rs := rands([]rune(str))
		for i := 0; i < len(rs); i++ {
			tmp.WriteRune(rs[i])
		}
		textmp = append(textmp, tmp.String())
	}
	return textmp[rand.Intn(number)]
}

func rands(runstr []rune) []rune {
	tmp := make([]rune, len(runstr))
	for i := range runstr {
		n := rand.Intn(len(runstr))
		if tmp[n] == 0 {
			tmp[n] = runstr[i]
		} else if tmp[n] != 0 { // 不为空
			rand.Seed(time.Now().UnixNano())
			switch rand.Intn(10) {
			case 0:
				for j := 0; j < len(runstr); j++ {
					if tmp[j] == 0 {
						tmp[j] = runstr[i]
						break
					}
				}
			default:
				for j := len(runstr) - 1; j > 0; j-- {
					if tmp[j] == 0 {
						tmp[j] = runstr[i]
						break
					}
				}
			}
		}
	}
	return tmp
}
