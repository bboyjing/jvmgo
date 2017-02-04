package math

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
	"math"
)

// int求余结构体
type IREM struct{ base.NoOperandsInstruction }

func (self *IREM) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	// 弹出栈顶元素，作为除数
	v2 := stack.PopInt()
	// 弹出栈顶元素，作为被除数
	v1 := stack.PopInt()
	// 若除数为0，报错
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	// 求余运算
	result := v1 % v2
	// 将求余的结果入栈
	stack.PushInt(result)
}

// long求余结构体
type LREM struct{ base.NoOperandsInstruction }

func (self *LREM) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 % v2
	stack.PushLong(result)
}

// float求余结构体
type FREM struct{ base.NoOperandsInstruction }

func (self *FREM) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := float32(math.Mod(float64(v1), float64(v2)))
	stack.PushFloat(result)
}

// double求余结构体
type DREM struct{ base.NoOperandsInstruction }

func (self *DREM) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := math.Mod(v1, v2)
	stack.PushDouble(result)
}

