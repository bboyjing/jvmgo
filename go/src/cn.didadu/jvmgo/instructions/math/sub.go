package math

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// double减法结构体
type DSUB struct{ base.NoOperandsInstruction }

func (self *DSUB) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 - v2
	stack.PushDouble(result)
}

// float减法结构体
type FSUB struct{ base.NoOperandsInstruction }

func (self *FSUB) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 - v2
	stack.PushFloat(result)
}

// int减法结构体
type ISUB struct{ base.NoOperandsInstruction }

func (self *ISUB) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	// 弹出栈顶元素
	v2 := stack.PopInt()
	// 弹出栈顶元素
	v1 := stack.PopInt()
	// 减法运算
	result := v1 - v2
	// 将减法运算结果入栈
	stack.PushInt(result)
}

// long减法结构体
type LSUB struct{ base.NoOperandsInstruction }

func (self *LSUB) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 - v2
	stack.PushLong(result)
}
