package prototype

//原型模式

//原型对象需要实现的接口
//拷贝原有的数据
type CloneAble interface {
	Clone() CloneAble
}

//原型的类
type ProtoTypeManager struct {
	prototypes map[string]CloneAble
}

func NewProtoTypeManager() *ProtoTypeManager  {
	return &ProtoTypeManager{prototypes:make(map[string]CloneAble)}
}

func (p *ProtoTypeManager)Get( name string)CloneAble  {
	return p.prototypes[name]
}

func (p *ProtoTypeManager)Set(name string,prototype CloneAble)  {
	p.prototypes[name] = prototype
}

type Type1 struct {
	name string
}

func (t *Type1)Clone() CloneAble  {
	tc := *t
	return &tc
}

type Type2 struct {
	name string
}

func (t *Type2)Clone() CloneAble  {
	tc := *t
	return &tc
}