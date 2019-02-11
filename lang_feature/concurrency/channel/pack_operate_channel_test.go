package channel

import (
	"testing"
	"time"
	"exercise/go-basic/lang_feature/util"
)

//channel实际还是通过锁实现，如果每次收发操作，都是批处理的，将会提升效率
func TestBatchChannel(t *testing.T) {
	start := time.Now()
	test()
	util.CountTime(start)
	//下面的速度快了将近两倍
	start2 := time.Now()
	testBlock()
	util.CountTime(start2)
}

const (
	max     = 50000000 // 最大容量
	block   = 500      // 单个元素：数组  的大小
	bufsize = 100      // 通道大小
)

func test() { // :
	done := make(chan struct{})
	c := make(chan int, bufsize)
	go func() {
		count := 0
		for x := range c {
			count += x
		}
		close(done)
	}()
	for i := 0; i < max; i++ {
		c <- i
	}
	close(c)
	<-done
}

func testBlock() { // : 500
	done := make(chan struct{})
	c := make(chan [block]int, bufsize)
	go func() {
		count := 0
		for a := range c {
			for _, x := range a {
				count += x
			}
		}
		close(done)
	}()

	for i := 0; i < max; i += block {
		var b [block]int //
		for n := 0; n < block; n++ {
			b[n] = i + n
			if i+n == max-1 {
				break
			}
		}
		c <- b
	}
	close(c)
	<-done
}
