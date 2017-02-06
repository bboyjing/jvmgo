package conversions

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// double强转成float结构体
type D2F struct{ base.NoOperandsInstruction }

func (self *D2F) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	// 弹出栈顶double型变量
	d := stack.PopDouble()
	// 强转成float
	f := float32(d)
	// 将强转后的float变量入栈
	stack.PushFloat(f)
}

// Convert double to int
type D2I struct{ base.NoOperandsInstruction }

func (self *D2I) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

// Convert double to long
type D2L struct{ base.NoOperandsInstruction }

func (self *D2L) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}