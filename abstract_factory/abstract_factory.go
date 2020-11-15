package abstract_factory

import "fmt"

type OderInfo interface {
	queryOrderInfo()
}
type OderDetail interface {
	saveOrderDetail()
}
type OrderFactory interface {
	CreateOrderInfo() OderInfo
	CreateOrderDetail() OderDetail
} 

type SnackOrder struct{}

func (s *SnackOrder)queryOrderInfo()  {
	fmt.Println("snack orderInfo")
}

type SnackOrderDetai struct {}

func (s *SnackOrderDetai)saveOrderDetail()  {
	fmt.Println("snack saveOrderDetail")
}

type SnackFactory struct {}

func (s *SnackFactory) CreateOrderInfo() OderInfo{
	return &SnackOrder{}
}

func (s *SnackFactory) CreateOrderDetail() OderDetail {
	return &SnackOrderDetai{}
}