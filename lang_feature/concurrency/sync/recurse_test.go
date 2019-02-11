package sync

import (
	"testing"
	"sync"
	"fmt"
	"time"
)

func TestRecurse(t *testing.T) {

	var lock sync.Mutex
	lock.Lock()
	{
		lock.Lock()
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
		lock.Unlock()
	}
	lock.Unlock()

}
func TestRecurse2(t *testing.T) {

	var lock sync.Mutex
	lock.Lock()
	go func() {
		lock.Lock()
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
		lock.Unlock()
	}()
	time.Sleep(5 * time.Second)
	lock.Unlock()

}
