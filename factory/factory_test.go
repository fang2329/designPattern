package factory

import (
	"fmt"
	"testing"
)

func TestPlusOperFactory_Create(t *testing.T) {
	var factory  OperatorFactory
	factory = PlusOperFactory{}
	op := factory.Create()
	op.SetLeft(10)
	op.SetRight(20)
	fmt.Println(op.Result())

	factory = SubOperatorFactory{}
	op = factory.Create()
	op.SetLeft(30)
	op.SetRight(10)
	fmt.Println(op.Result())
}