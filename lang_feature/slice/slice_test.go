package slice

import (
	"testing"
	"fmt"
	"github.com/pkg/errors"
)

func TestSlice(t *testing.T) {

	s := make([]int, 0, 100)
	s1 := s[:2:4]
	s2 := append(s1, 1, 2, 3, 4, 5, 6) // s1 cap
	fmt.Printf("s1: %p: %v\n", &s1[0], s1)
	fmt.Printf("s2: %p: %v\n", &s2[0], s2)
	fmt.Printf("s data: %v\n", s[:10])
	fmt.Printf("s1 cap: %d,s2 cap: %d\n", cap(s1), cap(s2))
}

func TestSliceCopy(t *testing.T) {
	names := []string{"Tom", "Jerry"}
	nums := []string{"one", "two", "three"}
	pNames := names // 确认 names 被更新

	fmt.Println(copy(names, nums))   //copy nums to names，返回长度以较短数组的长度为准
	fmt.Println(names, nums, pNames) //names= one two
}

//通过reslice实现一个栈
func TestSliceStack(t *testing.T) {
	stack := make([]int, 0, 5)
	//stack push
	push := func(x int) error {
		n := len(stack)
		if n == cap(stack) {
			return errors.New("stack is full")
		}
		stack = stack[:n+1]
		stack[n] = x
		return nil
	}
	//stack pop
	pop := func() (int, error) {
		n := len(stack)
		if n == 0 {
			return 0, errors.New("stack is empty")
		}
		x := stack[n-1]
		stack = stack[:n-1]
		return x, nil
	}
	//stack len
	stackLen := func() int {
		if stack == nil {
			return 0
		}
		return len(stack)
	}
	//
	for i := 0; i < 7; i++ {
		fmt.Printf("push%d: %v, %v\n", i, push(i), stack)
		fmt.Println(stackLen())
	}
	//
	for i := 0; i < 7; i++ {
		x, err := pop()
		fmt.Printf("pop: %d, %v, %v\n", x, err, stack)
		fmt.Println(stackLen())
	}
}
