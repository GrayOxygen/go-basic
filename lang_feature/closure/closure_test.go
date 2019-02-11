package closure

import (
	"testing"
	"fmt"
)

func TestClosure1(t *testing.T) {
	r1 := f(0)
	fmt.Println(r1(), r1())
}

//闭包：函数中的函数，不随外部函数执行完毕而消除自身函数和它引用的环境
//函数的局部变量escape到堆中，是实现闭包的基础
func f(i int) func() int {
	a := make(chan [12]int)
	fmt.Println(a)
	return func() int {
		i++ //return后，这个i会另外保存在堆中（此处命令为 i=new(int)  ）
		return i
	}

}
