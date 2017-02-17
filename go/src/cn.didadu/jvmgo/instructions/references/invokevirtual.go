package references

import (
	"cn.didadu/jvmgo/instructions/base"
	"fmt"
	"cn.didadu/jvmgo/rtdata"
	"cn.didadu/jvmgo/rtdata/heap"
)

// Invoke instance method; dispatch based on class
type INVOKE_VIRTUAL struct{ base.Index16Instruction }

// hack!
func (self *INVOKE_VIRTUAL) Execute(frame *rtdata.Frame) {
	// 获取运行时常量池
	cp := frame.Method().Class().ConstantPool()
	// 获取方法符号引用
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	// 当方法为println
	if methodRef.Name() == "println" {
		// 获取操作数栈
		stack := frame.OperandStack()
		// 通过描述符从操作数栈中弹出对应的值
		switch methodRef.Descriptor() {
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
			panic("println: " + methodRef.Descriptor())
		}
		stack.PopRef()
	}
}