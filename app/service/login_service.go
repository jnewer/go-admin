package service

import (
	"github.com/cilidm/toolbox/gconv"
	"pear-admin-go/app/core/cache"
	"pear-admin-go/app/util/e"
	"time"
)

func Lock(loginName string) {
	cache.Instance().Set(e.UserLock+loginName, true, time.Minute*5)
}

func UnLock(loginName string) {
	cache.Instance().Delete(e.UserLock + loginName)
}

func CheckLock(loginName string) bool {
	res, ok := cache.Instance().Get(e.UserLock + loginName)
	if ok && res == true {
		return true
	}
	return false
}

func SetPwdErrNum(loginName string) int {
	countNum := 0
	errNum, _ := cache.Instance().Get(e.UserLoginErr + loginName)
	if errNum != nil {
		countNum = gconv.Int(errNum)
	}
	countNum = countNum + 1
	cache.Instance().Set(e.UserLoginErr+loginName, countNum, time.Minute*1)
	if countNum >= 5 {
		Lock(loginName)
	}
	return countNum
}

func RemovePwdErrNum(loginName string) {
	cache.Instance().Delete(e.UserLoginErr + loginName)
}
