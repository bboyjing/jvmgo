package references

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
	"cn.didadu/jvmgo/rtdata/heap"
)

// invokestatic指令结构体
type INVOKE_STATIC struct{ base.Index16Instruction }

func (self *INVOKE_STATIC) Execute(frame *rtdata.Frame) {
	// 获取运行时常量池
	cp := frame.Method().Class().ConstantPool()
	// 通过索引获取方法符号引用
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	// 解析方法
	resolvedMethod := methodRef.ResolvedMethod()
	// 判断方法是否是static
	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	class := resolvedMethod.Class()
	// 判断类初始化是否已经开始
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	base.InvokeMethod(frame, resolvedMethod)
}
