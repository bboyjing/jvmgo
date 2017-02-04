package math

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// int或操作结构体
type IOR struct{ base.NoOperandsInstruction }

func (self *IOR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	// 弹出栈顶元素
	v2 := stack.PopInt()
	// 弹出栈顶元素
	v1 := stack.PopInt()
	// 将弹出的操作数进行或操作
	result := v1 | v2
	// 将或操作结果入栈
	stack.PushInt(result)
}

// long或操作结构体
type LOR struct{ base.NoOperandsInstruction }

func (self *LOR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 | v2
	stack.PushLong(result)
}
