package f

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// f.FOR(8, 5)
// for test
func FOR(g, t int) {
	if g > 256 {
		g = 256
	} else if g <= 0 {
		g = 1
	}
	if t > 60 {
		t = 60
	} else if t <= 0 {
		t = 1
	}
	j := int64(0)
	var wg sync.WaitGroup
	runtime.GOMAXPROCS(0)
	os := runtime.GOOS
	arch := runtime.GOARCH
	cpu := runtime.NumCPU()
	fmt.Printf("os: %v\tarch: %v\tCPU: %v\n", os, arch, cpu)
	start := time.Now()
	wg.Add(g)
	for i := 0; i < g; i++ {
		ch := make(chan bool)
		go func() {
			time.Sleep(time.Second * time.Duration(t))
			ch <- true
		}()
		go func() {
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
		}()
	}
	wg.Wait()
	ts := time.Since(start).Seconds()
	fmt.Printf("%.3f 秒内 %v 个协程加了 %v 次\n", ts, g, j)
}
