package heap

// Object结构体
type Object struct {
	// 对象的Class指针
	class  *Class
	// 实例变量，可以容纳任何类型的值
	data  interface{}
}

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		data: newSlots(class.instanceSlotCount),
	}
}

func (self *Object) Class() *Class {
	return self.class
}

// Fields()方法仍然只针对普通对象，转成Slots
func (self *Object) Fields() Slots {
	return self.data.(Slots)
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}