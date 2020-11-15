package abstract_factory

import "testing"

func TestNewSimpleCourseFactory(t *testing.T) {
	var factory OrderFactory
	factory = &SnackFactory{}
	factory.CreateOrderInfo().queryOrderInfo()
	factory.CreateOrderDetail().saveOrderDetail()
}