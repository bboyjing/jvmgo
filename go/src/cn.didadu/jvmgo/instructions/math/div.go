package math

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// double除法结构体
type DDIV struct{ base.NoOperandsInstruction }

func (self *DDIV) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 / v2
	stack.PushDouble(result)
}

// float除法结构体
type FDIV struct{ base.NoOperandsInstruction }

func (self *FDIV) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 / v2
	stack.PushFloat(result)
}

// int除法结构体
type IDIV struct{ base.NoOperandsInstruction }

func (self *IDIV) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	// 弹出栈顶元素，作为除数
	v2 := stack.PopInt()
	// 弹出栈顶元素，作为被除数
	v1 := stack.PopInt()
	// 若除数为0，报错
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	// 除法运算
	result := v1 / v2
	// 将除法运算结果入栈
	stack.PushInt(result)
}

// long除法结构体
type LDIV struct{ base.NoOperandsInstruction }

func (self *LDIV) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 / v2
	stack.PushLong(result)
}
