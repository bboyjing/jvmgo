package comparisons

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// ifeq指令结构体
type IFEQ struct{ base.BranchInstruction }

func (self *IFEQ) Execute(frame *rtdata.Frame) {
	// 弹出栈顶int变量
	val := frame.OperandStack().PopInt()
	if val == 0 {
		// 满足条件跳转
		base.Branch(frame, self.Offset)
	}
}

type IFNE struct{ base.BranchInstruction }

func (self *IFNE) Execute(frame *rtdata.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFLT struct{ base.BranchInstruction }

func (self *IFLT) Execute(frame *rtdata.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFLE struct{ base.BranchInstruction }

func (self *IFLE) Execute(frame *rtdata.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFGT struct{ base.BranchInstruction }

func (self *IFGT) Execute(frame *rtdata.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFGE struct{ base.BranchInstruction }

func (self *IFGE) Execute(frame *rtdata.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, self.Offset)
	}
}