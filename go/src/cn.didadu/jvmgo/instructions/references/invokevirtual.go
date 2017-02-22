package references

import (
	"cn.didadu/jvmgo/instructions/base"
	"fmt"
	"cn.didadu/jvmgo/rtdata"
	"cn.didadu/jvmgo/rtdata/heap"
)

// invokevirtual指令结构体
type INVOKE_VIRTUAL struct{ base.Index16Instruction }

func (self *INVOKE_VIRTUAL) Execute(frame *rtdata.Frame) {
	// 获取当前class
	currentClass := frame.Method().Class()
	// 获取当前类的运行时常量池
	cp := currentClass.ConstantPool()
	// 通过索引获取方法符号引用
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	// 解析需要调用的方法
	resolvedMethod := methodRef.ResolvedMethod()
	// 判断方法是否是static
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 获取距离操作数栈顶n个slot的引用变量，也就是需要调用的方法所属的类的实例引用
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		// hack!
		if methodRef.Name() == "println" {
			_println(frame.OperandStack(), methodRef.Descriptor())
			return
		}

		panic("java.lang.NullPointerException")
	}

	if resolvedMethod.IsProtected() &&
		resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&
		!ref.Class().IsSubClassOf(currentClass) {

		panic("java.lang.IllegalAccessError")
	}

	// 在获取的引用变量，也就是方法所属的类中查找真正要调用的方法
	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(),
		methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}

// hack!
func _println(stack *rtdata.OperandStack, descriptor string) {
	switch descriptor {
	case "(Z)V":
		fmt.Printf("%v\n", stack.PopInt() != 0)
	case "(C)V":
		fmt.Printf("%c\n", stack.PopInt())
	case "(I)V", "(B)V", "(S)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(F)V":
		fmt.Printf("%v\n", stack.PopFloat())
	case "(J)V":
		fmt.Printf("%v\n", stack.PopLong())
	case "(D)V":
		fmt.Printf("%v\n", stack.PopDouble())
	default:
		panic("println: " + descriptor)
	}
	stack.PopRef()
}