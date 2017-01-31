package loads

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// iload指令结构体
type ILOAD struct{ base.Index8Instruction }

// 通过索操作局部变量表
func (self *ILOAD) Execute(frame *rtdata.Frame) {
	_iload(frame, uint(self.Index))
}

// 统一的iload函数
func _iload(frame *rtdata.Frame, index uint) {
	// 通过索引读取局部变量表
	val := frame.LocalVars().GetInt(index)
	// 将局部变量表中的值推入栈顶
	frame.OperandStack().PushInt(val)
}

// 操作第0号局部变量，索引隐含在操作码中
type ILOAD_0 struct{ base.NoOperandsInstruction }
func (self *ILOAD_0) Execute(frame *rtdata.Frame) {
	_iload(frame, 0)
}

// 操作第1号局部变量，索引隐含在操作码中
type ILOAD_1 struct{ base.NoOperandsInstruction }
func (self *ILOAD_1) Execute(frame *rtdata.Frame) {
	_iload(frame, 1)
}

// 操作第2号局部变量，索引隐含在操作码中
type ILOAD_2 struct{ base.NoOperandsInstruction }
func (self *ILOAD_2) Execute(frame *rtdata.Frame) {
	_iload(frame, 2)
}

// 操作第3号局部变量，索引隐含在操作码中
type ILOAD_3 struct{ base.NoOperandsInstruction }
func (self *ILOAD_3) Execute(frame *rtdata.Frame) {
	_iload(frame, 3)
}