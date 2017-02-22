package references

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
	"cn.didadu/jvmgo/rtdata/heap"
)

// new指令结构体(2个字节操作数)
type NEW struct{ base.Index16Instruction }

func (self *NEW) Execute(frame *rtdata.Frame) {
	// 获取运行时常量池
	cp := frame.Method().Class().ConstantPool()
	// 通过索引从常量池中获取类符号引用
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	// 解析类
	class := classRef.ResolvedClass()

	// 判断类初始化是否已经开始
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	// 接口和抽象类不能实例化，抛出InstantiationError异常
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	// 创建对象引用
	ref := class.NewObject()
	// 将对象引用入栈
	frame.OperandStack().PushRef(ref)
}