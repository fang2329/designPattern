package prototype

import (
	"fmt"
	"testing"
)

func TestNewProtoTypeManager(t *testing.T) {
	mgr := NewProtoTypeManager()
	t1 := &Type1{name:"type1"}
	mgr.Set("t1",t1)
	t11 := mgr.Get("t1")
	t22 := t11.Clone()
	if t11 == t22 {
		fmt.Println("clone address")
	}else{
		fmt.Println("clone value")
	}
}
