package lang_feature

import (
	"fmt"
	"testing"
	"time"
)

type student struct {
	Name string
	Age  int
}

func paseStudent() map[string]*student {
	m := make(map[string]*student)
	stus := []student{
		{Name: "wang", Age: 28},
		{Name: "qian", Age: 23},
		{Name: "ma", Age: 29},
	}
	for _, stu := range stus {
		//&stu为stu的指针，而stu存的值只有最有一个元素
		//m[stu.Name] = &stu
		//如下每次指向一个单独的指针，则不会有问题
		temp := stu
		m[stu.Name] = &temp
	}

	return m
}

//引用和值的区分
func TestReferVal(t *testing.T) {
	students := paseStudent()
	for k, v := range students {
		fmt.Printf("key=%s,value=%v \n", k, v)
	}
}

func TestTime(t *testing.T) {
	mm, _ := time.ParseDuration("5m")
	mm1 := time.Now().Add(mm)
	fmt.Println(mm1)

}
