package singleton_util

import (
	"fmt"
	"sync"
	"time"
)

//单例模式
//map的并发访问
var reqLogLock *sync.Mutex
var once = new(sync.Once)
var onceLock = new(sync.Once)

//移除和新增的锁
func GetLockInstance() *sync.Mutex {
	onceLock.Do(func() {
		fmt.Println("GetInstanceLock")
		reqLogLock = new(sync.Mutex)
	})
	return reqLogLock
}

type HaitunReqLog struct {
	Param    string
	CallTime time.Time
}

//请求信息记录
var haitunReqLogMap map[string]*HaitunReqLog //url-log
//获取单例的请求信息
func GetInstance() map[string]*HaitunReqLog {
	once.Do(func() {
		fmt.Println("GetInstance")
		haitunReqLogMap = make(map[string]*HaitunReqLog, 1)
	})
	return haitunReqLogMap
}

//读取
func ValidReqFreq(url, param string) bool {
	GetLockInstance().Lock()
	defer GetLockInstance().Unlock()
	haitunReqLogMap = GetInstance()
	if _, ok := haitunReqLogMap[url]; !ok || haitunReqLogMap[url] == nil || haitunReqLogMap[url].Param != param {
		return true
	}
	dataTimeStr := time.Unix(haitunReqLogMap[url].CallTime.Unix(), 0).Format("2006-01-02 15:04:05")
	fmt.Println("ValidReqFreq:", dataTimeStr)
	return time.Now().Sub(haitunReqLogMap[url].CallTime).Minutes() <= 1
}

//新增上锁
func AddReqFreq(url, param string) {
	GetLockInstance().Lock()
	defer GetLockInstance().Unlock()
	haitunReqLogMap = GetInstance()
	reqLog := new(HaitunReqLog)
	reqLog.CallTime = time.Now()
	reqLog.Param = param
	haitunReqLogMap[url] = reqLog
	fmt.Println(url, reqLog)
}

//删除上锁
func DeleteMap(url string) {
	GetLockInstance().Lock()
	defer GetLockInstance().Unlock()
	delete(GetInstance(), url)
}

func RemReplicaReq() {
	instanceMap := GetInstance()
	GetLockInstance().Lock()
	defer GetLockInstance().Unlock()
	for key, val := range instanceMap {
		if time.Now().Sub(val.CallTime).Minutes() > 1 {
			dataTimeStr := time.Unix(val.CallTime.Unix(), 0).Format("2006-01-02 15:04:05")
			fmt.Printf("\nurl为【%s】param为【%s】调用时间为【%s】 已经过期\n", key, val.Param, dataTimeStr)

			delete(GetInstance(), key)
		}
	}
}
