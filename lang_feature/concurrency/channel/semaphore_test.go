package channel

import (
	"testing"
	"fmt"
	"sync"
	"time"
)

func TestSemaphore(t *testing.T) {
	sema := make(chan struct{}, 2)
	wg := new(sync.WaitGroup)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			sema <- struct{}{}
			defer func() {
				<-sema
			}()
			time.Sleep(2 * time.Second)
			fmt.Println("doing task:", n)
		}(i)
	}
	wg.Wait()
}
