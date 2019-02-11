package util

import (
	"time"
	"fmt"
)

func CountTime(start time.Time) {
	dis := time.Now().Sub(start).Seconds()
	fmt.Println("消耗秒数:", dis)
}
