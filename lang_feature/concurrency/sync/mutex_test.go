package sync

import (
	"testing"
	"sync"
	"fmt"
)

type operator struct {
	sync.Mutex
}

//因为Mutex为operator匿名属性，下面这个方法op必须是指针传递，否则锁失效，即op *operator
func (op operator) fix(str string) {
	op.Lock()
	defer op.Unlock()
	for i := 0; i < 5; i++ {
		fmt.Println(str, i)
	}

}

func TestMutex(t *testing.T) {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	operator := new(operator)
	go func() {
		defer wg.Done()
		operator.fix("read")
	}()
	go func() {
		defer wg.Done()
		operator.fix("write")
	}()

	wg.Wait()
}
