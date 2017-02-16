package references

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
	"cn.didadu/jvmgo/rtdata/heap"
)


// instanceof指令结构体
type CHECK_CAST struct{ base.Index16Instruction }

func (self *CHECK_CAST) Execute(frame *rtdata.Frame) {
	// 获取操作数栈
	stack := frame.OperandStack()
	// 弹出栈顶对象引用
	ref := stack.PopRef()
	// 将弹出的栈顶对象引用再次入栈，不改变操作数栈状态
	stack.PushRef(ref)
	// 如果引用为null，指令结束。表示null可以转成任何类型
	if ref == nil {
		return
	}

	// 获取运行时常量池
	cp := frame.Method().Class().ConstantPool()
	// 通过索引从常量池中获取类符号引用
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	// 解析类
	class := classRef.ResolvedClass()
	// 判断对象引用是不是class的实例，不是的话抛出ClassCastException异常
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}
