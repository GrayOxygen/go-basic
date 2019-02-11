package sync

import (
	"testing"
	"sync"
	"fmt"
)

type worker struct {
	name string
}

func (w *worker) work6Days() {
	fmt.Println(w.name, " worked over 5 days ")
}

//Add Done Wait方法，控制goroutine执行顺序
func TestWaitGroup(t *testing.T) {
	tasks := make(chan *worker, 64)

	// spawn four worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ { //模拟两个任务交叉执行：奇偶交叉
		wg.Add(1)
		go func(i int) {
			if i%2 == 1 {
				wg.Wait()
			}
			for t := range tasks {
				t.work6Days()
			}
			wg.Done()
		}(i)
		go func(i int) {
			if i%2 == 0 {
				wg.Wait()
			}
			for t := range tasks {
				t.work6Days()
			}
			wg.Done()
		}(i)
	}

	// generate some tasks
	for i := 0; i < 10; i++ {
		w := new(worker)
		w.name = fmt.Sprintf("%d号程序员", i)
		tasks <- w
	}
	close(tasks)

	// wait for the workers to finish
	wg.Wait()
}
