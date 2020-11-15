package factory

//操作的抽象 加法
type PlusOperFactory struct {}

//操作类中包含操作苏
type PlusOperator struct {
	*OperatorBase
}

//实际运行
func (o *PlusOperator)Result() int  {
	return o.left + o.right
}
func (PlusOperFactory)Create()  Operator {
	return &PlusOperator{&OperatorBase{}}
}
