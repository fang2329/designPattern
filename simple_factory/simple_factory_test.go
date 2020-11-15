package simple_factory

import "testing"

func TestSupermarketFactory_NewSupermarket(t *testing.T) {

	NewSupermarket("s").GetGoods()
	NewSupermarket("v").GetGoods()
}
