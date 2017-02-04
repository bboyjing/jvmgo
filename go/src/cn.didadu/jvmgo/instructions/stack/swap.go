package stack

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// swap结构体，交换栈顶两个变量
type SWAP struct{ base.NoOperandsInstruction }

func (self *SWAP) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	// 栈顶两个变量出栈
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	// 按弹出的顺序将两个变量入栈，达到交换效果
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}