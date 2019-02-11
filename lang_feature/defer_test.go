package lang_feature

import (
	"testing"
	"fmt"
)

func sum(x, y int) int {
	return x + y
}

/**
1，栈实现的LIFO

2，在return之前，即方法结束之前依次执行defer函数
	return xxx并非原子指令，实为：
	result=xxx
	defer func执行，defer函数引用到的值是在执行前就定下来的，即defer前面那部分代码的出来的
	return

3，defer中的变量在return前确定
 */
func TestDefer(t *testing.T) {
	defer func() {
		fmt.Println("第一个defer")
	}()

	var deferVal int
	datas := []int{1, 2, 3, 2, 3, 4, 1, 1, 2}
	for i := 0; i < len(datas); i = i + 3 {
		res := sum(datas[i], datas[i+1])
		if res != datas[i+2] {
			defer func() {
				deferVal = i
				fmt.Printf("\n第%d次计算出错，执行defer\n", i+1) //这里即最后的i为9
			}()
		}
	}

	fmt.Printf("\nreturn前，defer并未赋值 %d\n", deferVal)

}

//测试return的非原子性
func TestDeferReturn(t *testing.T) {
	if f() == 5 {
		fmt.Println("执行正确")
	}
}
func f() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

//测试defer stack实现
func TestDeferStack(t *testing.T) {
	a := 1
	b := 2
	defer calc(a, calc(a, b))
	a = 0
	defer calc(a, calc(a, b))
}

/**
1 2 3
0 2 2
0 2 2
1 3 4
 */
func calc(x, y int) int {
	fmt.Println(x, y, x+y)
	return x + y
}
