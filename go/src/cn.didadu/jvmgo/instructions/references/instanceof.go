package references

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
	"cn.didadu/jvmgo/rtdata/heap"
)

// instanceof指令结构体
type INSTANCE_OF struct{ base.Index16Instruction }

func (self *INSTANCE_OF) Execute(frame *rtdata.Frame) {
	// 获取操作数栈
	stack := frame.OperandStack()
	// 弹出栈顶对象引用
	ref := stack.PopRef()
	// 如果对象引用为空，将0入栈
	if ref == nil {
		stack.PushInt(0)
		return
	}

	// 获取运行时常量池
	cp := frame.Method().Class().ConstantPool()
	// 通过索引从常量池中获取类符号引用
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	// 解析类
	class := classRef.ResolvedClass()
	// 判断对象引用是不是class的实例
	if ref.IsInstanceOf(class) {
		// 如果是，将1入栈
		stack.PushInt(1)
	} else {
		// 如果不是，将0入栈
		stack.PushInt(0)
	}
}
