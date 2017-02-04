package stack

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// dup结构体，复制栈顶单个变量
type DUP struct{ base.NoOperandsInstruction }

func (self *DUP) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	// 先弹出栈顶变量
	slot := stack.PopSlot()
	// 将弹出的变量2次入栈，达到复制效果
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

// dup_x1结构体，复制栈顶单个变量到距栈顶第三位
type DUP_X1 struct{ base.NoOperandsInstruction }

func (self *DUP_X1) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// dup_x2结构体，复制栈顶单个变量到距栈顶第四位
type DUP_X2 struct{ base.NoOperandsInstruction }

func (self *DUP_X2) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}


// dup2结构体，复制栈顶两个变量
type DUP2 struct{ base.NoOperandsInstruction }

func (self *DUP2) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// dup2_x1结构体，复制栈顶2个变量到距栈顶第四、第五位
type DUP2_X1 struct{ base.NoOperandsInstruction }

func (self *DUP2_X1) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// dup2_x2结构体，复制栈顶2个变量到距栈顶第五、第六位
type DUP2_X2 struct{ base.NoOperandsInstruction }

func (self *DUP2_X2) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	slot4 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot4)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}