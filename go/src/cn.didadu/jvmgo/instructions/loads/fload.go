package loads

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// fload指令结构体
type FLOAD struct{ base.Index8Instruction }

// 通过索操作局部变量表
func (self *FLOAD) Execute(frame *rtdata.Frame) {
	_fload(frame, uint(self.Index))
}

// 统一的fload函数
func _fload(frame *rtdata.Frame, index uint) {
	// 通过索引读取局部变量表
	val := frame.LocalVars().GetFloat(index)
	// 将局部变量表中的值推入栈顶
	frame.OperandStack().PushFloat(val)
}

// 操作第0号局部变量，索引隐含在操作码中
type FLOAD_0 struct{ base.NoOperandsInstruction }
func (self *FLOAD_0) Execute(frame *rtdata.Frame) {
	_fload(frame, 0)
}

// 操作第1号局部变量，索引隐含在操作码中
type FLOAD_1 struct{ base.NoOperandsInstruction }
func (self *FLOAD_1) Execute(frame *rtdata.Frame) {
	_fload(frame, 1)
}

// 操作第2号局部变量，索引隐含在操作码中
type FLOAD_2 struct{ base.NoOperandsInstruction }
func (self *FLOAD_2) Execute(frame *rtdata.Frame) {
	_fload(frame, 2)
}

// 操作第3号局部变量，索引隐含在操作码中
type FLOAD_3 struct{ base.NoOperandsInstruction }
func (self *FLOAD_3) Execute(frame *rtdata.Frame) {
	_fload(frame, 3)
}