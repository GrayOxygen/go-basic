package channel

import (
	"testing"
	"sync"
	"fmt"
)

type Toy struct {
	id string
}
type toyFactory struct {
	sync.WaitGroup
	toyz chan *Toy
}

//每次生产5000件，50个任务并发
func (factory *toyFactory) prodToy() {
	factory.Add(50)
	for i := 0; i < 50; i++ {
		go func(n int) {
			fmt.Printf("\n第%d个生产者任务\n", n)
			for i := 0; i < 20; i++ { //每个任务都生产20个玩具，如果生产数大于channel缓冲区长度，则会死锁panic
				tempToy := new(Toy)
				tempToy.id = fmt.Sprintf("%d,%d", n, i)
				factory.toyz <- tempToy
			}
			factory.Done()
		}(i)
	}
	factory.Wait()
	close(factory.toyz) //只能关闭发送端
}

//消费：100个任务并发消费channel中数据
func (factory *toyFactory) cosumeToy() {
	factory.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Printf("\n第%d个消费者任务\n", i)
			for tempToy := range factory.toyz {
				fmt.Println("消费玩具：", tempToy.id)
			}
			factory.Done()
		}(i)
	}
	factory.Wait()
}

//通常用工厂方法将goroutine和channel绑定
func TestChannelWithFactory(t *testing.T) {
	factory := new(toyFactory)
	factory.toyz = make(chan *Toy, 5000)
	factory.prodToy()
	factory.cosumeToy()
}
