package singleton

//懒汉
type SingleTonDemon1 struct {
}

var singleton *SingleTonDemon1

func GetInstance1() *SingleTonDemon1 {
	if singleton == nil {
		singleton = &SingleTonDemon1{}
	}

	return singleton
}
