package control

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// goto指令结构体
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtdata.Frame) {
	// 无条件跳转
	base.Branch(frame, self.Offset)
}
