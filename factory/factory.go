package factory

//A X B
//AB操作数 X操作

/*
//实际运行类接口
type Operator interface {
	SetLeft(int)
	SetRight(int)
	Result() int
}

//工厂接口
type OperatorFactory interface {
	Create() Operator
}

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

type SubOperatorFactory struct {

}

type SubOperator struct {
	*OperatorBase
}

func (o *SubOperator)Result() int  {
	return o.left - o.right
}
func (SubOperatorFactory)Create()  Operator {
	return &SubOperator{&OperatorBase{}}
}
*/




