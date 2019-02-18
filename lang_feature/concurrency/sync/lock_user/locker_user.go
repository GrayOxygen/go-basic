package lock_user

import (
	"go-basic/lang_feature/concurrency/sync/lock"
)

//去掉超过一分钟的请求限制
func RemReplicaReq() {
	singleton_util.RemReplicaReq()
}
