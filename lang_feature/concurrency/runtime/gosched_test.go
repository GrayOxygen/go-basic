package runtime

import (
	"testing"
	"runtime"
	"fmt"
	"time"
	"sync"
)

func TestGoSched(t *testing.T) {
	wg := new(sync.WaitGroup)
	wg.Add(3)
	runtime.GOMAXPROCS(2)
	go func() {
		fmt.Println("runner a")
		wg.Done()
	}()

	go func() {
		fmt.Println("runner b")
		wg.Done()
	}()

	go func() {
		time.Sleep(2 * time.Second)
		runtime.Gosched()
		fmt.Println("休息了2秒的runner c并暂停了线程后，开始起跑了！")
		wg.Done()
	}()

	for i := 0; i < 10; i++ {
		fmt.Printf("runner d 跑了 %d百米\n", i)
		if i == 7 {
			fmt.Printf("还有%d百米的时候，d暂停自身线程\n", i)
			runtime.Gosched()
		}
	}
	wg.Wait()
}
