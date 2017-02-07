package extended

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// ifnull指令结构体
type IFNULL struct{ base.BranchInstruction }

func (self *IFNULL) Execute(frame *rtdata.Frame) {
	// 弹出栈顶引用变量
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		// 如果变量是null，则跳转
		base.Branch(frame, self.Offset)
	}
}


// Branch if reference not null
type IFNONNULL struct{ base.BranchInstruction }

func (self *IFNONNULL) Execute(frame *rtdata.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, self.Offset)
	}
}