package math

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// double取反结构体
type DNEG struct{ base.NoOperandsInstruction }

func (self *DNEG) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}

// float取反结构体
type FNEG struct{ base.NoOperandsInstruction }

func (self *FNEG) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}

// int取反结构体
type INEG struct{ base.NoOperandsInstruction }

func (self *INEG) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	// 弹出栈顶元素
	val := stack.PopInt()
	// 将取反结果入栈
	stack.PushInt(-val)
}

// long取反结构体
type LNEG struct{ base.NoOperandsInstruction }

func (self *LNEG) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}
