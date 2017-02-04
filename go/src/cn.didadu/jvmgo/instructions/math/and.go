package math

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// int与操作结构体
type IAND struct{ base.NoOperandsInstruction }

func (self *IAND) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	// 弹出栈顶元素作为操作数
	v2 := stack.PopInt()
	// 弹出栈顶元素作为操作数
	v1 := stack.PopInt()
	// 将弹出元素进行与操作
	result := v1 & v2
	// 将与操作结果入栈
	stack.PushInt(result)
}

// long与操作结构体
type LAND struct{ base.NoOperandsInstruction }

func (self *LAND) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 & v2
	stack.PushLong(result)
}