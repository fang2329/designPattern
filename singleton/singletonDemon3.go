package singleton

import (
	"sync"
)

//懒汉加锁
type SingleTonDemon3 struct {
}

var singletonDemon3 *SingleTonDemon3
var mutex sync.Mutex

func GetInstance3() *SingleTonDemon3{
	mutex.Lock()
	defer  mutex.Unlock()
	if singletonDemon3  == nil{
		singletonDemon3 = &SingleTonDemon3{}
	}
	return singletonDemon3
}
