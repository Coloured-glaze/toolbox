package f

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"

	"github.com/Coloured-glaze/toolbox/check"
	"github.com/Coloured-glaze/toolbox/cmd"
)

// f.Forc(8, 5)
func Forc(g, t int) {
	if g > 256 {
		g = 256
	} else if g <= 0 {
		g = 1
	}
	if t > 120 {
		t = 120
	} else if t <= 0 {
		t = 1
	}
	j := int64(0)
	times := float64(t) // 时间
	var wg sync.WaitGroup
	runtime.GOMAXPROCS(0)
	os := runtime.GOOS
	arch := runtime.GOARCH
	cpu := runtime.NumCPU()
	fmt.Printf("OS: %v\tARCH: %v\tCPU: %v\n", os, arch, cpu)

	ss, _ := strconv.ParseFloat(strconv.Itoa(int(times)), 64)
	sleep := "sleep " + strconv.Itoa(int(ss)) + "s" // shell
	start, err := cmd.Cmd("date '+%s%N'")
	check.Checker(err)
	starts, err := strconv.ParseFloat(start, 64)
	check.Checker(err)

	wg.Add(g)
	for i := 0; i < g; i++ {
		ch := make(chan struct{})
		go func(ch chan struct{}) {
			cmd.Cmd(sleep) // shell
			ch <- struct{}{}
		}(ch)
		go func(ch chan struct{}) {
			var k = int64(0)
			for {
				select {
				case <-ch:
					goto label
				default:
					k++
				}
			}
		label:
			j += k
			wg.Done()
		}(ch)
	}
	wg.Wait()
	end, err := cmd.Cmd("date '+%s%N'")
	check.Checker(err)
	ends, err := strconv.ParseFloat(end, 64)
	check.Checker(err)
	fmt.Printf("%.3f 秒内 %v 个协程加了 %v 次\n", (ends-starts)/1e9, g, j)
}
