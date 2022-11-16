package f

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
 f.FOR(8, 5)
*/

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
	times := time.Duration(t) // 时间

	wg := sync.WaitGroup{}
	wg.Add(g)
	os := runtime.GOOS
	arch := runtime.GOARCH
	cpu := runtime.NumCPU()
	// runtime.GOMAXPROCS(0)

	fmt.Printf("OS:%v\tARCH:%v\tCPU:%v\n", os, arch, cpu)

	start := time.Now()
	for i := 0; i < g; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*times)
		go func(ctx context.Context, cancel context.CancelFunc) {
			var k = int64(0)
			for {
				select {
				case <-ctx.Done():
					cancel()
					goto label
				default:
					k++
				}
			}
		label:
			j += k
			wg.Done()
		}(ctx, cancel)
	}
	wg.Wait()
	fmt.Printf("%.3f 秒内 %v 个协程加了 %v 次\n",
		time.Since(start).Seconds(), g, j)
}
