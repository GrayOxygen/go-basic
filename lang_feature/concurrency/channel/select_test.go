package channel

import (
	"testing"
	"sync"
	"fmt"
)

func TestSelect(t *testing.T) {
	wg := new(sync.WaitGroup)
	a, b := make(chan int), make(chan int)

	wg.Add(2)
	go func() { //receive select
		defer wg.Done()
		for {
			var (
				x    int
				ok   bool
				name string
			)
			select {
			case x, ok = <-a: //通道a关闭后，设成nil，则该case将不会被select选中
				if !ok {
					a = nil
				}
				name = "a"
			case x, ok = <-b:
				if !ok {
					b = nil
				}
				name = "b"
			default:
			}
			if a == nil && b == nil {
				return
			}
			fmt.Println(name, x)
		}
	}()

	go func() { //sender select
		defer wg.Done()
		defer close(a)
		defer close(b)

		for i := 0; i < 5; i++ {
			select {
			case a <- i:
			case b <- i * 10:
			}
		}
	}()
	wg.Wait() //等goroutine执行完
}
