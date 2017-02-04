package math

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// double加法结构体
type DADD struct{ base.NoOperandsInstruction }

func (self *DADD) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	result := v1 + v2
	stack.PushDouble(result)
}

// float加法结构体
type FADD struct{ base.NoOperandsInstruction }

func (self *FADD) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 + v2
	stack.PushFloat(result)
}

// int加法结构体
type IADD struct{ base.NoOperandsInstruction }

func (self *IADD) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	// 弹出栈顶元素
	v2 := stack.PopInt()
	// 弹出栈顶元素
	v1 := stack.PopInt()
	// 加法运算
	result := v1 + v2
	// 将加法运算结果入栈
	stack.PushInt(result)
}

// long加法结构体
type LADD struct{ base.NoOperandsInstruction }

func (self *LADD) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 + v2
	stack.PushLong(result)
}
