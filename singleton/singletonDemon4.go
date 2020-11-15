package singleton

import "sync"

//双重锁
type SingleTonDemon4 struct {
}

var singletonDemon4 *SingleTonDemon4
var mutx sync.Mutex

func GetInstance4() *SingleTonDemon4{
	if singletonDemon4 == nil{
		mutx.Lock()
		if singletonDemon4  == nil{
			 singletonDemon4 = &SingleTonDemon4{}
			 mutx.Unlock()
			 return singletonDemon4
		}

	}
	mutx.Unlock()
	return singletonDemon4
}
