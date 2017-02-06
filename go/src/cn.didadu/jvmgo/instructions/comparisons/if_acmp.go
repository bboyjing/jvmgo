package comparisons

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// if_acmpeq指令结构体
type IF_ACMPEQ struct{ base.BranchInstruction }

func (self *IF_ACMPEQ) Execute(frame *rtdata.Frame) {
	if _acmp(frame) {
		base.Branch(frame, self.Offset)
	}
}

func _acmp(frame *rtdata.Frame) bool {
	stack := frame.OperandStack()
	// 弹出栈顶引用
	ref2 := stack.PopRef()
	// 弹出栈顶引用
	ref1 := stack.PopRef()
	// 判断是否相等
	return ref1 == ref2
}

type IF_ACMPNE struct{ base.BranchInstruction }

func (self *IF_ACMPNE) Execute(frame *rtdata.Frame) {
	if !_acmp(frame) {
		base.Branch(frame, self.Offset)
	}
}