package main

import (
	"fmt"
	"go-basic/lang_feature/concurrency/sync/lock"
	"go-basic/lang_feature/concurrency/sync/lock_user"
	"strconv"
	"sync"
	"time"
)

func main() {

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 500; i++ {
			singleton_util.AddReqFreq("加入的url"+strconv.Itoa(i), "参数"+strconv.Itoa(i))
		}
	}()
	go func() {
		defer wg.Done()
		for {
			lock_user.RemReplicaReq()
			time.Sleep(20 * time.Second)
		}
	}()
	go func() {
		for {
			for i := 0; i < 15; i++ {
				fmt.Println("请求长度", len(singleton_util.GetInstance()))
				fmt.Println(strconv.Itoa(0), singleton_util.ValidReqFreq("加入的url"+strconv.Itoa(0), "参数"+strconv.Itoa(0)))
				time.Sleep(5 * time.Second)
			}
		}
	}()
	time.Sleep(10 * time.Minute)
	wg.Wait()
}
