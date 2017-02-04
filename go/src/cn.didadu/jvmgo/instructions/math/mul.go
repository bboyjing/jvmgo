package math

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)


// double乘法结构体
type DMUL struct{ base.NoOperandsInstruction }

func (self *DMUL) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 * v2
	stack.PushDouble(result)
}

// float乘法结构体
type FMUL struct{ base.NoOperandsInstruction }

func (self *FMUL) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 * v2
	stack.PushFloat(result)
}

// int乘法结构体
type IMUL struct{ base.NoOperandsInstruction }

func (self *IMUL) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	// 弹出栈顶元素
	v2 := stack.PopInt()
	// 弹出栈顶元素
	v1 := stack.PopInt()
	// 乘法运算
	result := v1 * v2
	// 将乘法运算结果入栈
	stack.PushInt(result)
}

// long乘法结构体
type LMUL struct{ base.NoOperandsInstruction }

func (self *LMUL) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 * v2
	stack.PushLong(result)
}
