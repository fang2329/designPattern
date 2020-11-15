package singleton

type SingletonDemon2 struct {

}

//饿汉
var singletonDemon2 = &SingletonDemon2{}

func GetInstance2()  *SingletonDemon2{
	return singletonDemon2
}
