package simple_factory

import "fmt"

//抽象的产品
type SuperMarket interface {
	GetGoods()
}

//参数的方式创建产品
func  NewSupermarket( name string) SuperMarket {
	switch name {
	case "s":
		return &Snacks{}
	case "v":
		return &Vegetables{}
	default:
		return nil
	}
}

//产品1
type Snacks struct {}

func (s *Snacks)GetGoods()  {
	fmt.Println("snacks")
}

//产品2
type Vegetables struct {}
func (v *Vegetables)GetGoods()  {
	fmt.Println("vegetables")
}







