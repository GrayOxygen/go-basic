package lang_feature

import (
	"fmt"
	"testing"
)

func TestRecover(t *testing.T) {
	defer func() {
		fmt.Println("c")
	}()
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("d")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容
		}
		fmt.Println("e")
	}()

	fn()
	fmt.Println("f") //不再执行
}

func fn() {
	fmt.Println("a")
	panic("异常信息")
	fmt.Println("b") //不再执行
}
