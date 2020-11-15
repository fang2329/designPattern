package factory


//数据抽离
type OperatorBase struct {
	left,right int
}

func (o *OperatorBase)SetLeft(leftval int)  {
	o.left = leftval
}
func (o *OperatorBase)SetRight(rightVal int)  {
	o.right = rightVal
}
