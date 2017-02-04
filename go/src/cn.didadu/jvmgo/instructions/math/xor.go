package math

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// int异或操作
type IXOR struct{ base.NoOperandsInstruction }

func (self *IXOR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	// 弹出栈顶元素
	v1 := stack.PopInt()
	// 弹出栈顶元素
	v2 := stack.PopInt()
	// 将弹出的操作数进行异或操作
	result := v1 ^ v2
	// 将异或操作结果入栈
	stack.PushInt(result)
}

// long异或操作
type LXOR struct{ base.NoOperandsInstruction }

func (self *LXOR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	result := v1 ^ v2
	stack.PushLong(result)
}
