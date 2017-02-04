package stack

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// pop结构体
type POP struct{ base.NoOperandsInstruction }

func (self *POP) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	// 弹出栈顶变量
	stack.PopSlot()
}

// pop2结构体
type POP2 struct{ base.NoOperandsInstruction }

func (self *POP2) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	// 弹出栈顶变量
	stack.PopSlot()
	stack.PopSlot()
}