package references

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
	"cn.didadu/jvmgo/rtdata/heap"
)

// invokeinterface指令结构体
type INVOKE_INTERFACE struct {
	index uint
	// count uint8
	// zero uint8
}

func (self *INVOKE_INTERFACE) FetchOperands(reader *base.BytecodeReader) {
	self.index = uint(reader.ReadUint16())
	// 下面两个字节读出来丢弃
	reader.ReadUint8() // count
	reader.ReadUint8() // must be 0
}

func (self *INVOKE_INTERFACE) Execute(frame *rtdata.Frame) {
	// 获取当前class的运行时常量池
	cp := frame.Method().Class().ConstantPool()
	// 通过索引获取方法符号引用
	methodRef := cp.GetConstant(self.index).(*heap.InterfaceMethodRef)
	// 解析需要调用的接口方法
	resolvedMethod := methodRef.ResolvedInterfaceMethod()
	if resolvedMethod.IsStatic() || resolvedMethod.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 获取距离操作数栈顶n个slot的引用变量
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException") // todo
	}
	if !ref.Class().IsImplements(methodRef.ResolvedClass()) {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 查找最终要调用的方法
	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(),
		methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	if !methodToBeInvoked.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}


