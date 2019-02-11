package channel

import (
	"testing"
	"sync"
)

func TestMemoryLeak(t *testing.T) {
	wg := new(sync.WaitGroup)
	channel := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		channel <- 1 //无缓冲chan，必须收发成对操作，否则报错fatal error: all goroutines are asleep - deadlock!
	}()
	wg.Wait()
}
